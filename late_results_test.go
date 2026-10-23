package noresponse

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/noresponsehq/go/v4/models/components"
)

const (
	lateSecret = "late-result-secret"
	lateKey    = "whk_sdktest01"
	lateNow    = int64(1791968402)
)

var lateSecrets = map[string]string{lateKey: lateSecret}

func lateFixture(t *testing.T, name string) map[string]any {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", name))
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}

// lateView wraps the assessment detail fixture name in the assessment
// envelope, with the account's decision set to the detail's recommendation.
// A pending detail has no decision, so its envelope has no account, as for an
// account's first assessment.
func lateView(t *testing.T, name string) map[string]any {
	t.Helper()
	detail := lateFixture(t, name)
	view := lateFixture(t, "assessment-view-complete.json")
	view["id"] = detail["assessment_id"]
	view["status"] = detail["status"]
	view["detail"] = detail
	recommendation, ok := detail["recommendation"].(string)
	if !ok {
		delete(view, "account")
		return view
	}
	account := view["account"].(map[string]any)
	account["assessment_id"] = detail["assessment_id"]
	account["risk"].(map[string]any)["decision"] = recommendation
	return view
}

// lateDelivery signs an assessment.updated delivery of data, overriding detail.account_id when given.
func lateDelivery(t *testing.T, data map[string]any, accountID string) ([]byte, http.Header) {
	t.Helper()
	return lateDeliveryWith(t, data, accountID, nil)
}

// lateDeliveryWith is lateDelivery with extra top-level body fields, such as "test": true.
func lateDeliveryWith(t *testing.T, data map[string]any, accountID string, extra map[string]any) ([]byte, http.Header) {
	t.Helper()
	copied := map[string]any{}
	for k, v := range data {
		copied[k] = v
	}
	if accountID != "" {
		detail := map[string]any{}
		for k, v := range copied["detail"].(map[string]any) {
			detail[k] = v
		}
		detail["account_id"] = accountID
		copied["detail"] = detail
	}
	id := "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2"
	body := map[string]any{"id": id, "type": "assessment.updated", "created_at": "2026-10-18T11:00:04Z", "data": copied}
	for k, v := range extra {
		body[k] = v
	}
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	stamp := fmt.Sprint(lateNow)
	mac := hmac.New(sha256.New, []byte(lateSecret))
	mac.Write([]byte(stamp + "."))
	mac.Write(raw)
	h := http.Header{}
	h.Set("NR-Webhook-Id", id)
	h.Set("NR-Webhook-Timestamp", stamp)
	h.Set("NR-Webhook-Key-Id", lateKey)
	h.Set("NR-Webhook-Signature", "v1="+hex.EncodeToString(mac.Sum(nil)))
	return raw, h
}

type mockCall struct {
	Method string
	Path   string // escaped, as sent
	Query  map[string]string
	Auth   string
	Accept string
	CType  string
	Body   map[string]any
}

type mockProvider struct {
	URL   string
	mu    sync.Mutex
	calls []mockCall
}

func (m *mockProvider) reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = nil
}

func (m *mockProvider) summary() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]string, len(m.calls))
	for i, c := range m.calls {
		out[i] = c.Method + " " + c.Path
	}
	return out
}

func (m *mockProvider) snapshot() []mockCall {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]mockCall(nil), m.calls...)
}

