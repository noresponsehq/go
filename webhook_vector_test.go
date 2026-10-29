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
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/noresponsehq/go/v9/models/components"
)

type signatureVector struct {
	Secret    string            `json:"secret_utf8"`
	KeyID     string            `json:"key_id"`
	Timestamp string            `json:"timestamp"`
	Payload   string            `json:"payload"`
	Headers   map[string]string `json:"headers"`
}

func at(seconds int64) WebhookOption {
	return WithWebhookClock(func() time.Time { return time.Unix(seconds, 0) })
}

func wantCode(t *testing.T, err error, code WebhookVerificationErrorCode) {
	t.Helper()
	var verificationErr *WebhookVerificationError
	if !errors.As(err, &verificationErr) {
		t.Fatalf("want WebhookVerificationError %s, got %v", code, err)
	}
	if verificationErr.Code != code {
		t.Fatalf("want code %s, got %s (%s)", code, verificationErr.Code, verificationErr.Message)
	}
}

func loadVector(t *testing.T) (signatureVector, http.Header, map[string]string, int64) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", "webhook-signature-vector.json"))
	if err != nil {
		t.Fatal(err)
	}
	var v signatureVector
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatal(err)
	}
	headers := http.Header{}
	for name, value := range v.Headers {
		headers.Set(name, value)
	}
	var seconds int64
	if err := json.Unmarshal([]byte(v.Timestamp), &seconds); err != nil {
		t.Fatal(err)
	}
	return v, headers, map[string]string{v.KeyID: v.Secret}, seconds
}

func TestTheVectorVerifiesIntoATypedAssessmentEvent(t *testing.T) {
	v, headers, secrets, now := loadVector(t)
	event, err := ConstructEvent([]byte(v.Payload), headers, secrets, at(now))
	if err != nil {
		t.Fatal(err)
	}
	if event.Type != components.WebhookEventTypeAssessmentUpdated || event.AssessmentUpdatedEvent.Data.Assessment == nil {
		t.Fatalf("event = %+v", event)
	}
	if got := event.AssessmentUpdatedEvent.Data.Assessment.Detail.AssessmentVersion; got != 2 {
		t.Fatalf("assessment version %d", got)
	}
}

func TestATamperedStaleUnknownOrRepeatedDeliveryIsRefused(t *testing.T) {
	v, headers, secrets, now := loadVector(t)
	repeated := headers.Clone()
	repeated["nr-webhook-signature"] = []string{v.Headers["NR-Webhook-Signature"]}
	missing := headers.Clone()
	missing.Del("NR-Webhook-Id")
	for name, c := range map[string]struct {
		payload string
		headers http.Header
		secrets map[string]string
		now     int64
		code    WebhookVerificationErrorCode
	}{
		"tampered": {v.Payload + " ", headers, secrets, now, WebhookInvalidSignature},
		"stale":    {v.Payload, headers, secrets, now + 301, WebhookInvalidTimestamp},
		"unknown":  {v.Payload, headers, map[string]string{"whk_other0001": v.Secret}, now, WebhookUnknownKey},
		"repeated": {v.Payload, repeated, secrets, now, WebhookDuplicateHeader},
		"missing":  {v.Payload, missing, secrets, now, WebhookMissingHeader},
	} {
		t.Run(name, func(t *testing.T) {
			_, err := ConstructEvent([]byte(c.payload), c.headers, c.secrets, at(c.now))
			wantCode(t, err, c.code)
		})
	}
}

// typedMember returns the one union member the event's type names, and
// whether every other member is unset.
func typedMember(event *components.WebhookEvent) (any, bool) {
	members := map[components.WebhookEventType]any{
		components.WebhookEventTypeAssessmentUpdated:       event.AssessmentUpdatedEvent,
		components.WebhookEventTypeEndpointVerification:    event.EndpointVerificationEvent,
		components.WebhookEventTypeAccountDeleted:          event.AccountDeletedEvent,
		components.WebhookEventTypeWorkspaceDeleted:        event.WorkspaceDeletedEvent,
		components.WebhookEventTypeWebhookEndpointDisabled: event.WebhookEndpointDisabledEvent,
		components.WebhookEventTypeAccountRiskChanged:      event.AccountRiskChangedEvent,
		components.WebhookEventTypeAuditEvent:              event.AuditEventEvent,
	}
	for eventType, member := range members {
		if eventType != event.Type && !reflect.ValueOf(member).IsNil() {
			return nil, false
		}
	}
	member, ok := members[event.Type]
	return member, ok && !reflect.ValueOf(member).IsNil()
}

func TestEveryContractEventFixtureParsesIntoItsTypedEvent(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("tests", "fixtures", "valid", "webhook-*.json"))
	if err != nil {
		t.Fatal(err)
	}
	const now = 1791968402
	parsed := 0
	for _, path := range paths {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Fatal(err)
		}
		eventType, ok := body["type"].(string)
		if !ok {
			continue
		}
		compact, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		mac := hmac.New(sha256.New, []byte("fixture-secret"))
		mac.Write([]byte("1791968402."))
		mac.Write(compact)
		headers := http.Header{}
		headers.Set("NR-Webhook-Id", body["id"].(string))
		headers.Set("NR-Webhook-Timestamp", "1791968402")
		headers.Set("NR-Webhook-Key-Id", "whk_fixture0001")
		headers.Set("NR-Webhook-Signature", "v1="+hex.EncodeToString(mac.Sum(nil)))
		event, err := ConstructEvent(compact, headers, map[string]string{"whk_fixture0001": "fixture-secret"}, at(now))
		if err != nil {
			t.Fatalf("%s: %v", filepath.Base(path), err)
		}
		if string(event.Type) != eventType {
			t.Fatalf("%s: type %s", filepath.Base(path), event.Type)
		}
		member, ok := typedMember(event)
		if !ok {
			t.Fatalf("%s: no single typed member for %s: %+v", filepath.Base(path), event.Type, event)
		}
		typed, err := json.Marshal(member)
		if err != nil {
			t.Fatal(err)
		}
		var roundTrip map[string]any
		if err := json.Unmarshal(typed, &roundTrip); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(roundTrip, body) {
			t.Fatalf("%s: typed event %s differs from the fixture %s", filepath.Base(path), typed, compact)
		}
		parsed++
	}
	if parsed == 0 || !strings.Contains(strings.Join(paths, ","), "webhook-account-risk-changed") {
		t.Fatalf("parsed %d fixtures from %v", parsed, paths)
	}
}
