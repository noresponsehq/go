package noresponse

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/noresponsehq/go/v10/models/apierrors"
	"github.com/noresponsehq/go/v10/models/components"
)

// Consumer smoke test of the client against a live API.
//
// Skipped unless NORESPONSE_LIVE_URL and NORESPONSE_LIVE_API_KEY name an API
// and a key with the assess and history scopes (a test key keeps it off real
// sources). It assesses a synthetic account, replays the request, reads it
// back, waits for it to settle, lists its history and checks the typed 404.
// NORESPONSE_LIVE_WEBHOOK_VECTORS names deliveries the same API signed, in
// the shape of webhook-signature-vector.json, for the verifier. Each step
// writes its name to NORESPONSE_LIVE_REPORT, and "done" at the end, so a
// failure names the operation (integrations/regression/suites/sdks.sh).
func TestLiveConsumer(t *testing.T) {
	url, key := os.Getenv("NORESPONSE_LIVE_URL"), os.Getenv("NORESPONSE_LIVE_API_KEY")
	if url == "" || key == "" {
		t.Skip("set NORESPONSE_LIVE_URL and NORESPONSE_LIVE_API_KEY to run against a live API")
	}
	step := func(name string) {
		if report := os.Getenv("NORESPONSE_LIVE_REPORT"); report != "" {
			if err := os.WriteFile(report, []byte(name), 0o600); err != nil {
				t.Fatal(err)
			}
		}
	}
	ctx := context.Background()
	client := New(WithServerURL(url), WithSecurity(key), WithTimeout(30*time.Second))
	accountID := "sdk-live-go-" + liveRun()
	idempotencyKey := "sdk-live-" + randomHex(16)
	email := "sdk-live@example.com"
	request := components.AssessmentRequest{
		AccountID:   accountID,
		Flow:        components.FlowAccountCreation,
		Identifiers: components.AssessmentRequestIdentifiers{Email: &email},
	}

	step("create")
	created, err := client.Assessments.Create(ctx, idempotencyKey, request)
	if err != nil {
		t.Fatal(err)
	}
	id := created.Result.ID
	if created.Result.Detail.AccountID != accountID {
		t.Fatalf("create answered account %q", created.Result.Detail.AccountID)
	}

	step("replay")
	replay, err := client.Assessments.Create(ctx, idempotencyKey, request)
	if err != nil || replay.Result.ID != id {
		t.Fatalf("replay answered %v, %v", replay, err)
	}

	step("get")
	fetched, err := client.Assessments.Get(ctx, id, nil)
	if err != nil || fetched.Result.ID != id {
		t.Fatalf("get answered %v, %v", fetched, err)
	}

	step("wait")
	waitCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	settled, err := WaitForAssessment(waitCtx, client, id, nil)
	if err != nil {
		t.Fatal(err)
	}
	if settled.Status != components.StatusComplete && settled.Status != components.StatusPartial || settled.Account == nil {
		t.Fatalf("wait answered status %s, account %v", settled.Status, settled.Account)
	}

	step("query")
	limit := int64(5)
	history, err := client.Assessments.Query(ctx, components.AssessmentQuery{AccountID: accountID, Limit: &limit})
	if err != nil {
		t.Fatal(err)
	}
	if !slices.ContainsFunc(history.Result.Items, func(item components.AssessmentDetail) bool { return item.AssessmentID == id }) ||
		slices.ContainsFunc(history.Result.Items, func(item components.AssessmentDetail) bool { return item.AccountID != accountID }) {
		t.Fatalf("query answered %+v", history.Result.Items)
	}

	step("list_account_assessments")
	page, err := client.Accounts.ListAssessments(ctx, accountID, &limit, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.ContainsFunc(page.Result.Items, func(item components.AssessmentDetail) bool { return item.AssessmentID == id }) {
		t.Fatalf("list answered %+v", page.Result.Items)
	}

	step("get_account")
	account, err := client.Accounts.Get(ctx, accountID)
	if err != nil {
		t.Fatal(err)
	}
	if account.Result.ID != accountID || account.Result.Risk.Decision != settled.Account.Risk.Decision {
		t.Fatalf("account answered %+v", account.Result)
	}

	// The typed error carries the problem document's documented fields; its
	// const Status field reads 0 after decoding (NOR-568), so the code says 404.
	step("not_found")
	_, err = client.Assessments.Get(ctx, "asm_00000000000000000000000000", nil)
	var missing *apierrors.NotFoundError
	if !errors.As(err, &missing) || missing.Code != components.NotFoundErrorCodeNotFound || missing.RequestID == "" {
		t.Fatalf("an unknown assessment answered %v", err)
	}

	if path := os.Getenv("NORESPONSE_LIVE_WEBHOOK_VECTORS"); path != "" {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var vectors []struct {
			signatureVector
			Type string `json:"type"`
		}
		if err := json.Unmarshal(raw, &vectors); err != nil {
			t.Fatal(err)
		}
		for _, vector := range vectors {
			step("verify_webhook " + vector.Type)
			headers := http.Header{}
			for name, value := range vector.Headers {
				headers.Set(name, value)
			}
			seconds, err := strconv.ParseInt(vector.Timestamp, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			secrets := map[string]string{vector.KeyID: vector.Secret}
			event, err := ConstructEvent([]byte(vector.Payload), headers, secrets, at(seconds))
			if err != nil || string(event.Type) != vector.Type {
				t.Fatalf("the signed %s delivery answered %v, %v", vector.Type, event, err)
			}
			_, err = ConstructEvent([]byte(vector.Payload+" "), headers, secrets, at(seconds))
			wantCode(t, err, WebhookInvalidSignature)
		}
	}
	step("done")
}

// liveRun is NORESPONSE_LIVE_RUN, or eight random hex characters.
func liveRun() string {
	if run := os.Getenv("NORESPONSE_LIVE_RUN"); run != "" {
		return run
	}
	return randomHex(4)
}

func randomHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