// startMock runs a real HTTP server on 127.0.0.1 standing in for a provider's
// admin API. It records each request and answers from route; a nil body with
// status 0 means 404.
func startMock(t *testing.T, route func(c mockCall) (int, any)) *mockProvider {
	t.Helper()
	m := &mockProvider{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		text, _ := io.ReadAll(r.Body)
		call := mockCall{
			Method: r.Method,
			Path:   r.URL.EscapedPath(),
			Query:  map[string]string{},
			Auth:   r.Header.Get("Authorization"),
			Accept: r.Header.Get("Accept"),
			CType:  r.Header.Get("Content-Type"),
		}
		for k, v := range r.URL.Query() {
			call.Query[k] = v[0]
		}
		if len(text) > 0 {
			_ = json.Unmarshal(text, &call.Body)
		}
		m.mu.Lock()
		m.calls = append(m.calls, call)
		m.mu.Unlock()
		status, body := route(call)
		if status == 0 {
			status, body = 404, map[string]string{"error": "not found"}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
	t.Cleanup(server.Close)
	m.URL = server.URL
	return m
}

func lateOptions(extra *LateResultOptions) *LateResultOptions {
	o := LateResultOptions{}
	if extra != nil {
		o = *extra
	}
	o.Webhook = append(o.Webhook, at(lateNow))
	return &o
}

func assertCalls(t *testing.T, m *mockProvider, want []string) {
	t.Helper()
	if got := m.summary(); !reflect.DeepEqual(got, want) {
		t.Fatalf("calls\n got %q\nwant %q", got, want)
	}
}

func assertEveryCall(t *testing.T, m *mockProvider, auth string) {
	t.Helper()
	for _, c := range m.snapshot() {
		if c.Auth != auth {
			t.Fatalf("authorization %q, want %q (%s %s)", c.Auth, auth, c.Method, c.Path)
		}
		if c.Accept != "application/json" {
			t.Fatalf("accept %q on %s %s", c.Accept, c.Method, c.Path)
		}
		if (c.Body != nil) != (c.CType == "application/json") {
			t.Fatalf("content-type %q with body %v on %s %s", c.CType, c.Body, c.Method, c.Path)
		}
	}
}

func lateEvent(status string, decision string) *Event {
	data := &components.Assessment{Status: components.Status(status)}
	if decision != "" {
		data.Account = &components.Account{Risk: components.AccountRisk{Decision: components.AccountDecision(decision)}}
	}
	return &Event{ID: "evt_x", Type: EventAssessmentUpdated, CreatedAt: "2026-10-18T11:00:04Z", Data: data}
}

func TestDefaultPolicyRevokesOnReviewAndBlock(t *testing.T) {
	expect := func(event *Event, policy LateResultPolicy, want LateResultAction) {
		t.Helper()
		got, err := LateResultActionFor(event, policy)
		if err != nil || got != want {
			t.Fatalf("got %q, %v; want %q", got, err, want)
		}
	}
	expect(lateEvent("complete", "review"), nil, LateResultRevokeSessions)
	expect(lateEvent("complete", "block"), nil, LateResultRevokeSessions)
	expect(lateEvent("complete", "step_up"), nil, LateResultNone)
	expect(lateEvent("pending", "review"), nil, LateResultNone)
	expect(lateEvent("partial", "allow"), nil, LateResultNone)
	expect(lateEvent("pending", ""), nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventAssessmentUpdated, Data: &AssessmentExpired{Status: "expired", ChangeReason: "expired"}}, nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventEndpointVerification, Data: &EndpointVerification{}}, nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventAccountDeleted, Data: &AccountDeleted{AccountID: "acct_1", AssessmentsDeleted: 1}}, nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventTenantDeleted, Data: &TenantDeleted{WorkspaceID: "tnt_1", Environment: "live"}}, nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventEndpointDisabled, Data: &EndpointDisabled{EndpointID: "whe_1"}}, nil, LateResultNone)
	blocked := components.Account{ID: "acct_1", Risk: components.AccountRisk{Score: 90, Decision: components.AccountDecisionBlock}}
	expect(&Event{ID: "evt_x", Type: EventAccountRiskChanged, Data: &AccountRiskChanged{Account: blocked}}, nil, LateResultNone)
	expect(lateEvent("complete", "step_up"), LateResultPolicy{components.AccountDecisionStepUp: LateResultBan}, LateResultBan)
	expect(lateEvent("complete", "review"), LateResultPolicy{components.AccountDecisionStepUp: LateResultBan}, LateResultNone)
	expect(lateEvent("complete", "block"), LateResultPolicy{components.AccountDecisionBlock: LateResultBan}, LateResultBan)

	// Only the account's decision counts: the detail's recommendation alone acts on nothing.
	detailOnly := &Event{Type: EventAssessmentUpdated, Data: &components.Assessment{Status: "complete", Detail: components.AssessmentDetail{Recommendation: components.RecommendationReview.ToPointer()}}}
	expect(detailOnly, nil, LateResultNone)

	if _, err := LateResultActionFor(lateEvent("complete", "review"), LateResultPolicy{components.AccountDecisionReview: "delete"}); !errors.Is(err, ErrUnknownLateResultAction) {
		t.Fatalf("unknown action error = %v", err)
	}
	want := LateResultPolicy{"allow": "none", "step_up": "none", "review": "revoke_sessions", "block": "revoke_sessions"}
	if got := DefaultLateResultPolicy(); !reflect.DeepEqual(got, want) {
		t.Fatalf("default policy = %v", got)
	}
}

func TestClerkReviewRevokesActiveSessionsAndBanPolicyBans(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) {
		switch {
		case c.Method == "GET" && c.Path == "/v1/sessions":
			if c.Query["offset"] == "0" {
				return 200, []map[string]string{{"id": "sess_1", "status": "active"}, {"id": "sess_2", "status": "active"}}
			}
			return 200, []map[string]string{}
		case c.Method == "POST" && strings.HasPrefix(c.Path, "/v1/sessions/sess_") && strings.HasSuffix(c.Path, "/revoke"):
			return 200, map[string]string{"status": "revoked"}
		case c.Method == "POST" && c.Path == "/v1/users/user_2EXAMPLE/ban":
			return 200, map[string]any{"id": "user_2EXAMPLE", "banned": true}
		}
		return 0, nil
	})
	clerk := &ClerkActions{SecretKey: "sk_test_mock", BaseURL: mock.URL + "/v1"}
	review := lateView(t, "assessment-phone-uncovered-complete.json")
	raw, headers := lateDelivery(t, review, "user_2EXAMPLE")
	ctx := context.Background()

	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, clerk, lateOptions(nil))
	if err != nil {
		t.Fatal(err)
	}
	if result.Action != LateResultRevokeSessions || result.UserID != "user_2EXAMPLE" {
		t.Fatalf("result = %+v", result)
	}
	if assessment, ok := result.Event.Data.(*components.Assessment); !ok || assessment.ID != review["id"] {
		t.Fatalf("event data = %+v", result.Event.Data)
	}
	assertCalls(t, mock, []string{"GET /v1/sessions", "POST /v1/sessions/sess_1/revoke", "POST /v1/sessions/sess_2/revoke"})
	if got, want := mock.snapshot()[0].Query, map[string]string{"user_id": "user_2EXAMPLE", "status": "active", "limit": "100", "offset": "0"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("query = %v", got)
	}
	assertEveryCall(t, mock, "Bearer sk_test_mock")

	mock.reset()
	banned, err := HandleLateResult(ctx, raw, headers, lateSecrets, clerk, lateOptions(&LateResultOptions{Policy: LateResultPolicy{components.AccountDecisionReview: LateResultBan}}))
	if err != nil {
		t.Fatal(err)
	}
	if banned.Action != LateResultBan {
		t.Fatalf("action = %q", banned.Action)
	}
	assertCalls(t, mock, []string{"POST /v1/users/user_2EXAMPLE/ban"})
	assertEveryCall(t, mock, "Bearer sk_test_mock")
}

