package noresponse

// Late-result handlers. Hand-written beside the generated client.
//
// An assessment can be revised after the signup was allowed: a slow source
// settles and the account's decision rises. NoResponse sends the new version as a
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
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/noresponsehq/go/v3/models/components"
)

// LateResultAction is what to do to the user. LateResultBan also ends every session.
type LateResultAction string

const (
	LateResultNone           LateResultAction = "none"
	LateResultRevokeSessions LateResultAction = "revoke_sessions"
	LateResultBan            LateResultAction = "ban"
)

// LateResultPolicy maps each account decision (account.risk.decision) to an
// action. A decision left out does nothing.
type LateResultPolicy map[components.AccountDecision]LateResultAction

// DefaultLateResultPolicy revokes sessions on review and block and does
// nothing on allow or step_up. Banning is opt-in. It returns a fresh map each
// call.
func DefaultLateResultPolicy() LateResultPolicy {
	return LateResultPolicy{
		components.AccountDecisionAllow:  LateResultNone,
		components.AccountDecisionStepUp: LateResultNone,
		components.AccountDecisionReview: LateResultRevokeSessions,
		components.AccountDecisionBlock:  LateResultRevokeSessions,
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

// LateResultActor is a LateResultProvider that reports the action that ran.
// HandleLateResult calls Act when the provider implements it, and Ban or
// RevokeSessions otherwise. A provider that has to do more than it was asked
// implements it: Auth0Actions returns LateResultBan when the tenant's plan
// has no sessions API and it blocked the user to end their sessions.
type LateResultActor interface {
	LateResultProvider
	// Act takes action (LateResultRevokeSessions or LateResultBan) and returns
	// the action that ran.
	Act(ctx context.Context, action LateResultAction, userID string) (LateResultAction, error)
}

// LateResultOptions adjusts HandleLateResult. The zero value is valid.
type LateResultOptions struct {
	// Policy maps account decision to action. Nil means DefaultLateResultPolicy.
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
	// Action is the action that ran. LateResultNone when the policy, the
	// event or the user id called for nothing. It can be stronger than the
	// policy's action when the provider had to fall back: Auth0Actions reports
	// LateResultBan when it blocked the user because the tenant's plan cannot
	// end sessions.
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
	// Code is the provider's error code from the response body, such as
	// Auth0's errorCode "insufficient_scope", when it sent one.
	Code string
}

func (e *ProviderActionError) Error() string { return "noresponse: " + e.Message }

// LateResultActionFor is the action your policy names for a verified event.
// Only a decided assessment (complete or partial) that carries the account's
// decision can act; a test event from WebhookEndpoints.SendTestEvent never does. A nil policy
// means DefaultLateResultPolicy.
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
	// A pending assessment carries the decision of an earlier one, or none.
	if assessment.Status == components.StatusPending || assessment.Account == nil {
		return LateResultNone, nil
	}
	decision := assessment.Account.Risk.Decision
	switch decision {
	case components.AccountDecisionAllow, components.AccountDecisionStepUp, components.AccountDecisionReview, components.AccountDecisionBlock:
	default:
		return LateResultNone, nil
	}
	action, ok := policy[decision]
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
	userID := assessment.Detail.AccountID
	switch {
	case options.ResolveUserID != nil:
		if userID, err = options.ResolveUserID(ctx, assessment.Detail.AccountID, event); err != nil {
			return nil, err
		}
	case options.UserID != nil:
		userID = options.UserID(assessment.Detail.AccountID, event)
	}
	if userID == "" {
		return &LateResult{Event: event, Action: LateResultNone}, nil
	}
	if provider == nil {
		return nil, errors.New("noresponse: HandleLateResult needs a provider")
	}
	ran := action
	if actor, ok := provider.(LateResultActor); ok {
		ran, err = actor.Act(ctx, action, userID)
	} else if action == LateResultBan {
		err = provider.Ban(ctx, userID)
	} else {
		err = provider.RevokeSessions(ctx, userID)
	}
	if err != nil {
		return nil, err
	}
	return &LateResult{Event: event, Action: ran, UserID: userID}, nil
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
	// retryRateLimits waits out a 429 and sends the call again; waited is
	// the time spent waiting so far, in one action.
	retryRateLimits bool
	waited          time.Duration
}

// Rate limits. Auth0 answers 429 when the Management API rate limit is hit,
// and a 429 means it did not act, so the call is safe to send again, a DELETE
// included. One revoke or ban waits at most rateLimitBudget in all, and each
// call retries at most rateLimitRetries times: it waits Retry-After seconds,
// else until X-RateLimit-Reset (Unix seconds), else 250 ms, 500 ms, 1 s, 2 s.
// A wait that would pass the budget is not started: the 429 returns at once
// as a ProviderActionError, the webhook answers 5xx and NoResponse delivers
// it again later. The budget keeps the handler inside the 5 s a delivery has
// to be answered.
const (
	rateLimitRetries = 4
	rateLimitBudget  = 3 * time.Second
	rateLimitBackoff = 250 * time.Millisecond
)

// rateLimitWait is how long a 429 asks to wait, capped at an hour so a huge
// header cannot overflow; anything past the budget gives up anyway.
func rateLimitWait(header http.Header, retry int) time.Duration {
	seconds := func(name string) (float64, bool) {
		value, err := strconv.ParseFloat(strings.TrimSpace(header.Get(name)), 64)
		return value, err == nil && !math.IsInf(value, 0) && !math.IsNaN(value)
	}
	if after, ok := seconds("Retry-After"); ok && after >= 0 {
		return time.Duration(math.Min(after, 3600) * float64(time.Second))
	}
	if reset, ok := seconds("X-RateLimit-Reset"); ok {
		left := reset - float64(time.Now().UnixNano())/float64(time.Second)
		return time.Duration(math.Max(0, math.Min(left, 3600)) * float64(time.Second))
	}
	return rateLimitBackoff << retry
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
	var encoded []byte
	if call.body != nil {
		var err error
		if encoded, err = json.Marshal(call.body); err != nil {
			return nil, err
		}
	}
	for retry := 0; ; retry++ {
		status, header, text, err := c.send(ctx, call.method, target, encoded, call.body != nil)
		if err != nil {
			return nil, err
		}
		if status == http.StatusTooManyRequests && c.retryRateLimits && retry < rateLimitRetries {
			if wait := rateLimitWait(header, retry); c.waited+wait <= rateLimitBudget {
				c.waited += wait
				timer := time.NewTimer(wait)
				select {
				case <-ctx.Done():
					timer.Stop()
					return nil, ctx.Err()
				case <-timer.C:
				}
				continue
			}
		}
		return c.answer(call, status, text)
	}
}

// send makes one request and reads the whole answer.
func (c *adminCaller) send(ctx context.Context, method, target string, body []byte, hasBody bool) (int, http.Header, []byte, error) {
	var reader io.Reader
	if hasBody {
		reader = strings.NewReader(string(body))
	}
	request, err := http.NewRequestWithContext(ctx, method, target, reader)
	if err != nil {
		return 0, nil, nil, err
	}
	request.Header.Set("Authorization", c.authorization)
	request.Header.Set("Accept", "application/json")
	if hasBody {
		request.Header.Set("Content-Type", "application/json")
	}
	response, err := c.client.Do(request)
	if err != nil {
		return 0, nil, nil, err
	}
	defer response.Body.Close()
	text, err := io.ReadAll(io.LimitReader(response.Body, 4<<20))
	if err != nil {
		return 0, nil, nil, err
	}
	return response.StatusCode, response.Header, text, nil
}

// answer maps a response to the body or a ProviderActionError.
func (c *adminCaller) answer(call adminCall, status int, text []byte) ([]byte, error) {
	if status == http.StatusNotFound {
		return nil, nil
	}
	if status < 200 || status > 299 {
		return nil, &ProviderActionError{
			Provider: c.provider,
			Status:   status,
			Message:  fmt.Sprintf("%s %s %s answered %d", c.provider, call.method, call.path, status),
			Code:     errorCode(text),
		}
	}
	return text, nil
}

// errorCode is the error code in a provider's error body: Auth0's errorCode,
// or a top-level code.
func errorCode(body []byte) string {
	var parsed struct {
		ErrorCode any `json:"errorCode"`
		Code      any `json:"code"`
	}
	if json.Unmarshal(body, &parsed) != nil {
		return ""
	}
	if code, ok := parsed.ErrorCode.(string); ok && code != "" {
		return code
	}
	if code, ok := parsed.Code.(string); ok {
		return code
	}
	return ""
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
//
// Auth0 offers the sessions and refresh-token endpoints on Enterprise plans
// only; on other plans they answer 403 feature_not_enabled. Then refresh
// tokens are deleted one by one through /device-credentials (types
// refresh_token and rotating_refresh_token), which every plan has. A login
// session cannot be ended, so RevokeSessions blocks the user before deleting
// the refresh tokens, and Act reports LateResultBan. Unblock the user in the
// dashboard to undo it. Any other refusal, such as 403 insufficient_scope, is
// returned.
//
// A 429 from the Management API is waited out and retried: Retry-After, else
// X-RateLimit-Reset, else 250 ms doubling, at most 4 retries a call and 3 s
// of waiting a revoke or ban. Past that the 429 is returned as a
// ProviderActionError.
type Auth0Actions struct {
	// Domain is your Auth0 tenant domain, such as example.eu.auth0.com. The
	// Management API is https://{Domain}/api/v2.
	Domain string
	// BaseURL is the Management API base URL, in place of Domain.
	BaseURL string
	// Token is a Management API access token with update:users,
	// delete:sessions, delete:refresh_tokens, read:device_credentials and
	// delete:device_credentials.
	Token string
	// HTTPClient defaults to http.DefaultClient.
	HTTPClient *http.Client
}

var _ LateResultActor = (*Auth0Actions)(nil)

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
	// A new caller for each action, so each revoke or ban has its own rate-limit budget.
	c := newAdminCaller("auth0", base, a.Token, a.HTTPClient)
	c.retryRateLimits = true
	return c, nil
}

// auth0NotOnPlan reports Auth0's 403 feature_not_enabled, its answer when the
// tenant's plan does not include an endpoint.
func auth0NotOnPlan(err error) bool {
	var refused *ProviderActionError
	return errors.As(err, &refused) && refused.Status == http.StatusForbidden && refused.Code == "feature_not_enabled"
}

func auth0Block(ctx context.Context, c *adminCaller, user string) error {
	_, err := c.do(ctx, adminCall{method: http.MethodPatch, path: "/users/" + user, body: map[string]bool{"blocked": true}})
	return err
}

// auth0DeleteSessions ends every session; false when the plan has no sessions API.
func auth0DeleteSessions(ctx context.Context, c *adminCaller, user string) (bool, error) {
	_, err := c.do(ctx, adminCall{method: http.MethodDelete, path: "/users/" + user + "/sessions"})
	if auth0NotOnPlan(err) {
		return false, nil
	}
	return err == nil, err
}

func auth0DeleteRefreshTokens(ctx context.Context, c *adminCaller, userID, user string) error {
	_, err := c.do(ctx, adminCall{method: http.MethodDelete, path: "/users/" + user + "/refresh-tokens"})
	if !auth0NotOnPlan(err) {
		return err
	}
	// Collect every id before deleting any, so the pages do not shift underneath.
	const page = 100
	var ids []string
	for _, kind := range []string{"refresh_token", "rotating_refresh_token"} {
		for n := 0; n < 10; n++ {
			query := url.Values{"user_id": {userID}, "type": {kind}, "page": {strconv.Itoa(n)}, "per_page": {strconv.Itoa(page)}}
			body, err := c.do(ctx, adminCall{method: http.MethodGet, path: "/device-credentials", query: query})
			if err != nil {
				return err
			}
			var listed []struct {
				ID any `json:"id"`
			}
			if len(body) > 0 && json.Unmarshal(body, &listed) != nil {
				listed = nil
			}
			for _, credential := range listed {
				if id, ok := credential.ID.(string); ok && id != "" {
					ids = append(ids, id)
				}
			}
			if len(listed) < page {
				break
			}
		}
	}
	for _, id := range ids {
		if _, err := c.do(ctx, adminCall{method: http.MethodDelete, path: "/device-credentials/" + url.PathEscape(id)}); err != nil {
			return err
		}
	}
	return nil
}

func (a *Auth0Actions) revoke(ctx context.Context, userID string) (LateResultAction, error) {
	user, err := pathSegment(userID)
	if err != nil {
		return LateResultNone, err
	}
	c, err := a.caller()
	if err != nil {
		return LateResultNone, err
	}
	ended, err := auth0DeleteSessions(ctx, c, user)
	if err != nil {
		return LateResultNone, err
	}
	// Without the sessions API a login session survives; blocking is the only way to end it.
	if !ended {
		if err := auth0Block(ctx, c, user); err != nil {
			return LateResultNone, err
		}
	}
	if err := auth0DeleteRefreshTokens(ctx, c, userID, user); err != nil {
		return LateResultNone, err
	}
	if ended {
		return LateResultRevokeSessions, nil
	}
	return LateResultBan, nil
}

func (a *Auth0Actions) RevokeSessions(ctx context.Context, userID string) error {
	_, err := a.revoke(ctx, userID)
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
	if err := auth0Block(ctx, c, user); err != nil {
		return err
	}
	// The block already ends a session the plan cannot delete.
	if _, err := auth0DeleteSessions(ctx, c, user); err != nil {
		return err
	}
	return auth0DeleteRefreshTokens(ctx, c, userID, user)
}

// Act bans or revokes and returns the action that ran: LateResultBan when
// RevokeSessions had to block the user.
func (a *Auth0Actions) Act(ctx context.Context, action LateResultAction, userID string) (LateResultAction, error) {
	if action == LateResultBan {
		if err := a.Ban(ctx, userID); err != nil {
			return LateResultNone, err
		}
		return LateResultBan, nil
	}
	return a.revoke(ctx, userID)
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
