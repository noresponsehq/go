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

	"github.com/noresponsehq/go/v8/models/components"
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