func TestClerkPagesUntilAShortPageAndAcceptsDataWrapper(t *testing.T) {
	page := func(prefix string, n int) []map[string]string {
		out := make([]map[string]string, n)
		for i := range out {
			out[i] = map[string]string{"id": fmt.Sprintf("%s%d", prefix, i), "status": "active"}
		}
		return out
	}
	mock := startMock(t, func(c mockCall) (int, any) {
		if c.Method == "GET" {
			switch c.Query["offset"] {
			case "0":
				return 200, map[string]any{"data": page("a", 100)}
			case "100":
				return 200, page("b", 1)
			}
		}
		if c.Method == "POST" {
			return 200, nil
		}
		return 0, nil
	})
	clerk := &ClerkActions{SecretKey: "sk", BaseURL: mock.URL + "/v1"}
	if err := clerk.RevokeSessions(context.Background(), "u1"); err != nil {
		t.Fatal(err)
	}
	calls := mock.summary()
	if len(calls) != 2+101 || calls[0] != "GET /v1/sessions" || calls[1] != "GET /v1/sessions" || calls[2] != "POST /v1/sessions/a0/revoke" || calls[102] != "POST /v1/sessions/b0/revoke" {
		t.Fatalf("unexpected call sequence (%d calls): %q ... %q", len(calls), calls[:3], calls[len(calls)-1])
	}
	if mock.snapshot()[1].Query["offset"] != "100" {
		t.Fatalf("second page offset = %q", mock.snapshot()[1].Query["offset"])
	}
}

func TestWorkOSRevokePagesThroughSessionsAndBanRevokesThenDeletes(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) {
		switch {
		case c.Method == "GET" && c.Path == "/user_management/users/user_01WORKOS/sessions":
			if _, more := c.Query["after"]; !more {
				return 200, map[string]any{"object": "list", "data": []map[string]string{{"id": "session_A", "status": "active"}, {"id": "session_X", "status": "revoked"}}, "list_metadata": map[string]any{"before": nil, "after": "session_X"}}
			}
			return 200, map[string]any{"object": "list", "data": []map[string]string{{"id": "session_B", "status": "active"}}, "list_metadata": map[string]any{"before": "session_B", "after": nil}}
		case c.Method == "POST" && c.Path == "/user_management/sessions/revoke":
			return 200, map[string]any{"id": c.Body["session_id"], "status": "revoked"}
		case c.Method == "DELETE" && c.Path == "/user_management/users/user_01WORKOS":
			return 202, nil
		}
		return 0, nil
	})
	workos := &WorkOSActions{APIKey: "sk_test_workos", BaseURL: mock.URL}
	step := lateView(t, "assessment-complete.json")
	raw, headers := lateDelivery(t, step, "acct_7f3a9c")
	result, err := HandleLateResult(context.Background(), raw, headers, lateSecrets, workos, lateOptions(&LateResultOptions{
		Policy: LateResultPolicy{components.AccountDecisionStepUp: LateResultBan},
		UserID: func(accountID string, _ *Event) string {
			if accountID == "acct_7f3a9c" {
				return "user_01WORKOS"
			}
			return ""
		},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if result.Action != LateResultBan || result.UserID != "user_01WORKOS" {
		t.Fatalf("result = %+v", result)
	}
	calls := mock.snapshot()
	got := make([]string, len(calls))
	for i, c := range calls {
		detail := fmt.Sprint(c.Query)
		if c.Body != nil {
			detail = fmt.Sprint(c.Body)
		} else if c.Method == "DELETE" {
			detail = "map[]"
		}
		got[i] = c.Method + " " + c.Path + " " + detail
	}
	want := []string{
		"GET /user_management/users/user_01WORKOS/sessions map[limit:100]",
		"GET /user_management/users/user_01WORKOS/sessions map[after:session_X limit:100]",
		"POST /user_management/sessions/revoke map[session_id:session_A]",
		"POST /user_management/sessions/revoke map[session_id:session_B]",
		"DELETE /user_management/users/user_01WORKOS map[]",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("calls\n got %q\nwant %q", got, want)
	}
	assertEveryCall(t, mock, "Bearer sk_test_workos")
}

func TestWorkOSDeleteOf404CountsAsDone(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) {
		if c.Method == "GET" {
			return 200, map[string]any{"data": []any{}, "list_metadata": map[string]any{"after": nil}}
		}
		return 0, nil // DELETE answers 404
	})
	workos := &WorkOSActions{APIKey: "k", BaseURL: mock.URL}
	if err := workos.Ban(context.Background(), "user_gone"); err != nil {
		t.Fatalf("ban = %v", err)
	}
	assertCalls(t, mock, []string{"GET /user_management/users/user_gone/sessions", "DELETE /user_management/users/user_gone"})
}

