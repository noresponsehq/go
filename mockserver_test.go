package noresponse

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/noresponsehq/go/v8/models/components"
)

// Smoke test of the client against the mock server.
//
// Skipped unless NORESPONSE_MOCK_URL names a running mock server
// (`generate_sdks.py mockserver --serve`). The mock server answers with the
// examples in the SDK document, so this proves the client's wiring (URL, auth
// header, request encoding, response parsing) and nothing about the live API.
func mockClient(t *testing.T) *NoResponse {
	t.Helper()
	url := os.Getenv("NORESPONSE_MOCK_URL")
	if url == "" {
		t.Skip("set NORESPONSE_MOCK_URL to a running mock server")
	}
	return New(WithServerURL(url), WithSecurity("nrk_mock_kkkkkkkkkkkkkkkkkkkk"))
}

func TestMockServerCreate(t *testing.T) {
	email := "new.user@example.com"
	res, err := mockClient(t).Assessments.Create(context.Background(), "mock-server-smoke-0001", components.AssessmentRequest{
		AccountID:   "acct_7f3a9c",
		Flow:        components.FlowAccountCreation,
		Identifiers: components.AssessmentRequestIdentifiers{Email: &email},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || !strings.HasPrefix(res.Result.ID, "asm_") {
		t.Fatalf("unexpected response: %+v", res)
	}
}

func TestMockServerGet(t *testing.T) {
	res, err := mockClient(t).Assessments.Get(context.Background(), "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || !strings.HasPrefix(res.Result.ID, "asm_") {
		t.Fatalf("unexpected response: %+v", res)
	}
}
