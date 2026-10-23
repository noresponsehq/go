package noresponse

// Webhook verification. Hand-written beside the generated client.

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/noresponsehq/go/v5/models/components"
)

// DefaultWebhookTolerance is how far a delivery's timestamp may be from now.
const DefaultWebhookTolerance = 300 * time.Second

// Event types a webhook endpoint receives.
const (
	EventAssessmentUpdated    = "assessment.updated"
	EventEndpointVerification = "endpoint.verification"
	EventAccountDeleted       = "account.deleted"
	EventWorkspaceDeleted     = "workspace.deleted"
	EventEndpointDisabled     = "webhook_endpoint.disabled"
	EventAccountRiskChanged   = "account.risk_changed"
)

var (
	webhookKeyID        = regexp.MustCompile(`^whk_[A-Za-z0-9]{8,64}$`)
	webhookSignature    = regexp.MustCompile(`^v1=[0-9a-f]{64}$`)
	webhookTimestamp    = regexp.MustCompile(`^[0-9]{1,12}$`)
	webhookChallenge    = regexp.MustCompile(`^[A-Za-z0-9_-]{43}$`)
	webhookEventID      = regexp.MustCompile(`^evt_[0-9A-HJKMNP-TV-Z]{26}$`)
	webhookAssessmentID = regexp.MustCompile(`^asm_[0-9A-HJKMNP-TV-Z]{26}$`)
	webhookAccountID    = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._:|-]{0,127}$`)
	webhookEndpointID   = regexp.MustCompile(`^whe_[0-9A-HJKMNP-TV-Z]{26}$`)
	webhookWorkspaceID  = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9_.-]{0,122}$`)
	webhookHeaders      = [...]string{"NR-Webhook-Id", "NR-Webhook-Timestamp", "NR-Webhook-Key-Id", "NR-Webhook-Signature"}
)

// ErrWebhook matches every *WebhookError with errors.Is.
var ErrWebhook = errors.New("noresponse: webhook verification failed")

// WebhookError is returned by ConstructEvent and VerificationResponse. Answer
// 400 and do not process the delivery. Code is one of missing_header,
// duplicate_header, invalid_timestamp, unknown_key, invalid_signature or
// invalid_payload.
type WebhookError struct {
	Code    string
	Message string
}

func (e *WebhookError) Error() string { return "noresponse: " + e.Message + " (" + e.Code + ")" }

// Is makes errors.Is(err, ErrWebhook) true for every WebhookError.
func (e *WebhookError) Is(target error) bool { return target == ErrWebhook }

func webhookError(code, message string) error { return &WebhookError{Code: code, Message: message} }

// AssessmentExpired is the final event for an assessment: it has expired and
// Get now answers 404. AssessmentVersion is the last served version.
type AssessmentExpired struct {
	AssessmentID      string `json:"assessment_id"`
	AccountID         string `json:"account_id"`
	AssessmentVersion int64  `json:"assessment_version"`
	Status            string `json:"status"`
	ChangeReason      string `json:"change_reason"`
}

// EndpointVerification is the challenge NoResponse sends when an endpoint is
// registered. Answer it with VerificationResponse.
type EndpointVerification struct {
	EndpointID string `json:"endpoint_id"`
	Challenge  string `json:"challenge"`
}

// AccountDeleted says an account you erased with Accounts.Delete is gone:
// delete what you hold for it. AssessmentsDeleted is at least 1, and
// DeletedAt equals deleted_at in the Delete answer.
type AccountDeleted struct {
	AccountID          string `json:"account_id"`
	AssessmentsDeleted int64  `json:"assessments_deleted"`
	DeletedAt          string `json:"deleted_at"`
}

// WorkspaceDeleted says the workspace was deleted. It is sent once per
// environment, before that environment's endpoints are torn down, and is the
// last event the endpoint receives. Environment is "live" or "test".
type WorkspaceDeleted struct {
	WorkspaceID string `json:"workspace_id"`
	Environment string `json:"environment"`
	DeletedAt   string `json:"deleted_at"`
}

// EndpointDisabled says another endpoint of this environment, EndpointID,
// was disabled after failing continuously since FailingSince. It receives
// nothing more until it is verified again.
type EndpointDisabled struct {
	EndpointID   string `json:"endpoint_id"`
	FailingSince string `json:"failing_since"`
	DisabledAt   string `json:"disabled_at"`
}