func TestAuth0ReviewDeletesSessionsAndRefreshTokensBanBlocksFirstPipeIsEscaped(t *testing.T) {
	id := "auth0|64abc0de1234567890abcdef"
	escaped := "auth0%7C64abc0de1234567890abcdef"
	mock := startMock(t, func(c mockCall) (int, any) {
		switch {
		case c.Path == "/api/v2/users/"+escaped+"/sessions" && c.Method == "DELETE":
			return 202, nil
		case c.Path == "/api/v2/users/"+escaped+"/refresh-tokens" && c.Method == "DELETE":
			return 202, nil
		case c.Path == "/api/v2/users/"+escaped && c.Method == "PATCH":
			return 200, map[string]any{"user_id": id, "blocked": c.Body["blocked"]}
		}
		return 0, nil
	})
	auth0 := &Auth0Actions{BaseURL: mock.URL + "/api/v2", Token: "mgmt_token"}
	review := lateView(t, "assessment-phone-uncovered-complete.json")
	raw, headers := lateDelivery(t, review, id)
	ctx := context.Background()

	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(nil))
	if err != nil {
		t.Fatal(err)
	}
	if result.Action != LateResultRevokeSessions || result.UserID != id {
		t.Fatalf("result = %+v", result)
	}
	assertCalls(t, mock, []string{"DELETE /api/v2/users/" + escaped + "/sessions", "DELETE /api/v2/users/" + escaped + "/refresh-tokens"})

	mock.reset()
	if _, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(&LateResultOptions{Policy: LateResultPolicy{components.AccountDecisionReview: LateResultBan}})); err != nil {
		t.Fatal(err)
	}
	assertCalls(t, mock, []string{"PATCH /api/v2/users/" + escaped, "DELETE /api/v2/users/" + escaped + "/sessions", "DELETE /api/v2/users/" + escaped + "/refresh-tokens"})
	if body := mock.snapshot()[0].Body; !reflect.DeepEqual(body, map[string]any{"blocked": true}) {
		t.Fatalf("patch body = %v", body)
	}
	assertEveryCall(t, mock, "Bearer mgmt_token")
}

var auth0FeatureNotEnabled = map[string]any{"statusCode": 403, "error": "Forbidden", "message": "This feature is not enabled for this tenant.", "errorCode": "feature_not_enabled"}

type auth0Credential struct{ id, kind string }

// auth0WithoutSessionsAPI is an Auth0 tenant on a plan without the sessions
// API, as a Free tenant answers: both bulk endpoints refuse with 403
// feature_not_enabled, while /device-credentials and PATCH /users/{id} work.
// It keeps the refresh tokens and the blocked flag, so a repeated delivery
// sees what the first left.
type auth0WithoutSessionsAPI struct {
	mu               sync.Mutex
	escaped          string
	refreshTokensAPI bool
	blocked          bool
	credentials      []auth0Credential
}

func newAuth0WithoutSessionsAPI(escaped string, refreshTokensAPI bool) *auth0WithoutSessionsAPI {
	return &auth0WithoutSessionsAPI{escaped: escaped, refreshTokensAPI: refreshTokensAPI, credentials: []auth0Credential{{"dcr_a", "refresh_token"}, {"dcr_b", "refresh_token"}, {"dcr_r", "rotating_refresh_token"}}}
}

func (a *auth0WithoutSessionsAPI) state() (bool, int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.blocked, len(a.credentials)
}

func (a *auth0WithoutSessionsAPI) set(blocked bool, add ...auth0Credential) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.blocked = blocked
	a.credentials = append(a.credentials, add...)
}

func (a *auth0WithoutSessionsAPI) route(c mockCall) (int, any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	user := "/api/v2/users/" + a.escaped
	switch {
	case c.Method == "DELETE" && c.Path == user+"/sessions":
		return 403, auth0FeatureNotEnabled
	case c.Method == "DELETE" && c.Path == user+"/refresh-tokens":
		if !a.refreshTokensAPI {
			return 403, auth0FeatureNotEnabled
		}
		a.credentials = nil
		return 202, nil
	case c.Method == "PATCH" && c.Path == user:
		a.blocked = c.Body["blocked"] == true
		return 200, map[string]any{"blocked": a.blocked}
	case c.Method == "GET" && c.Path == "/api/v2/device-credentials":
		listed := []map[string]string{}
		for _, d := range a.credentials {
			if d.kind == c.Query["type"] {
				listed = append(listed, map[string]string{"id": d.id, "type": d.kind, "user_id": c.Query["user_id"]})
			}
		}
		return 200, listed
	case c.Method == "DELETE" && strings.HasPrefix(c.Path, "/api/v2/device-credentials/"):
		id := strings.TrimPrefix(c.Path, "/api/v2/device-credentials/")
		for i, d := range a.credentials {
			if d.id == id {
				a.credentials = append(a.credentials[:i:i], a.credentials[i+1:]...)
				return 204, nil
			}
		}
		return 404, map[string]string{"errorCode": "inexistent_device_credential"}
	}
	return 0, nil
}

