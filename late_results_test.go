package noresponse

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/noresponsehq/go/v2/models/components"
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

// lateDelivery signs an assessment.updated delivery of data, overriding account_id when given.
func lateDelivery(t *testing.T, data map[string]any, accountID string) ([]byte, http.Header) {
	t.Helper()
	copied := map[string]any{}
	for k, v := range data {
		copied[k] = v
	}
	if accountID != "" {
		copied["account_id"] = accountID
	}
	id := "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2"
	raw, err := json.Marshal(map[string]any{"id": id, "type": "assessment.updated", "created_at": "2026-10-18T11:00:04Z", "data": copied})
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

func lateEvent(status string, recommendation string) *Event {
	data := &components.Assessment{Status: components.Status(status)}
	if recommendation != "" {
		data.Recommendation = components.Recommendation(recommendation).ToPointer()
	}
	return &Event{ID: "evt_x", Type: EventAssessmentUpdated, CreatedAt: "2026-10-18T11:00:04Z", Data: data}
}

func TestDefaultPolicyRevokesOnReviewOnly(t *testing.T) {
	expect := func(event *Event, policy LateResultPolicy, want LateResultAction) {
		t.Helper()
		got, err := LateResultActionFor(event, policy)
		if err != nil || got != want {
			t.Fatalf("got %q, %v; want %q", got, err, want)
		}
	}
	expect(lateEvent("complete", "review"), nil, LateResultRevokeSessions)
	expect(lateEvent("complete", "step_up"), nil, LateResultNone)
	expect(lateEvent("partial", "allow"), nil, LateResultNone)
	expect(lateEvent("pending", ""), nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventAssessmentUpdated, Data: &AssessmentExpired{Status: "expired", ChangeReason: "expired"}}, nil, LateResultNone)
	expect(&Event{ID: "evt_x", Type: EventEndpointVerification, Data: &EndpointVerification{}}, nil, LateResultNone)
	expect(lateEvent("complete", "step_up"), LateResultPolicy{components.RecommendationStepUp: LateResultBan}, LateResultBan)
	expect(lateEvent("complete", "review"), LateResultPolicy{components.RecommendationStepUp: LateResultBan}, LateResultNone)

	// The recommendation falls back to risk.recommendation.
	fallback := &Event{Type: EventAssessmentUpdated, Data: &components.Assessment{Status: "complete", Risk: &components.Risk{Recommendation: components.RecommendationReview}}}
	expect(fallback, nil, LateResultRevokeSessions)

	if _, err := LateResultActionFor(lateEvent("complete", "review"), LateResultPolicy{components.RecommendationReview: "delete"}); !errors.Is(err, ErrUnknownLateResultAction) {
		t.Fatalf("unknown action error = %v", err)
	}
	want := LateResultPolicy{"allow": "none", "step_up": "none", "review": "revoke_sessions"}
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
	review := lateFixture(t, "assessment-phone-uncovered-complete.json")
	raw, headers := lateDelivery(t, review, "user_2EXAMPLE")
	ctx := context.Background()

	result, err := HandleLateResult(ctx, raw, headers, lateSecrets, clerk, lateOptions(nil))
	if err != nil {
		t.Fatal(err)
	}
	if result.Action != LateResultRevokeSessions || result.UserID != "user_2EXAMPLE" {
		t.Fatalf("result = %+v", result)
	}
	if assessment, ok := result.Event.Data.(*components.Assessment); !ok || assessment.AssessmentID != review["assessment_id"] {
		t.Fatalf("event data = %+v", result.Event.Data)
	}
	assertCalls(t, mock, []string{"GET /v1/sessions", "POST /v1/sessions/sess_1/revoke", "POST /v1/sessions/sess_2/revoke"})
	if got, want := mock.snapshot()[0].Query, map[string]string{"user_id": "user_2EXAMPLE", "status": "active", "limit": "100", "offset": "0"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("query = %v", got)
	}
	assertEveryCall(t, mock, "Bearer sk_test_mock")

	mock.reset()
	banned, err := HandleLateResult(ctx, raw, headers, lateSecrets, clerk, lateOptions(&LateResultOptions{Policy: LateResultPolicy{components.RecommendationReview: LateResultBan}}))
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
	step := lateFixture(t, "assessment-complete.json")
	raw, headers := lateDelivery(t, step, "acct_7f3a9c")
	result, err := HandleLateResult(context.Background(), raw, headers, lateSecrets, workos, lateOptions(&LateResultOptions{
		Policy: LateResultPolicy{components.RecommendationStepUp: LateResultBan},
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
	review := lateFixture(t, "assessment-phone-uncovered-complete.json")
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
	if _, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(&LateResultOptions{Policy: LateResultPolicy{components.RecommendationReview: LateResultBan}})); err != nil {
		t.Fatal(err)
	}
	assertCalls(t, mock, []string{"PATCH /api/v2/users/" + escaped, "DELETE /api/v2/users/" + escaped + "/sessions", "DELETE /api/v2/users/" + escaped + "/refresh-tokens"})
	if body := mock.snapshot()[0].Body; !reflect.DeepEqual(body, map[string]any{"blocked": true}) {
		t.Fatalf("patch body = %v", body)
	}
	assertEveryCall(t, mock, "Bearer mgmt_token")
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
	review := lateFixture(t, "assessment-phone-uncovered-complete.json")
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
		"step_up": lateFixture(t, "assessment-complete.json"),
		"pending": lateFixture(t, "assessment-pending.json"),
		"expired": expired,
	} {
		raw, headers := lateDelivery(t, data, "")
		result, err := HandleLateResult(ctx, raw, headers, lateSecrets, auth0, lateOptions(nil))
		if err != nil || result.Action != LateResultNone || result.UserID != "" {
			t.Fatalf("%s: %+v, %v", name, result, err)
		}
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
