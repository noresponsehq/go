package noresponse

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/noresponsehq/go/v13/models/components"
	"github.com/noresponsehq/go/v13/models/operations"
)

func TestGroupCreationKeepsItsArgumentsAndAcceptsARequestKeyOption(t *testing.T) {
	var keys []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys = append(keys, r.Header.Get("Idempotency-Key"))
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(`{"type":"https://docs.noresponse.org/problems/service-unavailable","title":"Service unavailable","status":503,"detail":"Try again later.","code":"service_unavailable","request_id":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}`))
	}))
	defer server.Close()
	client := NewClient("nrk_test_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	body := components.GroupRequest{Name: "Retry test", StartsAt: "2026-12-03T09:30:00Z"}
	ctx := context.Background()
	workspace := "tnt_01JC8Y0A2B4C6D8E0F2G4H6J8K"
	if _, err := client.Groups.Create(ctx, workspace, components.KeyEnvironmentTest, body); err == nil {
		t.Fatal("an unavailable API returned success")
	}
	key := NewIdempotencyKey()
	if _, err := client.Groups.Create(ctx, workspace, components.KeyEnvironmentTest, body, operations.WithGroupsCreateIdempotencyKey(key)); err == nil {
		t.Fatal("an unavailable API returned success")
	}
	if len(keys) != 2 || keys[0] != "" || keys[1] != key {
		t.Fatalf("request keys or automatic retries changed: %v", keys)
	}
}
