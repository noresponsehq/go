package noresponse

// Late-result handlers. Hand-written beside the generated client.
//
// An assessment can be revised after the signup was allowed: a slow source
// settles and the recommendation rises. NoResponse sends the new version as a
// signed assessment.updated webhook. HandleLateResult verifies that delivery
// with ConstructEvent, picks an action from your policy and acts through the
// auth provider's admin API: revoke the user's sessions or ban the user.
//
// The handler verifies the delivery itself, so it never acts on an unverified
// event, and it never acts on a test event from WebhookEndpoints.SendTestEvent.
// The provider calls are thin: you pass the credentials, and each action is
// the provider's documented admin endpoint. A provider 404 for a user or
// session that is already gone counts as done, so a repeated delivery is safe.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/noresponsehq/go/v2/models/components"
)

// LateResultAction is what to do to the user. LateResultBan also ends every session.
type LateResultAction string

const (
	LateResultNone           LateResultAction = "none"
	LateResultRevokeSessions LateResultAction = "revoke_sessions"
	LateResultBan            LateResultAction = "ban"
)

// LateResultPolicy maps each recommendation to an action. A recommendation
// left out does nothing.
type LateResultPolicy map[components.Recommendation]LateResultAction

// DefaultLateResultPolicy revokes sessions on review and does nothing on allow
// or step_up. Banning is opt-in. It returns a fresh map each call.
func DefaultLateResultPolicy() LateResultPolicy {
	return LateResultPolicy{
		components.RecommendationAllow:  LateResultNone,
		components.RecommendationStepUp: LateResultNone,
		components.RecommendationReview: LateResultRevokeSessions,
	}
}

// ErrUnknownLateResultAction is returned when a policy names an action other
// than none, revoke_sessions or ban.
var ErrUnknownLateResultAction = errors.New("noresponse: unknown late result action")

// ErrEmptyUserID is returned by a provider asked to act on an empty user id.
var ErrEmptyUserID = errors.New("noresponse: user id must be a non-empty string")

// LateResultProvider is the two admin actions for one auth provider. Build one
// with ClerkActions, WorkOSActions or Auth0Actions.
type LateResultProvider interface {
	Name() string
	// RevokeSessions ends every active session of the user.
	RevokeSessions(ctx context.Context, userID string) error
	// Ban stops the user signing in again and ends their sessions.
	Ban(ctx context.Context, userID string) error
}

// LateResultOptions adjusts HandleLateResult. The zero value is valid.
type LateResultOptions struct {
	// Policy maps recommendation to action. Nil means DefaultLateResultPolicy.
	Policy LateResultPolicy
	// ResolveUserID returns the provider user id for the assessment's account
	// id, such as from your user store. It gets the context HandleLateResult
	// was called with. Return "" to take no action, or an error to fail the
	// delivery: HandleLateResult returns it unchanged and calls no provider,
	// so answer the webhook with a 5xx and NoResponse retries it. When nil,
	// UserID is used; when both are nil, the account id itself, as when you
	// assess with the provider's user id.
	ResolveUserID func(ctx context.Context, accountID string, event *Event) (string, error)
	// UserID is the older form of ResolveUserID, with no context and no way
	// to report a failed lookup. Ignored when ResolveUserID is set.
	//
	// Deprecated: use ResolveUserID.
	UserID func(accountID string, event *Event) string
	// Webhook options for ConstructEvent, such as WithTolerance and WithClock.
	Webhook []WebhookOption
}

// LateResult is the outcome of HandleLateResult.
type LateResult struct {
	// Event is the verified event. Answer an endpoint.verification with
	// VerificationResponse(event).
	Event *Event
	// Action taken. LateResultNone when the policy, the event or the user id
	// called for nothing.
	Action LateResultAction
	// UserID is the provider user id acted on; empty when no action was taken.
	UserID string
}

// ProviderActionError is returned when the provider's admin API refuses a
// call. Answer the webhook with a 5xx so NoResponse retries it. A 404 is never
// returned: the user or session is already gone, so the action is done.
type ProviderActionError struct {
	Provider string
	Status   int
	Message  string
}

func (e *ProviderActionError) Error() string { return "noresponse: " + e.Message }

