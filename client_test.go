package noresponse

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/noresponsehq/go/v9/models/components"
)

func TestNewClientSendsTheKeyAsBearer(t *testing.T) {
	var auth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "assessment-view-complete.json"))
		if err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	client := NewClient("nrk_test_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	if _, err := client.Assessments.Get(context.Background(), "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil); err != nil {
		t.Fatal(err)
	}
	if auth != "Bearer nrk_test_kkkkkkkkkkkkkkkkkkkk" {
		t.Fatalf("Authorization was %q", auth)
	}
}

func TestNewIdempotencyKeyFitsTheContract(t *testing.T) {
	allowed := regexp.MustCompile(`^[A-Za-z0-9._:-]{16,128}$`)
	seen := map[string]bool{}
	for range 50 {
		key := NewIdempotencyKey()
		if !allowed.MatchString(key) {
			t.Fatalf("key %q is outside the contract", key)
		}
		if seen[key] {
			t.Fatalf("key %q repeated", key)
		}
		seen[key] = true
	}
}

func TestAllAssessmentsFollowsCursorsAndStopsEarly(t *testing.T) {
	fixture, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "assessment-complete.json"))
	if err != nil {
		t.Fatal(err)
	}
	var assessment json.RawMessage = fixture
	var cursors []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Cursor string `json:"cursor"`
		}
		_ = json.NewDecoder(r.Body).Decode(&request)
		cursors = append(cursors, request.Cursor)
		page := map[string]any{
			"schema_version": "1",
			"account_id":     "acct_7f3a9c",
			"items":          []json.RawMessage{assessment, assessment},
		}
		if request.Cursor == "" {
			page["next_cursor"] = "cursor-two"
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()

	client := NewClient("nrk_test_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	query := components.AssessmentQuery{AccountID: "acct_7f3a9c"}
	count := 0
	for _, err := range AllAssessments(context.Background(), client, query) {
		if err != nil {
			t.Fatal(err)
		}
		count++
	}
	if count != 4 || len(cursors) != 2 || cursors[1] != "cursor-two" {
		t.Fatalf("got %d items over cursors %v", count, cursors)
	}

	cursors = nil
	count = 0
	for _, err := range AllAssessments(context.Background(), client, query) {
		if err != nil {
			t.Fatal(err)
		}
		count++
		break
	}
	if count != 1 || len(cursors) != 1 {
		t.Fatalf("early stop fetched %d pages for %d items", len(cursors), count)
	}
}

// A management key creates keys and changes settings; a settings change
// sends only what it names.
func TestKeysAndSettingsSendOnlyWhatTheyName(t *testing.T) {
	type seen struct {
		method, path string
		body         map[string]any
	}
	var requests []seen
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		requests = append(requests, seen{r.Method, r.URL.Path, body})
		name := "key-created-by-management-key.json"
		if r.Method == http.MethodPatch {
			name = "settings.json"
		}
		fixture, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", name))
		if err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusCreated)
		}
		_, _ = w.Write(fixture)
	}))
	defer server.Close()

	ctx := context.Background()
	client := NewClient("nrk_live_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	created, err := client.Keys.Create(ctx, "tnt_a", components.KeyRequest{Environment: components.KeyEnvironmentLive,
		Scopes: []components.KeyRequestScope{components.KeyRequestScopeAssess, components.KeyRequestScopeManagement}})
	if err != nil || created.Result.Key == nil || *created.Result.CreatedBy != "nrk_live_01JA2ZE4K5M6N7P8Q9R0S1T2V3" {
		t.Fatalf("create: %+v %v", created, err)
	}
	stepUp := int64(25)
	if _, err := client.Settings.Update(ctx, "tnt_a", components.KeyEnvironmentLive, components.SettingsRequest{
		Flows: &components.SettingsRequestFlows{Login: &components.FlowChange{StepUp: &stepUp}}}); err != nil {
		t.Fatal(err)
	}
	if len(requests) != 2 || requests[0].path != "/v1/workspaces/tnt_a/keys" || requests[1].method != http.MethodPatch ||
		requests[1].path != "/v1/workspaces/tnt_a/settings/live" {
		t.Fatalf("requests %+v", requests)
	}
	login := requests[1].body["flows"].(map[string]any)["login"].(map[string]any)
	if len(login) != 1 || login["step_up"] != float64(25) {
		t.Fatalf("settings change sent %v", requests[1].body)
	}
}

