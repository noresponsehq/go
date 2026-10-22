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

	"github.com/noresponsehq/go/v2/models/components"
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
	if assessment.AssessmentID != "asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM" || assessment.AssessmentVersion != 2 {
		t.Fatalf("assessment %+v", assessment)
	}
	if assessment.Risk == nil || assessment.Risk.Recommendation != components.RecommendationStepUp {
		t.Fatalf("risk %+v", assessment.Risk)
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
	complete := fixture(t, "valid/assessment-complete.json", nil)
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

// The API adds fields without a new version, so a signed event that carries
// keys this SDK does not know still verifies.
func TestUnknownFieldsAreAccepted(t *testing.T) {
	secrets := map[string]string{"whk_sdktest01": "secret"}
	for _, name := range []string{
		"invalid/webhook-tenant-id.json",
		"invalid/webhook-identifier-leak.json",
		"invalid/webhook-expired-with-assessment-url.json",
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