// AccountRiskChanged says the account's decision for at least one flow
// changed. Account is the account exactly as Accounts.Get returns it, with its
// new risk and decisions; PreviousDecisionsByFlow holds the decisions it
// replaced. It is not sent when the score moves within the same decision, nor
// for an account's first risk.
type AccountRiskChanged struct {
	Account                 components.Account                `json:"account"`
	PreviousDecisionsByFlow components.AccountDecisionsByFlow `json:"previous_decisions_by_flow"`
}

// Event is a verified delivery. Data is *components.Assessment for an
// assessment.updated event (exactly as Assessments.Get returns it),
// *AssessmentExpired for the terminal expiry notice, *EndpointVerification
// for endpoint.verification, *AccountDeleted for account.deleted,
// *WorkspaceDeleted for workspace.deleted, *EndpointDisabled for
// webhook_endpoint.disabled and *AccountRiskChanged for
// account.risk_changed. Deliveries are at least
// once and unordered: deduplicate on ID and keep the highest
// Detail.AssessmentVersion per assessment.
//
//	switch data := event.Data.(type) {
//	case *components.Assessment:
//	case *noresponse.AssessmentExpired:
//	case *noresponse.EndpointVerification:
//	case *noresponse.AccountDeleted:
//	case *noresponse.WorkspaceDeleted:
//	case *noresponse.EndpointDisabled:
//	case *noresponse.AccountRiskChanged:
//	}
type Event struct {
	ID        string
	Type      string
	CreatedAt string
	// Test is true on an event sent by WebhookEndpoints.SendTestEvent: Data
	// is a fixed example assessment, not one of yours. Acknowledge it and
	// apply nothing.
	Test bool
	Data any
}

// Expired reports whether this is the terminal expiry event. Delete what you
// hold for the assessment and fetch nothing.
func (e *Event) Expired() bool {
	_, ok := e.Data.(*AssessmentExpired)
	return ok
}

// WebhookOption adjusts ConstructEvent.
type WebhookOption func(*webhookOptions)

type webhookOptions struct {
	tolerance time.Duration
	now       func() time.Time
}

// WithTolerance sets how far NR-Webhook-Timestamp may be from now. The
// default is DefaultWebhookTolerance.
func WithTolerance(tolerance time.Duration) WebhookOption {
	return func(o *webhookOptions) { o.tolerance = tolerance }
}

// WithClock sets the receiver's clock, for tests.
func WithClock(now func() time.Time) WebhookOption {
	return func(o *webhookOptions) { o.now = now }
}

