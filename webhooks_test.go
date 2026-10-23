package noresponse

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/noresponsehq/go/v4/models/components"
)

type signatureVector struct {
	SecretUTF8 string            `json:"secret_utf8"`
	KeyID      string            `json:"key_id"`
	Timestamp  string            `json:"timestamp"`
	Payload    string            `json:"payload"`
	Headers    map[string]string `json:"headers"`
}

func fixture(t *testing.T, name string, into any) []byte {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", name))
	if err != nil {
		t.Fatal(err)
	}
	if into != nil {
		if err := json.Unmarshal(raw, into); err != nil {
			t.Fatal(err)
		}
	}
	return raw
}

func vectorHeaders(v signatureVector) http.Header {
	h := http.Header{}
	for name, value := range v.Headers {
		h.Set(name, value)
	}
	return h
}

func at(seconds int64) WebhookOption {
	return WithClock(func() time.Time { return time.Unix(seconds, 0) })
}

func sign(t *testing.T, body []byte, secret, timestamp, id, keyID string) http.Header {
	t.Helper()
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "."))
	mac.Write(body)
	h := http.Header{}
	h.Set("NR-Webhook-Id", id)
	h.Set("NR-Webhook-Timestamp", timestamp)
	h.Set("NR-Webhook-Key-Id", keyID)
	h.Set("NR-Webhook-Signature", "v1="+hex.EncodeToString(mac.Sum(nil)))
	return h
}

func wantCode(t *testing.T, err error, code string) {
	t.Helper()
	var webhookErr *WebhookError
	if !errors.As(err, &webhookErr) || !errors.Is(err, ErrWebhook) {
		t.Fatalf("want WebhookError %s, got %v", code, err)
	}
	if webhookErr.Code != code {
		t.Fatalf("want code %s, got %s (%s)", code, webhookErr.Code, webhookErr.Message)
	}
}