// LateResultActionFor is the action your policy names for a verified event.
// Only a served assessment with a recommendation can act; a test event from
// WebhookEndpoints.SendTestEvent never does. A nil policy means
// DefaultLateResultPolicy.
func LateResultActionFor(event *Event, policy LateResultPolicy) (LateResultAction, error) {
	if event == nil || event.Type != EventAssessmentUpdated || event.Test {
		return LateResultNone, nil
	}
	assessment, ok := event.Data.(*components.Assessment)
	if !ok || assessment == nil {
		return LateResultNone, nil
	}
	if policy == nil {
		policy = DefaultLateResultPolicy()
	}
	var recommendation components.Recommendation
	switch {
	case assessment.Recommendation != nil:
		recommendation = *assessment.Recommendation
	case assessment.Risk != nil:
		recommendation = assessment.Risk.Recommendation
	}
	switch recommendation {
	case components.RecommendationAllow, components.RecommendationStepUp, components.RecommendationReview:
	default:
		return LateResultNone, nil
	}
	action, ok := policy[recommendation]
	if !ok || action == "" {
		return LateResultNone, nil
	}
	switch action {
	case LateResultNone, LateResultRevokeSessions, LateResultBan:
		return action, nil
	}
	return LateResultNone, fmt.Errorf("%w %q", ErrUnknownLateResultAction, string(action))
}

// HandleLateResult verifies a webhook delivery and applies your policy to it.
//
// payload, headers, secrets and options.Webhook are exactly what
// ConstructEvent takes; a delivery that fails verification returns the
// *WebhookError before any provider call. A test event acts on nothing.
// Revoke and ban are safe to repeat, but deduplicate on result.Event.ID and
// ignore versions older than one you have acted on. options may be nil.
func HandleLateResult(ctx context.Context, payload []byte, headers http.Header, secrets map[string]string, provider LateResultProvider, options *LateResultOptions) (*LateResult, error) {
	if options == nil {
		options = &LateResultOptions{}
	}
	event, err := ConstructEvent(payload, headers, secrets, options.Webhook...)
	if err != nil {
		return nil, err
	}
	action, err := LateResultActionFor(event, options.Policy)
	if err != nil {
		return nil, err
	}
	assessment, ok := event.Data.(*components.Assessment)
	if action == LateResultNone || !ok {
		return &LateResult{Event: event, Action: LateResultNone}, nil
	}
	userID := assessment.AccountID
	switch {
	case options.ResolveUserID != nil:
		if userID, err = options.ResolveUserID(ctx, assessment.AccountID, event); err != nil {
			return nil, err
		}
	case options.UserID != nil:
		userID = options.UserID(assessment.AccountID, event)
	}
	if userID == "" {
		return &LateResult{Event: event, Action: LateResultNone}, nil
	}
	if provider == nil {
		return nil, errors.New("noresponse: HandleLateResult needs a provider")
	}
	if action == LateResultBan {
		err = provider.Ban(ctx, userID)
	} else {
		err = provider.RevokeSessions(ctx, userID)
	}
	if err != nil {
		return nil, err
	}
	return &LateResult{Event: event, Action: action, UserID: userID}, nil
}

type adminCall struct {
	method string
	path   string
	query  url.Values
	body   any
}

type adminCaller struct {
	provider      string
	base          string
	authorization string
	client        *http.Client
}

func newAdminCaller(provider, base, token string, client *http.Client) *adminCaller {
	if client == nil {
		client = http.DefaultClient
	}
	return &adminCaller{provider: provider, base: strings.TrimRight(base, "/"), authorization: "Bearer " + token, client: client}
}