func TestAuth0WithoutSessionsAPIBlocksDeletesDeviceCredentialsAndReportsBan(t *testing.T) {
	id := "auth0|64abc0de1234567890abcdef"
	escaped := "auth0%7C64abc0de1234567890abcdef"
	tenant := newAuth0WithoutSessionsAPI(escaped, false)
	mock := startMock(t, tenant.route)
	auth0 := &Auth0Actions{BaseURL: mock.URL + "/api/v2", Token: "mgmt_token"}
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), id)
	ctx := context.Background()

	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(nil))
	if err != nil {
		t.Fatal(err)
	}
	if result.Action != LateResultBan || result.UserID != id {
		t.Fatalf("result = %+v", result)
	}
	if blocked, left := tenant.state(); !blocked || left != 0 {
		t.Fatalf("blocked %v, %d credentials left", blocked, left)
	}
	assertCalls(t, mock, []string{
		"DELETE /api/v2/users/" + escaped + "/sessions",
		"PATCH /api/v2/users/" + escaped,
		"DELETE /api/v2/users/" + escaped + "/refresh-tokens",
		"GET /api/v2/device-credentials",
		"GET /api/v2/device-credentials",
		"DELETE /api/v2/device-credentials/dcr_a",
		"DELETE /api/v2/device-credentials/dcr_b",
		"DELETE /api/v2/device-credentials/dcr_r",
	})
	calls := mock.snapshot()
	if calls[3].Query["type"] != "refresh_token" || calls[4].Query["type"] != "rotating_refresh_token" || calls[3].Query["user_id"] != id {
		t.Fatalf("device credential queries %v, %v", calls[3].Query, calls[4].Query)
	}
	assertEveryCall(t, mock, "Bearer mgmt_token")

	// The same delivery again: still blocked, nothing left to delete, same report.
	mock.reset()
	if result, err = HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(nil)); err != nil || result.Action != LateResultBan {
		t.Fatalf("repeat = %+v, %v", result, err)
	}
	for _, c := range mock.snapshot() {
		if c.Method == "DELETE" && strings.HasPrefix(c.Path, "/api/v2/device-credentials/") {
			t.Fatalf("repeat deleted %s", c.Path)
		}
	}

	// Ban blocks first, and the plan refusing the sessions endpoint is no failure.
	tenant.set(false, auth0Credential{"dcr_c", "refresh_token"})
	mock.reset()
	ban := lateOptions(&LateResultOptions{Policy: LateResultPolicy{components.AccountDecisionReview: LateResultBan}})
	if result, err = HandleLateResult(ctx, raw, headers, lateSecrets, auth0, ban); err != nil || result.Action != LateResultBan {
		t.Fatalf("ban = %+v, %v", result, err)
	}
	if first := mock.snapshot()[0]; first.Method != "PATCH" {
		t.Fatalf("ban began with %s %s", first.Method, first.Path)
	}
	if blocked, left := tenant.state(); !blocked || left != 0 {
		t.Fatalf("after ban: blocked %v, %d credentials left", blocked, left)
	}

	// Called directly, RevokeSessions takes the same fallback.
	tenant.set(false)
	if err := auth0.RevokeSessions(ctx, id); err != nil {
		t.Fatal(err)
	}
	if blocked, _ := tenant.state(); !blocked {
		t.Fatal("RevokeSessions did not block")
	}
}

var auth0TooManyRequests = map[string]any{"statusCode": 429, "error": "Too Many Requests", "message": "Global limit has been reached", "errorCode": "too_many_requests"}

// rateLimited answers 429 once to each call named in limited, with that
// entry's headers, and passes everything else to next. A GET is named by its
// type query.
func rateLimited(t *testing.T, limited map[string]map[string]string, next func(c mockCall) (int, any)) (*httptest.Server, *[]string, *sync.Mutex) {
	t.Helper()
	var mu sync.Mutex
	var sent []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		text, _ := io.ReadAll(r.Body)
		call := mockCall{Method: r.Method, Path: r.URL.EscapedPath(), Query: map[string]string{}}
		for k, v := range r.URL.Query() {
			call.Query[k] = v[0]
		}
		if len(text) > 0 {
			_ = json.Unmarshal(text, &call.Body)
		}
		key := call.Method + " " + call.Path
		if call.Method == "GET" {
			key = "GET " + call.Query["type"]
		}
		mu.Lock()
		sent = append(sent, key)
		headers, limit := limited[key]
		delete(limited, key)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		status, body := 429, any(auth0TooManyRequests)
		if limit {
			for k, v := range headers {
				w.Header().Set(k, v)
			}
		} else if status, body = next(call); status == 0 {
			status, body = 404, map[string]string{"error": "not found"}
		}
		w.WriteHeader(status)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
	t.Cleanup(server.Close)
	return server, &sent, &mu
}

func TestAuth0WaitsOutA429AndRetriesTheCall(t *testing.T) {
	tenant := newAuth0WithoutSessionsAPI("auth0%7Cabc", false)
	// One 429 each: Retry-After, an X-RateLimit-Reset already past, and no header (the backoff).
	limited := map[string]map[string]string{
		"GET refresh_token":                       {"Retry-After": "0"},
		"DELETE /api/v2/device-credentials/dcr_a": {"X-RateLimit-Reset": strconv.FormatInt(time.Now().Unix()-5, 10)},
		"DELETE /api/v2/device-credentials/dcr_b": {},
	}
	server, sent, mu := rateLimited(t, limited, tenant.route)
	auth0 := &Auth0Actions{BaseURL: server.URL + "/api/v2", Token: "mgmt_token"}
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), "auth0|abc")
	result, err := HandleLateResult(context.Background(), raw, headers, lateSecrets, auth0, lateOptions(nil))
	if err != nil {
		t.Fatal(err)
	}
	if result.Action != LateResultBan {
		t.Fatalf("result = %+v", result)
	}
	if blocked, left := tenant.state(); !blocked || left != 0 {
		t.Fatalf("blocked %v, %d credentials left", blocked, left)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(limited) != 0 {
		t.Fatalf("calls never sent: %v", limited)
	}
	count := map[string]int{}
	for _, k := range *sent {
		count[k]++
	}
	for _, k := range []string{"GET refresh_token", "DELETE /api/v2/device-credentials/dcr_a", "DELETE /api/v2/device-credentials/dcr_b"} {
		if count[k] != 2 {
			t.Fatalf("%s sent %d times; calls %v", k, count[k], *sent)
		}
	}
}