func TestSignatureVectorYieldsTypedAssessment(t *testing.T) {
	var v signatureVector
	fixture(t, "webhook-signature-vector.json", &v)
	event, err := ConstructEvent([]byte(v.Payload), vectorHeaders(v), map[string]string{v.KeyID: v.SecretUTF8}, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	assessment, ok := event.Data.(*components.Assessment)
	if !ok {
		t.Fatalf("data is %T", event.Data)
	}
	if event.Type != EventAssessmentUpdated || event.ID != v.Headers["NR-Webhook-Id"] || event.CreatedAt != "2026-10-14T09:00:02.000Z" {
		t.Fatalf("envelope %+v", event)
	}
	if assessment.ID != "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM" || assessment.Detail.AssessmentVersion != 2 {
		t.Fatalf("assessment %+v", assessment)
	}
	if assessment.Account == nil || assessment.Account.Risk.Decision != components.AccountDecisionStepUp {
		t.Fatalf("account %+v", assessment.Account)
	}
	if event.Expired() {
		t.Fatal("a complete assessment is not the expiry event")
	}
}

func TestRefusals(t *testing.T) {
	var v signatureVector
	fixture(t, "webhook-signature-vector.json", &v)
	payload := []byte(v.Payload)
	secrets := map[string]string{v.KeyID: v.SecretUTF8}
	const now = 1791968402

	_, err := ConstructEvent(append(append([]byte{}, payload...), ' '), vectorHeaders(v), secrets, at(now))
	wantCode(t, err, "invalid_signature")
	_, err = ConstructEvent(payload, vectorHeaders(v), secrets, at(now+301))
	wantCode(t, err, "invalid_timestamp")
	_, err = ConstructEvent(payload, vectorHeaders(v), secrets, at(now-301))
	wantCode(t, err, "invalid_timestamp")
	if _, err = ConstructEvent(payload, vectorHeaders(v), secrets, at(now+301), WithTolerance(600*time.Second)); err != nil {
		t.Fatalf("a wider tolerance accepts the delivery: %v", err)
	}

	repeated := vectorHeaders(v)
	repeated.Add("NR-Webhook-Id", "duplicate")
	_, err = ConstructEvent(payload, repeated, secrets, at(now))
	wantCode(t, err, "duplicate_header")
	spelled := vectorHeaders(v)
	spelled["nr-webhook-ID"] = []string{v.Headers["NR-Webhook-Id"]}
	_, err = ConstructEvent(payload, spelled, secrets, at(now))
	wantCode(t, err, "duplicate_header")

	missing := vectorHeaders(v)
	missing.Del("NR-Webhook-Signature")
	_, err = ConstructEvent(payload, missing, secrets, at(now))
	wantCode(t, err, "missing_header")

	_, err = ConstructEvent(payload, vectorHeaders(v), map[string]string{"whk_other0001": v.SecretUTF8}, at(now))
	wantCode(t, err, "unknown_key")
	_, err = ConstructEvent(payload, vectorHeaders(v), map[string]string{v.KeyID: "other-secret"}, at(now))
	wantCode(t, err, "invalid_signature")
	malformed := vectorHeaders(v)
	malformed.Set("NR-Webhook-Signature", "v2=abc")
	_, err = ConstructEvent(payload, malformed, secrets, at(now))
	wantCode(t, err, "invalid_signature")
}

func TestRotationAndEventID(t *testing.T) {
	complete := fixture(t, "valid/assessment-view-complete.json", nil)
	const stamp = "1791968402"
	body, _ := json.Marshal(map[string]any{"id": "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2", "type": "assessment.updated", "created_at": "2026-10-18T11:00:04Z", "data": json.RawMessage(complete)})
	both := map[string]string{"whk_sdktest01": "old-secret", "whk_sdktest02": "new-secret"}

	event, err := ConstructEvent(body, sign(t, body, "new-secret", stamp, "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2", "whk_sdktest02"), both, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := event.Data.(*components.Assessment); !ok {
		t.Fatalf("data is %T", event.Data)
	}
	_, err = ConstructEvent(body, sign(t, body, "new-secret", stamp, "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ9", "whk_sdktest02"), both, at(1791968402))
	wantCode(t, err, "invalid_payload")

	// A pointer without the assessment body is not an assessment.
	pointer, _ := json.Marshal(map[string]any{"id": "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2", "type": "assessment.updated", "created_at": "2026-10-18T11:00:04Z",
		"data": map[string]any{"assessment_id": "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "assessment_url": "/v1/assessments/x"}})
	_, err = ConstructEvent(pointer, sign(t, pointer, "new-secret", stamp, "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2", "whk_sdktest02"), both, at(1791968402))
	wantCode(t, err, "invalid_payload")

	// A bare assessment detail, without the account and detail envelope, is
	// not an assessment either.
	flat, _ := json.Marshal(map[string]any{"id": "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2", "type": "assessment.updated", "created_at": "2026-10-18T11:00:04Z",
		"data": json.RawMessage(fixture(t, "valid/assessment-complete.json", nil))})
	_, err = ConstructEvent(flat, sign(t, flat, "new-secret", stamp, "evt_01JA2ZB4N5P6Q7R8S9T0VWXYZ2", "whk_sdktest02"), both, at(1791968402))
	wantCode(t, err, "invalid_payload")
}

func TestExpiredAndVerificationEvents(t *testing.T) {
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}

	expired := fixture(t, "valid/webhook-assessment-expired.json", nil)
	event, err := ConstructEvent(expired, sign(t, expired, "secret", stamp, "evt_01JA2ZC4N5P6Q7R8S9T0VWXYZ3", "whk_sdktest01"), secrets, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	notice, ok := event.Data.(*AssessmentExpired)
	if !ok || !event.Expired() {
		t.Fatalf("data is %T", event.Data)
	}
	if *notice != (AssessmentExpired{AssessmentID: "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM", AccountID: "acct_7f3a9c", AssessmentVersion: 2, Status: "expired", ChangeReason: "expired"}) {
		t.Fatalf("notice %+v", notice)
	}
	if _, err := VerificationResponse(event); err == nil {
		t.Fatal("an expiry event has no challenge")
	}

	ping := fixture(t, "valid/webhook-endpoint-verification.json", nil)
	event, err = ConstructEvent(ping, sign(t, ping, "secret", stamp, "evt_01JA2ZD4N5P6Q7R8S9T0VWXYZ4", "whk_sdktest01"), secrets, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	answer, err := VerificationResponse(event)
	if err != nil || answer["challenge"] != "q3Jt8mV0x2Yb7Lk9Pz4Rw6Hn1Sd5Fg8Ja0Tc3Ue7Wi4" {
		t.Fatalf("answer %v, %v", answer, err)
	}
	if event.Data.(*EndpointVerification).EndpointID != "whe_01JA2Z0A1B2C3D4E5F6G7H8J9K" {
		t.Fatalf("endpoint %+v", event.Data)
	}
}

func TestAccountDeletedEvent(t *testing.T) {
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}
	erased := fixture(t, "valid/webhook-account-deleted.json", nil)
	event, err := ConstructEvent(erased, sign(t, erased, "secret", stamp, "evt_01JA2ZE5N6P7Q8R9S0T1VWXYZ4", "whk_sdktest01"), secrets, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	deleted, ok := event.Data.(*AccountDeleted)
	if !ok || event.Type != EventAccountDeleted || event.Expired() || event.Test {
		t.Fatalf("event %+v", event)
	}
	if *deleted != (AccountDeleted{AccountID: "acct_7f3a9c", AssessmentsDeleted: 3, DeletedAt: "2026-10-23T09:00:00.000Z"}) {
		t.Fatalf("data %+v", deleted)
	}
	if _, err := VerificationResponse(event); err == nil {
		t.Fatal("an account.deleted event has no challenge")
	}
	for _, mutate := range []func(map[string]any){
		func(data map[string]any) { data["assessments_deleted"] = 0 },
		func(data map[string]any) { data["assessments_deleted"] = 1.5 },
		func(data map[string]any) { data["account_id"] = "jane@example.com" },
		func(data map[string]any) { data["deleted_at"] = "tomorrow" },
	} {
		var body map[string]any
		if err := json.Unmarshal(erased, &body); err != nil {
			t.Fatal(err)
		}
		mutate(body["data"].(map[string]any))
		raw, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		_, err = ConstructEvent(raw, sign(t, raw, "secret", stamp, body["id"].(string), "whk_sdktest01"), secrets, at(1791968402))
		wantCode(t, err, "invalid_payload")
	}
}

const (
	tenantDeletedBody    = `{"id":"evt_01JA2ZD3K4M5N6P7Q8R9S0T1V2","type":"tenant.deleted","created_at":"2026-10-23T10:00:00.000Z","data":{"workspace_id":"tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM","environment":"live","deleted_at":"2026-10-23T10:00:00.000Z"}}`
	endpointDisabledBody = `{"id":"evt_01JA2ZD3K4M5N6P7Q8R9S0T1V3","type":"webhook_endpoint.disabled","created_at":"2026-10-23T10:00:00.000Z","data":{"endpoint_id":"whe_01JA2ZD0E1F2G3H4J5K6M7N8P9","failing_since":"2026-10-20T10:00:00.000Z","disabled_at":"2026-10-23T10:00:00.000Z"}}`
)

// refuseMutatedData signs one copy of body per mutation of its data and checks
// each is refused as invalid_payload.
func refuseMutatedData(t *testing.T, body string, mutations []func(map[string]any)) {
	t.Helper()
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}
	for _, mutate := range mutations {
		var event map[string]any
		if err := json.Unmarshal([]byte(body), &event); err != nil {
			t.Fatal(err)
		}
		mutate(event["data"].(map[string]any))
		raw, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}
		_, err = ConstructEvent(raw, sign(t, raw, "secret", stamp, event["id"].(string), "whk_sdktest01"), secrets, at(1791968402))
		wantCode(t, err, "invalid_payload")
	}
}

func TestTenantDeletedEvent(t *testing.T) {
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}
	raw := []byte(tenantDeletedBody)
	event, err := ConstructEvent(raw, sign(t, raw, "secret", stamp, "evt_01JA2ZD3K4M5N6P7Q8R9S0T1V2", "whk_sdktest01"), secrets, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	deleted, ok := event.Data.(*TenantDeleted)
	if !ok || event.Type != EventTenantDeleted || event.Expired() || event.Test || event.CreatedAt != "2026-10-23T10:00:00.000Z" {
		t.Fatalf("event %+v", event)
	}
	if *deleted != (TenantDeleted{WorkspaceID: "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", Environment: "live", DeletedAt: "2026-10-23T10:00:00.000Z"}) {
		t.Fatalf("data %+v", deleted)
	}
	if _, err := VerificationResponse(event); err == nil {
		t.Fatal("a tenant.deleted event has no challenge")
	}
	refuseMutatedData(t, tenantDeletedBody, []func(map[string]any){
		func(data map[string]any) { delete(data, "workspace_id") },
		func(data map[string]any) { data["workspace_id"] = "-tnt" },
		func(data map[string]any) { data["workspace_id"] = "tnt/01" },
		func(data map[string]any) { data["environment"] = "staging" },
		func(data map[string]any) { delete(data, "environment") },
		func(data map[string]any) { data["deleted_at"] = "tomorrow" },
		func(data map[string]any) { delete(data, "deleted_at") },
	})
}

func TestEndpointDisabledEvent(t *testing.T) {
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}
	raw := []byte(endpointDisabledBody)
	event, err := ConstructEvent(raw, sign(t, raw, "secret", stamp, "evt_01JA2ZD3K4M5N6P7Q8R9S0T1V3", "whk_sdktest01"), secrets, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	disabled, ok := event.Data.(*EndpointDisabled)
	if !ok || event.Type != EventEndpointDisabled || event.Expired() || event.Test {
		t.Fatalf("event %+v", event)
	}
	if *disabled != (EndpointDisabled{EndpointID: "whe_01JA2ZD0E1F2G3H4J5K6M7N8P9", FailingSince: "2026-10-20T10:00:00.000Z", DisabledAt: "2026-10-23T10:00:00.000Z"}) {
		t.Fatalf("data %+v", disabled)
	}
	if _, err := VerificationResponse(event); err == nil {
		t.Fatal("a webhook_endpoint.disabled event has no challenge")
	}
	refuseMutatedData(t, endpointDisabledBody, []func(map[string]any){
		func(data map[string]any) { data["endpoint_id"] = "whk_01JA2ZD0E1F2G3H4J5K6M7N8P9" },
		func(data map[string]any) { delete(data, "endpoint_id") },
		func(data map[string]any) { data["failing_since"] = "three days ago" },
		func(data map[string]any) { delete(data, "failing_since") },
		func(data map[string]any) { data["disabled_at"] = "now" },
		func(data map[string]any) { delete(data, "disabled_at") },
	})
}

func TestAccountRiskChangedEvent(t *testing.T) {
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}
	changed := fixture(t, "valid/webhook-account-risk-changed.json", nil)
	event, err := ConstructEvent(changed, sign(t, changed, "secret", stamp, "evt_01JA2ZE5N6P7Q8R9S0T1VWXYZ6", "whk_sdktest01"), secrets, at(1791968402))
	if err != nil {
		t.Fatal(err)
	}
	data, ok := event.Data.(*AccountRiskChanged)
	if !ok || event.Type != EventAccountRiskChanged || event.Expired() || event.Test {
		t.Fatalf("event %+v", event)
	}
	// data.Account is the same type Accounts.Get returns.
	var account components.Account = data.Account
	if account.ID != "acct_7f3a9c" || account.Risk.Score != 70 || account.Risk.Decision != components.AccountDecisionReview ||
		account.AssessmentID != "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM" || account.UpdatedAt != "2026-10-23T09:00:04.812Z" {
		t.Fatalf("account %+v", account)
	}
	if len(account.Risk.Reasons) != 2 || account.Risk.Reasons[0].Code != "recent_password_exposure" || account.Risk.Reasons[0].Weight != 55 {
		t.Fatalf("reasons %+v", account.Risk.Reasons)
	}
	review := components.AccountDecisionsByFlow{Signup: components.AccountDecisionReview, Login: components.AccountDecisionReview, Purchase: components.AccountDecisionReview}
	allow := components.AccountDecisionsByFlow{Signup: components.AccountDecisionAllow, Login: components.AccountDecisionAllow, Purchase: components.AccountDecisionAllow}
	if account.DecisionsByFlow != review || data.PreviousDecisionsByFlow != allow {
		t.Fatalf("decisions %+v, previous %+v", account.DecisionsByFlow, data.PreviousDecisionsByFlow)
	}
	if _, err := VerificationResponse(event); err == nil {
		t.Fatal("an account.risk_changed event has no challenge")
	}
	for name, mutate := range map[string]func(map[string]any){
		"no previous":            func(data map[string]any) { delete(data, "previous_decisions_by_flow") },
		"previous missing login": func(data map[string]any) { delete(data["previous_decisions_by_flow"].(map[string]any), "login") },
		"no decisions_by_flow":   func(data map[string]any) { delete(data["account"].(map[string]any), "decisions_by_flow") },
		"no risk score":          func(data map[string]any) { delete(data["account"].(map[string]any)["risk"].(map[string]any), "score") },
		"account is a string":    func(data map[string]any) { data["account"] = "acct_7f3a9c" },
		"account id is an email": func(data map[string]any) { data["account"].(map[string]any)["id"] = "jane@example.com" },
	} {
		var body map[string]any
		if err := json.Unmarshal(changed, &body); err != nil {
			t.Fatal(err)
		}
		mutate(body["data"].(map[string]any))
		raw, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		_, err = ConstructEvent(raw, sign(t, raw, "secret", stamp, body["id"].(string), "whk_sdktest01"), secrets, at(1791968402))
		if err == nil {
			t.Fatalf("%s: accepted", name)
		}
		wantCode(t, err, "invalid_payload")
	}
}

// The API adds fields without a new version, so a signed event that carries
// keys this SDK does not know still verifies.
func TestUnknownFieldsAreAccepted(t *testing.T) {
	secrets := map[string]string{"whk_sdktest01": "secret"}
	for _, name := range []string{
		"invalid/webhook-tenant-id.json",
		"invalid/webhook-identifier-leak.json",
		"invalid/webhook-expired-with-assessment-url.json",
		"invalid/webhook-account-deleted-assessment-id.json",
	} {
		raw := fixture(t, name, nil)
		var body map[string]any
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Fatal(err)
		}
		if _, err := ConstructEvent(raw, sign(t, raw, "secret", "1791968402", body["id"].(string), "whk_sdktest01"), secrets, at(1791968402)); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
	}
}

func TestContractInvalidFixturesAndMalformedIdentifiersAreRefused(t *testing.T) {
	const stamp = "1791968402"
	secrets := map[string]string{"whk_sdktest01": "secret"}
	for _, name := range []string{
		"invalid/webhook-update-names-source.json",
		"invalid/webhook-unknown-event-type.json",
		"invalid/webhook-update-pointer-only.json",
		"invalid/webhook-update-with-expired-status.json",
		"invalid/webhook-verification-without-challenge.json",
		"invalid/webhook-account-deleted-none.json",
	} {
		t.Run(name, func(t *testing.T) {
			raw := fixture(t, name, nil)
			var body map[string]any
			if err := json.Unmarshal(raw, &body); err != nil {
				t.Fatal(err)
			}
			_, err := ConstructEvent(raw, sign(t, raw, "secret", stamp, body["id"].(string), "whk_sdktest01"), secrets, at(1791968402))
			if err == nil {
				t.Fatalf("want invalid_payload, got <nil>")
			}
			wantCode(t, err, "invalid_payload")
		})
	}

	valid := fixture(t, "valid/webhook-endpoint-verification.json", nil)
	for _, mutate := range []func(map[string]any){
		func(body map[string]any) { body["id"] = "evt_bad" },
		func(body map[string]any) { body["data"].(map[string]any)["endpoint_id"] = "whe_bad" },
		func(body map[string]any) { body["created_at"] = "tomorrow" },
		func(body map[string]any) { body["created_at"] = "2026-10-14" },
	} {
		var body map[string]any
		if err := json.Unmarshal(valid, &body); err != nil {
			t.Fatal(err)
		}
		mutate(body)
		raw, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		_, err = ConstructEvent(raw, sign(t, raw, "secret", stamp, body["id"].(string), "whk_sdktest01"), secrets, at(1791968402))
		wantCode(t, err, "invalid_payload")
	}
}

func TestATestEventSaysSoAndARealOneDoesNot(t *testing.T) {
	secrets := map[string]string{"whk_sdktest01": "secret"}
	for name, want := range map[string]bool{"valid/webhook-assessment-updated-test.json": true, "valid/webhook-assessment-updated.json": false} {
		body := fixture(t, name, nil)
		var envelope struct {
			ID string `json:"id"`
		}
		if err := json.Unmarshal(body, &envelope); err != nil {
			t.Fatal(err)
		}
		event, err := ConstructEvent(body, sign(t, body, "secret", "1791968402", envelope.ID, "whk_sdktest01"), secrets, at(1791968402))
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if _, ok := event.Data.(*components.Assessment); !ok || event.Test != want {
			t.Fatalf("%s: test = %v, data %T", name, event.Test, event.Data)
		}
	}
}