// do sends one request and returns the response body, nil when empty or when
// a 404 answered. Every call names a user or a session, so a 404 means it is
// already gone: revoke and ban count it as done and a repeated delivery does
// not fail.
func (c *adminCaller) do(ctx context.Context, call adminCall) ([]byte, error) {
	target := c.base + call.path
	if len(call.query) > 0 {
		target += "?" + call.query.Encode()
	}
	var reader io.Reader
	if call.body != nil {
		encoded, err := json.Marshal(call.body)
		if err != nil {
			return nil, err
		}
		reader = strings.NewReader(string(encoded))
	}
	request, err := http.NewRequestWithContext(ctx, call.method, target, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", c.authorization)
	request.Header.Set("Accept", "application/json")
	if call.body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	text, err := io.ReadAll(io.LimitReader(response.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, &ProviderActionError{
			Provider: c.provider,
			Status:   response.StatusCode,
			Message:  fmt.Sprintf("%s %s %s answered %d", c.provider, call.method, call.path, response.StatusCode),
		}
	}
	return text, nil
}

func pathSegment(id string) (string, error) {
	if id == "" {
		return "", ErrEmptyUserID
	}
	return url.PathEscape(id), nil
}

type providerSession struct {
	ID     any `json:"id"`
	Status any `json:"status"`
}

// activeSessionIDs reads a bare JSON array or {"data": [...]} and returns the
// ids of the sessions whose status is "active", with the page length.
func activeSessionIDs(body []byte) (ids []string, count int, listAfter string) {
	var list []providerSession
	if json.Unmarshal(body, &list) != nil {
		var wrapped struct {
			Data         []providerSession `json:"data"`
			ListMetadata struct {
				After any `json:"after"`
			} `json:"list_metadata"`
		}
		if json.Unmarshal(body, &wrapped) == nil {
			list = wrapped.Data
			if after, ok := wrapped.ListMetadata.After.(string); ok {
				listAfter = after
			}
		}
	}
	for _, session := range list {
		id, _ := session.ID.(string)
		status, _ := session.Status.(string)
		if status == "active" && id != "" {
			ids = append(ids, id)
		}
	}
	return ids, len(list), listAfter
}

const adminPageSize = 100

// ClerkActions acts through the Clerk Backend API. RevokeSessions lists the
// user's active sessions (GET /sessions?user_id&status=active) and revokes
// each (POST /sessions/{id}/revoke). Ban is POST /users/{id}/ban, which
// revokes every session and blocks sign-in.
type ClerkActions struct {
	// SecretKey is the Clerk secret key (sk_...).
	SecretKey string
	// BaseURL defaults to https://api.clerk.com/v1.
	BaseURL string
	// HTTPClient defaults to http.DefaultClient.
	HTTPClient *http.Client
}

func (a *ClerkActions) Name() string { return "clerk" }

func (a *ClerkActions) caller() (*adminCaller, error) {
	if a == nil || a.SecretKey == "" {
		return nil, errors.New("noresponse: ClerkActions needs SecretKey")
	}
	base := a.BaseURL
	if base == "" {
		base = "https://api.clerk.com/v1"
	}
	return newAdminCaller("clerk", base, a.SecretKey, a.HTTPClient), nil
}

func (a *ClerkActions) RevokeSessions(ctx context.Context, userID string) error {
	if _, err := pathSegment(userID); err != nil {
		return err
	}
	c, err := a.caller()
	if err != nil {
		return err
	}
	var active []string
	for offset := 0; offset < adminPageSize*1000; offset += adminPageSize {
		body, err := c.do(ctx, adminCall{method: http.MethodGet, path: "/sessions", query: url.Values{
			"user_id": {userID}, "status": {"active"}, "limit": {strconv.Itoa(adminPageSize)}, "offset": {strconv.Itoa(offset)},
		}})
		if err != nil {
			return err
		}
		ids, count, _ := activeSessionIDs(body)
		active = append(active, ids...)
		if count < adminPageSize {
			break
		}
	}
	for _, id := range active {
		if _, err := c.do(ctx, adminCall{method: http.MethodPost, path: "/sessions/" + url.PathEscape(id) + "/revoke"}); err != nil {
			return err
		}
	}
	return nil
}

func (a *ClerkActions) Ban(ctx context.Context, userID string) error {
	user, err := pathSegment(userID)
	if err != nil {
		return err
	}
	c, err := a.caller()
	if err != nil {
		return err
	}
	_, err = c.do(ctx, adminCall{method: http.MethodPost, path: "/users/" + user + "/ban"})
	return err
}

// WorkOSActions acts through WorkOS User Management. RevokeSessions pages
// through GET /user_management/users/{id}/sessions and revokes each active
// one with POST /user_management/sessions/revoke. WorkOS has no block or ban,
// so Ban revokes the sessions and then deletes the user
// (DELETE /user_management/users/{id}). Deletion is permanent.
type WorkOSActions struct {
	// APIKey is the WorkOS API key (sk_...).
	APIKey string
	// BaseURL defaults to https://api.workos.com.
	BaseURL string
	// HTTPClient defaults to http.DefaultClient.
	HTTPClient *http.Client
}

func (a *WorkOSActions) Name() string { return "workos" }

func (a *WorkOSActions) caller() (*adminCaller, error) {
	if a == nil || a.APIKey == "" {
		return nil, errors.New("noresponse: WorkOSActions needs APIKey")
	}
	base := a.BaseURL
	if base == "" {
		base = "https://api.workos.com"
	}
	return newAdminCaller("workos", base, a.APIKey, a.HTTPClient), nil
}

func (a *WorkOSActions) RevokeSessions(ctx context.Context, userID string) error {
	user, err := pathSegment(userID)
	if err != nil {
		return err
	}
	c, err := a.caller()
	if err != nil {
		return err
	}
	var active []string
	after := ""
	for round := 0; round < 1000; round++ {
		query := url.Values{"limit": {"100"}}
		if after != "" {
			query.Set("after", after)
		}
		body, err := c.do(ctx, adminCall{method: http.MethodGet, path: "/user_management/users/" + user + "/sessions", query: query})
		if err != nil {
			return err
		}
		ids, _, next := activeSessionIDs(body)
		active = append(active, ids...)
		if next == "" {
			break
		}
		after = next
	}
	for _, id := range active {
		if _, err := c.do(ctx, adminCall{method: http.MethodPost, path: "/user_management/sessions/revoke", body: map[string]string{"session_id": id}}); err != nil {
			return err
		}
	}
	return nil
}

func (a *WorkOSActions) Ban(ctx context.Context, userID string) error {
	user, err := pathSegment(userID)
	if err != nil {
		return err
	}
	if err := a.RevokeSessions(ctx, userID); err != nil {
		return err
	}
	c, err := a.caller()
	if err != nil {
		return err
	}
	_, err = c.do(ctx, adminCall{method: http.MethodDelete, path: "/user_management/users/" + user})
	return err
}

// Auth0Actions acts through the Auth0 Management API. RevokeSessions deletes
// every session (DELETE /users/{id}/sessions) and refresh token
// (DELETE /users/{id}/refresh-tokens). Ban first sets blocked: true
// (PATCH /users/{id}), then does the same.
type Auth0Actions struct {
	// Domain is your Auth0 tenant domain, such as example.eu.auth0.com. The
	// Management API is https://{Domain}/api/v2.
	Domain string
	// BaseURL is the Management API base URL, in place of Domain.
	BaseURL string
	// Token is a Management API access token with update:users,
	// delete:sessions and delete:refresh_tokens.
	Token string
	// HTTPClient defaults to http.DefaultClient.
	HTTPClient *http.Client
}

func (a *Auth0Actions) Name() string { return "auth0" }

func (a *Auth0Actions) caller() (*adminCaller, error) {
	if a == nil || a.Token == "" {
		return nil, errors.New("noresponse: Auth0Actions needs Token")
	}
	base := a.BaseURL
	if base == "" && a.Domain != "" {
		base = "https://" + a.Domain + "/api/v2"
	}
	if base == "" {
		return nil, errors.New("noresponse: Auth0Actions needs Domain or BaseURL")
	}
	return newAdminCaller("auth0", base, a.Token, a.HTTPClient), nil
}

func (a *Auth0Actions) RevokeSessions(ctx context.Context, userID string) error {
	user, err := pathSegment(userID)
	if err != nil {
		return err
	}
	c, err := a.caller()
	if err != nil {
		return err
	}
	if _, err := c.do(ctx, adminCall{method: http.MethodDelete, path: "/users/" + user + "/sessions"}); err != nil {
		return err
	}
	_, err = c.do(ctx, adminCall{method: http.MethodDelete, path: "/users/" + user + "/refresh-tokens"})
	return err
}

func (a *Auth0Actions) Ban(ctx context.Context, userID string) error {
	user, err := pathSegment(userID)
	if err != nil {
		return err
	}
	c, err := a.caller()
	if err != nil {
		return err
	}
	if _, err := c.do(ctx, adminCall{method: http.MethodPatch, path: "/users/" + user, body: map[string]bool{"blocked": true}}); err != nil {
		return err
	}
	return a.RevokeSessions(ctx, userID)
}

// LogOnlyActions is a provider for an app with its own login, or for trying
// the handler before you wire up a real one. RevokeSessions and Ban each log
// one line saying what they would do and call nothing: no request leaves your
// process. Replace it with your own LateResultProvider once you know what
// revoking and banning mean in your app. The zero value logs to
// slog.Default().
type LogOnlyActions struct {
	// Logger receives one Info record per action, with "action" and "user_id"
	// attributes. Nil means slog.Default().
	Logger *slog.Logger
}

func (a *LogOnlyActions) Name() string { return "log_only" }

func (a *LogOnlyActions) log(ctx context.Context, action LateResultAction, userID, message string) error {
	if userID == "" {
		return ErrEmptyUserID
	}
	logger := slog.Default()
	if a != nil && a.Logger != nil {
		logger = a.Logger
	}
	logger.InfoContext(ctx, message, "action", string(action), "user_id", userID)
	return nil
}

// RevokeSessions logs that it would end every session of the user.
func (a *LogOnlyActions) RevokeSessions(ctx context.Context, userID string) error {
	return a.log(ctx, LateResultRevokeSessions, userID, "NoResponse late result (log only): would revoke every session of user "+userID)
}

// Ban logs that it would ban the user and end their sessions.
func (a *LogOnlyActions) Ban(ctx context.Context, userID string) error {
	return a.log(ctx, LateResultBan, userID, "NoResponse late result (log only): would ban user "+userID+" and end their sessions")
}
