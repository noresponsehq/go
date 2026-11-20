package noresponse

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/noresponsehq/go/v20/models/components"
	"github.com/noresponsehq/go/v20/models/operations"
	"github.com/noresponsehq/go/v20/optionalnullable"
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
	start := "2026-12-03T09:30:00Z"
	body := components.GroupRequest{Name: "Retry test", StartsAt: optionalnullable.From(&start)}
	ctx := context.Background()
	if _, err := client.Groups.Create(ctx, body); err == nil {
		t.Fatal("an unavailable API returned success")
	}
	key := NewIdempotencyKey()
	if _, err := client.Groups.Create(ctx, body, operations.WithGroupsCreateIdempotencyKey(key)); err == nil {
		t.Fatal("an unavailable API returned success")
	}
	if len(keys) != 2 || keys[0] != "" || keys[1] != key {
		t.Fatalf("request keys or automatic retries changed: %v", keys)
	}
}