func TestAuth0ReturnsA429AtOnceWhenTheWaitPassesTheBudget(t *testing.T) {
	limited := map[string]map[string]string{"DELETE /api/v2/users/auth0%7Cabc/sessions": {"Retry-After": "60"}}
	server, sent, mu := rateLimited(t, limited, func(mockCall) (int, any) { return 200, nil })
	auth0 := &Auth0Actions{BaseURL: server.URL + "/api/v2", Token: "mgmt_token"}
	started := time.Now()
	_, err := auth0.Act(context.Background(), LateResultRevokeSessions, "auth0|abc")
	var refused *ProviderActionError
	if !errors.As(err, &refused) || refused.Status != 429 || refused.Code != "too_many_requests" {
		t.Fatalf("err = %v", err)
	}
	if elapsed := time.Since(started); elapsed > time.Second {
		t.Fatalf("waited %v", elapsed)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(*sent) != 1 {
		t.Fatalf("calls %v", *sent)
	}
}

func TestAuth0WithRefreshTokensButNoSessionsAPIStillBlocksAndAScopeRefusalNeverFallsBack(t *testing.T) {
	tenant := newAuth0WithoutSessionsAPI("auth0%7Cabc", true)
	mock := startMock(t, tenant.route)
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), "auth0|abc")
	ctx := context.Background()
	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, &Auth0Actions{BaseURL: mock.URL + "/api/v2", Token: "t"}, lateOptions(nil))
	if err != nil || result.Action != LateResultBan {
		t.Fatalf("result = %+v, %v", result, err)
	}
	if blocked, _ := tenant.state(); !blocked {
		t.Fatal("not blocked")
	}
	for _, c := range mock.snapshot() {
		if c.Path == "/api/v2/device-credentials" {
			t.Fatal("listed device credentials although the refresh-token endpoint worked")
		}
	}

	scoped := startMock(t, func(c mockCall) (int, any) {
		if c.Method == "DELETE" {
			return 403, map[string]any{"statusCode": 403, "error": "Forbidden", "message": "Insufficient scope, expected any of: delete:sessions", "errorCode": "insufficient_scope"}
		}
		return 200, map[string]any{}
	})
	_, err = HandleLateResult(ctx, raw, headers, lateSecrets, &Auth0Actions{BaseURL: scoped.URL + "/api/v2", Token: "t"}, lateOptions(nil))
	var refusal *ProviderActionError
	if !errors.As(err, &refusal) || refusal.Status != 403 || refusal.Code != "insufficient_scope" {
		t.Fatalf("err = %v", err)
	}
	assertCalls(t, scoped, []string{"DELETE /api/v2/users/auth0%7Cabc/sessions"})
}

type plainProvider struct{ done []string }

func (p *plainProvider) Name() string { return "plain" }
func (p *plainProvider) RevokeSessions(_ context.Context, userID string) error {
	p.done = append(p.done, "revoke "+userID)
	return nil
}
func (p *plainProvider) Ban(_ context.Context, userID string) error {
	p.done = append(p.done, "ban "+userID)
	return nil
}

func TestAProviderWithoutActReportsTheActionAskedFor(t *testing.T) {
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), "u_1")
	ctx := context.Background()
	p := &plainProvider{}
	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, p, lateOptions(nil))
	if err != nil || result.Action != LateResultRevokeSessions {
		t.Fatalf("revoke = %+v, %v", result, err)
	}
	result, err = HandleLateResult(ctx, raw, headers, lateSecrets, p, lateOptions(&LateResultOptions{Policy: LateResultPolicy{components.AccountDecisionReview: LateResultBan}}))
	if err != nil || result.Action != LateResultBan {
		t.Fatalf("ban = %+v, %v", result, err)
	}
	if !reflect.DeepEqual(p.done, []string{"revoke u_1", "ban u_1"}) {
		t.Fatalf("done = %v", p.done)
	}
}

func TestAuth0DomainBuildsTheManagementBase(t *testing.T) {
	c, err := (&Auth0Actions{Domain: "example.eu.auth0.com", Token: "t"}).caller()
	if err != nil || c.base != "https://example.eu.auth0.com/api/v2" {
		t.Fatalf("base = %v, %v", c, err)
	}
	if _, err := (&Auth0Actions{Token: "t"}).caller(); err == nil {
		t.Fatal("want an error without domain or base URL")
	}
}

func TestUnverifiedQuietOrUnmappedEventMakesNoProviderCallAndRefusalReturnsError(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) {
		if c.Method == "DELETE" {
			return 403, map[string]string{"error": "insufficient_scope", "secret": "mgmt_token"}
		}
		return 0, nil
	})
	auth0 := &Auth0Actions{BaseURL: mock.URL + "/api/v2", Token: "mgmt_token"}
	review := lateView(t, "assessment-phone-uncovered-complete.json")
	raw, headers := lateDelivery(t, review, "auth0|abc")
	ctx := context.Background()

	_, err := HandleLateResult(ctx, raw, headers, map[string]string{lateKey: "wrong"}, auth0, lateOptions(nil))
	wantCode(t, err, "invalid_signature")
	_, err = HandleLateResult(ctx, append(append([]byte{}, raw...), ' '), headers, lateSecrets, auth0, lateOptions(nil))
	wantCode(t, err, "invalid_signature")
	_, err = HandleLateResult(ctx, raw, headers, lateSecrets, auth0, &LateResultOptions{Webhook: []WebhookOption{at(lateNow + 301)}})
	wantCode(t, err, "invalid_timestamp")

	expired := lateFixture(t, "webhook-assessment-expired.json")["data"].(map[string]any)
	for name, data := range map[string]map[string]any{
		"step_up": lateView(t, "assessment-complete.json"),
		"pending": lateView(t, "assessment-pending.json"),
		"expired": expired,
	} {
		raw, headers := lateDelivery(t, data, "")
		result, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(nil))
		if err != nil || result.Action != LateResultNone || result.UserID != "" {
			t.Fatalf("%s: %+v, %v", name, result, err)
		}
	}
	// account.risk_changed is verified and typed, and acts on nothing: the
	// assessment.updated that moved the decision does.
	changed := lateFixture(t, "webhook-account-risk-changed.json")
	changedRaw, changedHeaders := lateDeliveryWith(t, changed["data"].(map[string]any), "", map[string]any{"type": changed["type"]})
	ignored, err := HandleLateResult(ctx, changedRaw, changedHeaders, lateSecrets, auth0, lateOptions(nil))
	if err != nil || ignored.Action != LateResultNone || ignored.Event.Type != EventAccountRiskChanged {
		t.Fatalf("account.risk_changed: %+v, %v", ignored, err)
	}
	if data, ok := ignored.Event.Data.(*AccountRiskChanged); !ok || data.Account.DecisionsByFlow.Login != components.AccountDecisionReview {
		t.Fatalf("account.risk_changed data %+v", ignored.Event.Data)
	}
	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(&LateResultOptions{UserID: func(string, *Event) string { return "" }}))
	if err != nil || result.Action != LateResultNone {
		t.Fatalf("empty resolver: %+v, %v", result, err)
	}
	if n := len(mock.summary()); n != 0 {
		t.Fatalf("%d provider calls before any action", n)
	}

	_, err = HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(nil))
	var refusal *ProviderActionError
	if !errors.As(err, &refusal) || refusal.Status != 403 || refusal.Provider != "auth0" {
		t.Fatalf("err = %v", err)
	}
	if strings.Contains(err.Error(), "mgmt_token") || strings.Contains(err.Error(), "insufficient_scope") {
		t.Fatalf("error leaks response or credential: %v", err)
	}
	if n := len(mock.summary()); n != 1 {
		t.Fatalf("%d provider calls, want 1", n)
	}
}

