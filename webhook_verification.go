package noresponse

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/noresponsehq/go/v9/models/components"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// WebhookVerificationErrorCode says why a webhook delivery was refused.
type WebhookVerificationErrorCode string

const (
	WebhookMissingHeader    WebhookVerificationErrorCode = "missing_header"
	WebhookDuplicateHeader  WebhookVerificationErrorCode = "duplicate_header"
	WebhookInvalidTimestamp WebhookVerificationErrorCode = "invalid_timestamp"
	WebhookUnknownKey       WebhookVerificationErrorCode = "unknown_key"
	WebhookInvalidSignature WebhookVerificationErrorCode = "invalid_signature"
	WebhookInvalidPayload   WebhookVerificationErrorCode = "invalid_payload"
)

// WebhookVerificationError is returned by ConstructEvent for a delivery that must not be processed. Answer it with a 400.
type WebhookVerificationError struct {
	Code    WebhookVerificationErrorCode
	Message string
}

func (e *WebhookVerificationError) Error() string {
	return "webhook verification failed: " + e.Message + " (" + string(e.Code) + ")"
}

func webhookError(code WebhookVerificationErrorCode, message string) error {
	return &WebhookVerificationError{Code: code, Message: message}
}

// DefaultWebhookTolerance is how far the NR-Webhook-Timestamp header may be from the receiver's clock.
const DefaultWebhookTolerance = 300 * time.Second

type webhookOptions struct {
	tolerance time.Duration
	now       func() time.Time
}

// WebhookOption configures ConstructEvent.
type WebhookOption func(*webhookOptions)

// WithWebhookTolerance replaces DefaultWebhookTolerance.
func WithWebhookTolerance(tolerance time.Duration) WebhookOption {
	return func(o *webhookOptions) { o.tolerance = tolerance }
}

// WithWebhookClock sets the receiver's clock used for the tolerance check.
func WithWebhookClock(now func() time.Time) WebhookOption {
	return func(o *webhookOptions) { o.now = now }
}

var webhookHeaderNames = []string{"NR-Webhook-Id", "NR-Webhook-Timestamp", "NR-Webhook-Key-Id",
	"NR-Webhook-Signature",
}

// ConstructEvent verifies a webhook delivery and parses its body.
//
// payload is the raw request body exactly as received: re-encoding parsed JSON breaks the signature.
// secrets maps each key id, sent in NR-Webhook-Key-Id, to its secret, so a receiver can hold two during a rotation.
// A header repeated in any letter case, or sent as a comma-joined list, is refused.
func ConstructEvent(payload []byte, headers http.Header, secrets map[string]string, opts ...WebhookOption) (*components.WebhookEvent, error) {
	o := webhookOptions{tolerance: DefaultWebhookTolerance, now: time.Now}
	for _, opt := range opts {
		opt(&o)
	}

	values, err := webhookHeaderValues(headers)
	if err != nil {
		return nil, err
	}
	id := values["NR-Webhook-Id"]
	timestamp := values["NR-Webhook-Timestamp"]
	keyID := values["NR-Webhook-Key-Id"]
	signature := values["NR-Webhook-Signature"]

	seconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil || len(timestamp) > 12 || strings.TrimLeft(timestamp, "0123456789") != "" {
		return nil, webhookError(WebhookInvalidTimestamp, "NR-Webhook-Timestamp is not a Unix time in seconds")
	}
	if skew := o.now().Sub(time.Unix(seconds, 0)); skew > o.tolerance || skew < -o.tolerance {
		return nil, webhookError(WebhookInvalidTimestamp, "NR-Webhook-Timestamp is outside the tolerance")
	}
	secret, ok := secrets[keyID]
	if !ok {
		return nil, webhookError(WebhookUnknownKey, "NR-Webhook-Key-Id names a key this receiver does not hold")
	}
	if secret == "" {
		return nil, webhookError(WebhookInvalidSignature, "the webhook secret is empty")
	}

	encoded, ok := strings.CutPrefix(signature, "v1=")
	if !ok {
		return nil, webhookError(WebhookInvalidSignature, "NR-Webhook-Signature is malformed")
	}
	received, err := hex.DecodeString(encoded)
	if err != nil {
		return nil, webhookError(WebhookInvalidSignature, "NR-Webhook-Signature is malformed")
	}

	substitute := strings.NewReplacer("{timestamp}", timestamp, "{id}", id)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(substitute.Replace("{timestamp}.")))
	mac.Write(payload)
	if !hmac.Equal(mac.Sum(nil), received) {
		return nil, webhookError(WebhookInvalidSignature, "NR-Webhook-Signature does not match")
	}
	var envelope map[string]json.RawMessage
	var bodyID string
	if json.Unmarshal(payload, &envelope) != nil || json.Unmarshal(envelope["id"], &bodyID) != nil || bodyID != id {
		return nil, webhookError(WebhookInvalidPayload, "the body's id does not equal NR-Webhook-Id")
	}

	var event components.WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, webhookError(WebhookInvalidPayload, "the body is not a known webhook event")
	}
	return &event, nil
}

func webhookHeaderValues(headers http.Header) (map[string]string, error) {
	values := make(map[string]string, len(webhookHeaderNames))
	for _, name := range webhookHeaderNames {
		var found []string
		for key, vals := range headers {
			if strings.EqualFold(key, name) {
				found = append(found, vals...)
			}
		}
		if len(found) > 1 || (len(found) == 1 && strings.Contains(found[0], ",")) {
			return nil, webhookError(WebhookDuplicateHeader, name+" is repeated")
		}
		if len(found) == 0 || found[0] == "" {
			return nil, webhookError(WebhookMissingHeader, name+" is missing")
		}
		values[name] = found[0]
	}
	return values, nil
}