func TestIdentifierSearchRequestsAndMatchEvidence(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", "investigation-responses.json"))
	if err != nil {
		t.Fatal(err)
	}
	var examples map[string]json.RawMessage
	if err := json.Unmarshal(raw, &examples); err != nil {
		t.Fatal(err)
	}
	var revision map[string]any
	if err := json.Unmarshal(examples["revision"], &revision); err != nil {
		t.Fatal(err)
	}
	match := map[string]any{"kind": "username", "value": "JaneDoe", "match_basis": "canonical", "namespace": "customer", "provenance": "submitted", "assessment_id": revision["id"]}
	matchedRevision := map[string]any{}
	for key, value := range revision {
		matchedRevision[key] = value
	}
	matchedRevision["matched_identifier"] = match
	pages := map[string]any{
		"/v1/accounts/search":    map[string]any{"schema_version": "1", "items": []any{map[string]any{"account_id": "acct_7f3a9c", "latest": revision, "assessments": 2, "matched_identifier": match}}, "total": 1, "total_capped": false},
		"/v1/assessments/search": map[string]any{"retention_days": 365, "schema_version": "1", "items": []any{matchedRevision}, "total": 1, "total_capped": false},
	}
	seen := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen++
		if r.Method != http.MethodPost || r.Header.Get("Authorization") != "Bearer nrk_test_identifier_fixture" {
			t.Error("incorrect method or credential")
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		identifier, ok := body["identifier"].(map[string]any)
		if !ok || identifier["kind"] != "username" || identifier["value"] != "janedoe" || identifier["match"] != "canonical" || identifier["namespace"] != "customer" {
			t.Errorf("identifier was not preserved: %#v", identifier)
		}
		page, ok := pages[r.URL.Path]
		if !ok {
			t.Errorf("unexpected path %s", r.URL.Path)
			w.WriteHeader(404)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()
	identifierJSON := `{"kind":"username","value":"janedoe","match":"canonical","namespace":"customer"}`
	var identifier components.IdentifierSearch
	if err := json.Unmarshal([]byte(identifierJSON), &identifier); err != nil {
		t.Fatal(err)
	}
	var accountRequest components.AccountDirectorySearchRequest
	if err := json.Unmarshal([]byte(`{"identifier":`+identifierJSON+`,"limit":1}`), &accountRequest); err != nil {
		t.Fatal(err)
	}
	var assessmentRequest components.AssessmentSearchRequest
	if err := json.Unmarshal([]byte(`{"from":"2026-10-01T00:00:00Z","to":"2026-10-02T00:00:00Z","identifier":`+identifierJSON+`}`), &assessmentRequest); err != nil {
		t.Fatal(err)
	}
	client := NewClient("nrk_test_identifier_fixture", WithServerURL(server.URL))
	accounts, err := client.Accounts.Search(context.Background(), accountRequest)
	if err != nil {
		t.Fatal(err)
	}
	assessments, err := client.Assessments.Search(context.Background(), assessmentRequest)
	if err != nil {
		t.Fatal(err)
	}
	for _, found := range []*components.IdentifierMatch{accounts.Result.Items[0].MatchedIdentifier, assessments.Result.Items[0].MatchedIdentifier} {
		if found == nil || found.Value != "JaneDoe" || string(found.MatchBasis) != "canonical" || string(found.Provenance) != "submitted" || found.AssessmentID != revision["id"] {
			t.Errorf("match evidence changed: %#v", found)
		}
	}
	if accounts.Result.Items[0].Current != nil {
		t.Error("missing current risk must remain absent")
	}
	if identifier.Value != "janedoe" || seen != 2 {
		t.Fatalf("identifier value %q, calls %d", identifier.Value, seen)
	}
}

func TestCombinedIdentifierSearchRequestsAndWitnesses(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", "investigation-responses.json"))
	if err != nil {
		t.Fatal(err)
	}
	var examples map[string]json.RawMessage
	if err := json.Unmarshal(raw, &examples); err != nil {
		t.Fatal(err)
	}
	var revision map[string]any
	if err := json.Unmarshal(examples["revision"], &revision); err != nil {
		t.Fatal(err)
	}
	expressionJSON := `{"scope":"assessment","expression":{"and":[{"term":{"kind":"phone","value":"+14155552671","match":"canonical"}},{"or":[{"term":{"kind":"email","value":"other@example.com","match":"exact"}},{"term":{"kind":"username","value":"janedoe","match":"canonical","namespace":"customer"}}]}]}}`
	var expression map[string]any
	if err := json.Unmarshal([]byte(expressionJSON), &expression); err != nil {
		t.Fatal(err)
	}
	assessmentExpected, _ := json.Marshal(expression)
	expression["scope"] = "account"
	accountExpected, _ := json.Marshal(expression)
	matches := []any{
		map[string]any{"kind": "phone", "value": "+14155552671", "match_basis": "canonical", "provenance": "submitted", "assessment_id": revision["id"], "term_index": 0},
		map[string]any{"kind": "username", "value": "JaneDoe", "match_basis": "canonical", "namespace": "customer", "provenance": "submitted", "assessment_id": revision["id"], "term_index": 2},
	}
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != http.MethodPost || r.Header.Get("Authorization") != "Bearer nrk_test_identifier_fixture" {
			t.Error("incorrect method or credential")
		}
		var request map[string]any
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Error(err)
		}
		actual, _ := json.Marshal(request["identifier_expression"])
		expected := assessmentExpected
		if r.URL.Path == "/v1/accounts/search" {
			expected = accountExpected
		}
		if string(actual) != string(expected) || request["identifier"] != nil {
			t.Errorf("expression did not survive: %s", actual)
		}
		item := map[string]any{}
		for key, value := range revision {
			item[key] = value
		}
		item["matched_identifiers"] = matches
		if r.URL.Path == "/v1/accounts/search" {
			item = map[string]any{"account_id": "acct_7f3a9c", "latest": revision, "assessments": 2, "matched_identifiers": matches}
		}
		page := map[string]any{"schema_version": "1", "total": 1, "total_capped": false, "items": []any{item}}
		if r.URL.Path == "/v1/assessments/search" {
			page["retention_days"] = 365
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()
	var accountRequest components.AccountDirectorySearchRequest
	if err := json.Unmarshal([]byte(`{"identifier_expression":`+string(accountExpected)+`}`), &accountRequest); err != nil {
		t.Fatal(err)
	}
	var assessmentRequest components.AssessmentSearchRequest
	if err := json.Unmarshal([]byte(`{"from":"2026-10-01T00:00:00Z","to":"2026-10-02T00:00:00Z","identifier_expression":`+expressionJSON+`}`), &assessmentRequest); err != nil {
		t.Fatal(err)
	}
	client := NewClient("nrk_test_identifier_fixture", WithServerURL(server.URL))
	accounts, err := client.Accounts.Search(context.Background(), accountRequest)
	if err != nil {
		t.Fatal(err)
	}
	assessments, err := client.Assessments.Search(context.Background(), assessmentRequest)
	if err != nil {
		t.Fatal(err)
	}
	for _, witnesses := range [][]components.IdentifierTermMatch{accounts.Result.Items[0].MatchedIdentifiers, assessments.Result.Items[0].MatchedIdentifiers} {
		if len(witnesses) != 2 || witnesses[0].TermIndex != 0 || witnesses[1].TermIndex != 2 || witnesses[1].Value != "JaneDoe" || witnesses[1].AssessmentID != revision["id"] {
			t.Fatalf("witnesses changed: %#v", witnesses)
		}
	}
	if accounts.Result.Items[0].MatchedIdentifier != nil || assessments.Result.Items[0].MatchedIdentifier != nil || calls != 2 {
		t.Fatal("combined search mixed singular evidence or omitted a call")
	}
}