func TestEmptyUserIDAndMissingCredentialsAreErrors(t *testing.T) {
	ctx := context.Background()
	for _, p := range []LateResultProvider{
		&ClerkActions{SecretKey: "k", BaseURL: "http://127.0.0.1:1"},
		&WorkOSActions{APIKey: "k", BaseURL: "http://127.0.0.1:1"},
		&Auth0Actions{Token: "k", BaseURL: "http://127.0.0.1:1"},
	} {
		if err := p.RevokeSessions(ctx, ""); !errors.Is(err, ErrEmptyUserID) {
			t.Fatalf("%s revoke: %v", p.Name(), err)
		}
		if err := p.Ban(ctx, ""); !errors.Is(err, ErrEmptyUserID) {
			t.Fatalf("%s ban: %v", p.Name(), err)
		}
	}
	if err := (&ClerkActions{}).Ban(ctx, "u"); err == nil {
		t.Fatal("clerk without a key must fail")
	}
	if err := (&WorkOSActions{}).Ban(ctx, "u"); err == nil {
		t.Fatal("workos without a key must fail")
	}
	if err := (&Auth0Actions{}).Ban(ctx, "u"); err == nil {
		t.Fatal("auth0 without a token must fail")
	}
}

func TestWebhookTestEventActsOnNothing(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) { return 200, map[string]any{} })
	workos := &WorkOSActions{APIKey: "sk_test_workos", BaseURL: mock.URL}
	raw, headers := lateDeliveryWith(t, lateView(t, "assessment-complete.json"), "acct_7f3a9c", map[string]any{"test": true})
	resolved := 0
	ban := LateResultPolicy{components.AccountDecisionAllow: LateResultBan, components.AccountDecisionStepUp: LateResultBan, components.AccountDecisionReview: LateResultBan}
	for name, extra := range map[string]*LateResultOptions{
		"UserID":        {Policy: ban, UserID: func(string, *Event) string { resolved++; return "user_01WORKOS" }},
		"ResolveUserID": {Policy: ban, ResolveUserID: func(context.Context, string, *Event) (string, error) { resolved++; return "user_01WORKOS", nil }},
		"default":       {Policy: ban},
	} {
		result, err := HandleLateResult(context.Background(), raw, headers, lateSecrets, workos, lateOptions(extra))
		if err != nil || !result.Event.Test || result.Action != LateResultNone || result.UserID != "" {
			t.Fatalf("%s: %+v, %v", name, result, err)
		}
		if action, err := LateResultActionFor(result.Event, ban); err != nil || action != LateResultNone {
			t.Fatalf("%s: LateResultActionFor = %q, %v", name, action, err)
		}
	}
	if resolved != 0 {
		t.Fatalf("resolver called %d times for a test event", resolved)
	}
	if n := len(mock.summary()); n != 0 {
		t.Fatalf("%d provider calls for a test event", n)
	}
}

func TestRepeatedDeliveryTreatsAProvider404AsDone(t *testing.T) {
	// Every user and session is already gone, as after a first WorkOS ban deleted the user.
	mock := startMock(t, func(c mockCall) (int, any) { return 404, map[string]string{"error": "not_found"} })
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), "user_gone")
	for _, provider := range []LateResultProvider{
		&WorkOSActions{APIKey: "k", BaseURL: mock.URL},
		&ClerkActions{SecretKey: "k", BaseURL: mock.URL + "/v1"},
		&Auth0Actions{Token: "k", BaseURL: mock.URL + "/api/v2"},
	} {
		for _, action := range []LateResultAction{LateResultBan, LateResultRevokeSessions} {
			for attempt := 0; attempt < 2; attempt++ {
				policy := LateResultPolicy{components.AccountDecisionReview: action}
				result, err := HandleLateResult(context.Background(), raw, headers, lateSecrets, provider, lateOptions(&LateResultOptions{Policy: policy}))
				if err != nil || result.Action != action || result.UserID != "user_gone" {
					t.Fatalf("%s %s attempt %d: %+v, %v", provider.Name(), action, attempt, result, err)
				}
			}
		}
	}
}