// ConstructEvent verifies one delivery and returns the typed Event.
//
// payload is the request body exactly as received: the signature is
// HMAC-SHA256 over "<timestamp>.<body>", so re-encoding parsed JSON breaks
// it. Each NR-Webhook-* header must appear exactly once, in any letter case.
// secrets maps each whk_ key id NoResponse gave you to its secret; hold both
// during a rotation. Every failure is a *WebhookError matching ErrWebhook.
func ConstructEvent(payload []byte, headers http.Header, secrets map[string]string, opts ...WebhookOption) (*Event, error) {
	o := webhookOptions{tolerance: DefaultWebhookTolerance, now: time.Now}
	for _, opt := range opts {
		opt(&o)
	}
	values, err := webhookHeaderValues(headers)
	if err != nil {
		return nil, err
	}
	id, timestamp, keyID, signature := values[0], values[1], values[2], values[3]
	if !webhookTimestamp.MatchString(timestamp) {
		return nil, webhookError("invalid_timestamp", "webhook timestamp is outside the replay window")
	}
	seconds, _ := strconv.ParseInt(timestamp, 10, 64)
	skew := o.now().Unix() - seconds
	if skew < 0 {
		skew = -skew
	}
	if skew > int64(o.tolerance/time.Second) {
		return nil, webhookError("invalid_timestamp", "webhook timestamp is outside the replay window")
	}
	secret, ok := secrets[keyID]
	if !webhookKeyID.MatchString(keyID) || !ok {
		return nil, webhookError("unknown_key", "webhook key id is not one this receiver holds")
	}
	if !webhookSignature.MatchString(signature) {
		return nil, webhookError("invalid_signature", "webhook signature is malformed")
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write(payload)
	given, _ := hex.DecodeString(signature[3:])
	if !hmac.Equal(mac.Sum(nil), given) {
		return nil, webhookError("invalid_signature", "webhook signature does not match")
	}
	var body struct {
		ID        *string         `json:"id"`
		Type      string          `json:"type"`
		CreatedAt *string         `json:"created_at"`
		Test      bool            `json:"test"`
		Data      json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(payload, &body); err != nil {
		return nil, webhookError("invalid_payload", "webhook body is not JSON")
	}
	if body.ID == nil || *body.ID != id || !webhookEventID.MatchString(*body.ID) {
		return nil, webhookError("invalid_payload", "webhook event id does not match NR-Webhook-Id")
	}
	if body.CreatedAt == nil || !bytes.HasPrefix(bytes.TrimSpace(body.Data), []byte("{")) {
		return nil, webhookError("invalid_payload", "webhook event envelope is malformed")
	}
	if _, err := time.Parse(time.RFC3339Nano, *body.CreatedAt); err != nil {
		return nil, webhookError("invalid_payload", "webhook event envelope is malformed")
	}
	data, err := eventData(body.Type, body.Data)
	if err != nil {
		return nil, err
	}
	return &Event{ID: id, Type: body.Type, CreatedAt: *body.CreatedAt, Test: body.Test && body.Type == EventAssessmentUpdated, Data: data}, nil
}

func eventData(eventType string, raw json.RawMessage) (any, error) {
	switch eventType {
	case EventEndpointVerification:
		var verification EndpointVerification
		if err := json.Unmarshal(raw, &verification); err != nil ||
			!webhookEndpointID.MatchString(verification.EndpointID) ||
			!webhookChallenge.MatchString(verification.Challenge) {
			return nil, webhookError("invalid_payload", "endpoint verification event is malformed")
		}
		return &verification, nil
	case EventAccountDeleted:
		var deleted AccountDeleted
		if err := json.Unmarshal(raw, &deleted); err != nil || !webhookAccountID.MatchString(deleted.AccountID) || deleted.AssessmentsDeleted < 1 {
			return nil, webhookError("invalid_payload", "account deleted event is malformed")
		}
		if _, err := time.Parse(time.RFC3339Nano, deleted.DeletedAt); err != nil {
			return nil, webhookError("invalid_payload", "account deleted event is malformed")
		}
		return &deleted, nil
	case EventWorkspaceDeleted:
		var deleted WorkspaceDeleted
		if err := json.Unmarshal(raw, &deleted); err != nil || !webhookWorkspaceID.MatchString(deleted.WorkspaceID) ||
			(deleted.Environment != "live" && deleted.Environment != "test") || !webhookTime(deleted.DeletedAt) {
			return nil, webhookError("invalid_payload", "tenant deleted event is malformed")
		}
		return &deleted, nil
	case EventEndpointDisabled:
		var disabled EndpointDisabled
		if err := json.Unmarshal(raw, &disabled); err != nil || !webhookEndpointID.MatchString(disabled.EndpointID) ||
			!webhookTime(disabled.FailingSince) || !webhookTime(disabled.DisabledAt) {
			return nil, webhookError("invalid_payload", "endpoint disabled event is malformed")
		}
		return &disabled, nil
	case EventAccountRiskChanged:
		var changed AccountRiskChanged
		var envelope struct {
			Account                 json.RawMessage `json:"account"`
			PreviousDecisionsByFlow json.RawMessage `json:"previous_decisions_by_flow"`
		}
		var account struct {
			Risk            json.RawMessage `json:"risk"`
			DecisionsByFlow json.RawMessage `json:"decisions_by_flow"`
		}
		if json.Unmarshal(raw, &changed) != nil || json.Unmarshal(raw, &envelope) != nil ||
			!hasRequiredFields(raw, accountRiskChangedRequired) ||
			!hasRequiredFields(envelope.Account, accountRequired) || json.Unmarshal(envelope.Account, &account) != nil ||
			!hasRequiredFields(account.Risk, accountRiskRequired) ||
			!hasRequiredFields(account.DecisionsByFlow, decisionsByFlowRequired) ||
			!hasRequiredFields(envelope.PreviousDecisionsByFlow, decisionsByFlowRequired) ||
			!webhookAccountID.MatchString(changed.Account.ID) {
			return nil, webhookError("invalid_payload", "account risk changed event is malformed")
		}
		return &changed, nil
	case EventAssessmentUpdated:
		var probe struct {
			Status string `json:"status"`
		}
		_ = json.Unmarshal(raw, &probe)
		if probe.Status == "expired" {
			var expired AssessmentExpired
			if json.Unmarshal(raw, &expired) != nil ||
				!webhookAssessmentID.MatchString(expired.AssessmentID) ||
				!webhookAccountID.MatchString(expired.AccountID) || expired.AssessmentVersion < 1 ||
				expired.Status != "expired" || expired.ChangeReason != "expired" {
				return nil, webhookError("invalid_payload", "expired assessment event is malformed")
			}
			return &expired, nil
		}
		var assessment components.Assessment
		var envelope struct {
			Detail json.RawMessage `json:"detail"`
		}
		if json.Unmarshal(raw, &assessment) != nil || json.Unmarshal(raw, &envelope) != nil ||
			!hasRequiredFields(raw, assessmentRequired) || !hasRequiredFields(envelope.Detail, assessmentDetailRequired) ||
			!incompleteSourcesComplete(envelope.Detail) || !webhookAssessmentID.MatchString(assessment.ID) ||
			assessment.Detail.AssessmentID != assessment.ID || !webhookAccountID.MatchString(assessment.Detail.AccountID) {
			return nil, webhookError("invalid_payload", "webhook data is not an assessment")
		}
		return &assessment, nil
	}
	return nil, webhookError("invalid_payload", "unknown webhook event type")
}

func webhookTime(value string) bool {
	_, err := time.Parse(time.RFC3339Nano, value)
	return err == nil
}

// incompleteSourcesComplete checks each incomplete_sources entry names its
// lane and reason, which the generated decoder does not enforce. Values and
// extra keys are left open so a new lane or reason still verifies.
func incompleteSourcesComplete(raw json.RawMessage) bool {
	var body struct {
		IncompleteSources []map[string]json.RawMessage `json:"incomplete_sources"`
	}
	if json.Unmarshal(raw, &body) != nil {
		return false
	}
	for _, entry := range body.IncompleteSources {
		var lane, reason string
		if json.Unmarshal(entry["lane"], &lane) != nil || json.Unmarshal(entry["reason"], &reason) != nil || lane == "" || reason == "" {
			return false
		}
	}
	return true
}

// assessmentRequired and assessmentDetailRequired are every field
// components.Assessment and its detail always carry: the ones whose json tag
// is neither omitzero nor omitempty. The generated decoder does not check
// them, so a pointer such as {assessment_id, assessment_url} would otherwise
// decode as an assessment.
var (
	assessmentRequired       = requiredJSONFields(reflect.TypeFor[components.Assessment]())
	assessmentDetailRequired = requiredJSONFields(reflect.TypeFor[components.AssessmentDetail]())
	// The same for account.risk_changed: the event's data, its account (as
	// Accounts.Get returns it), the account's risk and each decisions_by_flow.
	accountRiskChangedRequired = requiredJSONFields(reflect.TypeFor[AccountRiskChanged]())
	accountRequired            = requiredJSONFields(reflect.TypeFor[components.Account]())
	accountRiskRequired        = requiredJSONFields(reflect.TypeFor[components.AccountRisk]())
	decisionsByFlowRequired    = requiredJSONFields(reflect.TypeFor[components.AccountDecisionsByFlow]())
)

func requiredJSONFields(t reflect.Type) []string {
	var names []string
	for i := range t.NumField() {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		name, options, _ := strings.Cut(tag, ",")
		if name == "" || name == "-" || strings.Contains(options, "omitzero") || strings.Contains(options, "omitempty") {
			continue
		}
		names = append(names, name)
	}
	return names
}

func hasRequiredFields(raw json.RawMessage, names []string) bool {
	var present map[string]json.RawMessage
	if json.Unmarshal(raw, &present) != nil {
		return false
	}
	for _, name := range names {
		if value, ok := present[name]; !ok || string(value) == "null" {
			return false
		}
	}
	return true
}

// webhookHeaderValues finds each NR-Webhook-* header exactly once, matching
// names without regard to case.
func webhookHeaderValues(headers http.Header) ([4]string, error) {
	var values [4]string
	var counts [4]int
	for name, received := range headers {
		for i, canonical := range webhookHeaders {
			if strings.EqualFold(name, canonical) {
				counts[i] += len(received)
				if len(received) > 0 {
					values[i] = received[0]
				}
			}
		}
	}
	for i, canonical := range webhookHeaders {
		if counts[i] > 1 {
			return values, webhookError("duplicate_header", "webhook header "+strings.ToLower(canonical)+" is repeated")
		}
		if values[i] == "" {
			return values, webhookError("missing_header", "webhook header "+strings.ToLower(canonical)+" is missing")
		}
	}
	return values, nil
}

// VerificationResponse is the body to answer an endpoint.verification event
// with, as JSON, status 200.
func VerificationResponse(event *Event) (map[string]string, error) {
	if event != nil {
		if verification, ok := event.Data.(*EndpointVerification); ok {
			return map[string]string{"challenge": verification.Challenge}, nil
		}
	}
	return nil, webhookError("invalid_payload", "not an endpoint.verification event")
}