func TestClerkSessionGoneBetweenListAndRevokeIsDoneAndOtherRefusalsStillFail(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) {
		switch {
		case c.Method == "GET" && c.Path == "/v1/sessions":
			return 200, []any{map[string]any{"id": "sess_1", "status": "active"}}
		case c.Method == "POST" && c.Path == "/v1/sessions/sess_1/revoke":
			return 404, map[string]any{"errors": []any{map[string]string{"code": "resource_not_found"}}}
		case c.Method == "POST" && c.Path == "/v1/users/user_2EXAMPLE/ban":
			return 500, map[string]string{"error": "boom"}
		}
		return 0, nil
	})
	clerk := &ClerkActions{SecretKey: "k", BaseURL: mock.URL + "/v1"}
	ctx := context.Background()
	if err := clerk.RevokeSessions(ctx, "user_2EXAMPLE"); err != nil {
		t.Fatalf("revoke = %v", err)
	}
	var refusal *ProviderActionError
	if err := clerk.Ban(ctx, "user_2EXAMPLE"); !errors.As(err, &refusal) || refusal.Status != 500 {
		t.Fatalf("ban = %v", err)
	}
}

func TestResolveUserIDGetsTheContextAndItsErrorFailsTheDelivery(t *testing.T) {
	mock := startMock(t, func(c mockCall) (int, any) { return 202, nil })
	auth0 := &Auth0Actions{Token: "k", BaseURL: mock.URL + "/api/v2"}
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), "acct_7f3a9c")
	type key struct{}
	ctx := context.WithValue(context.Background(), key{}, "request")
	storeDown := errors.New("store unavailable")

	_, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(&LateResultOptions{
		ResolveUserID: func(ctx context.Context, accountID string, event *Event) (string, error) {
			if ctx.Value(key{}) != "request" || accountID != "acct_7f3a9c" || event == nil {
				t.Fatalf("resolver got ctx %v, account %q, event %v", ctx.Value(key{}), accountID, event)
			}
			return "", storeDown
		},
		UserID: func(string, *Event) string { t.Fatal("UserID must not run when ResolveUserID is set"); return "" },
	}))
	if !errors.Is(err, storeDown) {
		t.Fatalf("err = %v, want the resolver's error", err)
	}
	if n := len(mock.summary()); n != 0 {
		t.Fatalf("%d provider calls after a failed resolve", n)
	}

	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(&LateResultOptions{
		ResolveUserID: func(context.Context, string, *Event) (string, error) { return "auth0|abc", nil },
	}))
	if err != nil || result.Action != LateResultRevokeSessions || result.UserID != "auth0|abc" {
		t.Fatalf("%+v, %v", result, err)
	}
	result, err = HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(&LateResultOptions{
		ResolveUserID: func(context.Context, string, *Event) (string, error) { return "", nil },
	}))
	if err != nil || result.Action != LateResultNone {
		t.Fatalf("empty resolve: %+v, %v", result, err)
	}
}

type refuseTransport struct{ t *testing.T }

func (r refuseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r.t.Errorf("the log-only provider called out: %s %s", req.Method, req.URL)
	return nil, errors.New("refused")
}

func TestLogOnlyActionsLogsWhatItWouldDoAndCallsNothing(t *testing.T) {
	saved := http.DefaultTransport
	http.DefaultTransport = refuseTransport{t}
	t.Cleanup(func() { http.DefaultTransport = saved })

	var buf bytes.Buffer
	logOnly := &LogOnlyActions{Logger: slog.New(slog.NewJSONHandler(&buf, nil))}
	if logOnly.Name() != "log_only" {
		t.Fatalf("Name() = %q", logOnly.Name())
	}
	raw, headers := lateDelivery(t, lateView(t, "assessment-phone-uncovered-complete.json"), "usr_8f2k1")
	ctx := context.Background()
	revoked, err := HandleLateResult(ctx, raw, headers, lateSecrets, logOnly, lateOptions(nil))
	if err != nil || revoked.Action != LateResultRevokeSessions || revoked.UserID != "usr_8f2k1" {
		t.Fatalf("revoke: %+v, %v", revoked, err)
	}
	ban := LateResultPolicy{components.AccountDecisionReview: LateResultBan}
	banned, err := HandleLateResult(ctx, raw, headers, lateSecrets, logOnly, lateOptions(&LateResultOptions{Policy: ban}))
	if err != nil || banned.Action != LateResultBan {
		t.Fatalf("ban: %+v, %v", banned, err)
	}
	testRaw, testHeaders := lateDeliveryWith(t, lateView(t, "assessment-phone-uncovered-complete.json"), "usr_8f2k1", map[string]any{"test": true})
	if result, err := HandleLateResult(ctx, testRaw, testHeaders, lateSecrets, logOnly, lateOptions(nil)); err != nil || result.Action != LateResultNone {
		t.Fatalf("test event: %+v, %v", result, err)
	}

	type record struct {
		Level  string `json:"level"`
		Msg    string `json:"msg"`
		Action string `json:"action"`
		UserID string `json:"user_id"`
	}
	var got []record
	for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
		var r record
		if err := json.Unmarshal([]byte(line), &r); err != nil {
			t.Fatalf("%q: %v", line, err)
		}
		got = append(got, r)
	}
	want := []record{
		{"INFO", "NoResponse late result (log only): would revoke every session of user usr_8f2k1", "revoke_sessions", "usr_8f2k1"},
		{"INFO", "NoResponse late result (log only): would ban user usr_8f2k1 and end their sessions", "ban", "usr_8f2k1"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("log = %+v, want %+v", got, want)
	}
	if err := logOnly.Ban(ctx, ""); !errors.Is(err, ErrEmptyUserID) {
		t.Fatalf("empty user id: %v", err)
	}
	var zero LogOnlyActions
	if err := zero.RevokeSessions(ctx, "usr_8f2k1"); err != nil {
		t.Fatalf("zero value: %v", err)
	}
}
