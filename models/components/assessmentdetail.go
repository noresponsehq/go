package components

import (
	"github.com/noresponsehq/go/v3/internal/utils"
)

// AssessmentDetail - One assessment's own facts: its score and recommendation, the signals behind them, the lanes still running and its timing. risk here is this assessment's own score; the account's current risk is the account member of the assessment. risk and recommendation are present for complete and partial; assessment_url is present for pending.
type AssessmentDetail struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the tenant; another tenant's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// Your own opaque identifier for the account being checked, scoped to your tenant. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// An optional provider correlation id, echoed back. It is not an idempotency key. Refuses an exact-shape JWT or a bare 40, 64 or 128 character hex digest; a UUID is accepted.
	RequestID *string `json:"request_id,omitzero"`
	// complete: final; nothing is still running. A lane that failed or missed its 10 s deadline stays listed in incomplete_sources, and confidence is lower. partial: a recommendation exists and at least one lane is still running or did not settle (see incomplete_sources). pending: nothing usable before the deadline; no recommendation yet, poll or wait for the webhook.
	Status Status `json:"status"`
	// The recommended action: allow, step_up or review. Your application applies it.
	Recommendation *Recommendation `json:"recommendation,omitzero"`
	// Ordinal confidence in the recommendation under the current policy: low, medium, high or unavailable.
	Confidence Confidence `json:"confidence"`
	// The one risk decision for this event. score is 0 (safe) to 100 (risky): the sum of the weights of the reasons that fired, clamped to 0 to 100, from hand-set versioned weights. Every answer with a recommendation carries it, and the top-level recommendation equals risk.recommendation. It names no source or provider.
	Risk *Risk `json:"risk,omitzero"`
	// Detailed view only (response_mode=detailed, which needs the detailed scope): a standard answer omits signals[], and a detailed one always carries it, possibly empty. The body does not name its mode, so that rule is the serving gate's; the schema states the half it can: an answer with risk.details is a detailed view and carries signals[]. At most one signal per (type, source_id, identifier).
	Signals []Signal `json:"signals,omitzero"`
	// The lanes that did not settle, each listed once with a reason; lanes, not sources. In a detailed view a lane is listed exactly when one of its signals did not settle; a standard view carries no signals to match.
	IncompleteSources []IncompleteSource `json:"incomplete_sources"`
	// Version of the hand-authored heuristic weights that produced the recommendation, for example 'risk-2026-10.7'. A tenant threshold override adds a +t<N> suffix, e.g. 'risk-2026-10.7+t3'.
	PolicyVersion string `json:"policy_version"`
	// Labels the v0 policy as hand-authored and not statistically validated.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	policyBasis string `const:"heuristic_unvalidated" json:"policy_basis"`
	// Monotonic revision of this assessment. 1 on the synchronous answer; each material change while the assessment is served (an async continuation settling a source, a change of recommendation, confidence or status) increments it. Expiry does not increment it: an expired assessment is gone, not revised, and the expiry webhook carries the last served version.
	AssessmentVersion int64 `json:"assessment_version"`
	// When the assessment was created. Fixed for the life of the assessment: every revision carries the same created_at.
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// After this instant the assessment is no longer served: GET /v1/assessments/{assessment_id} answers 404 and it leaves the account history. It is created_at plus your retention period at creation (default and maximum 365 days) and is fixed per record: a later revision keeps it, and a new assessment for the same account starts its own clock. If you later shorten retention, the effective expiry becomes created_at plus the shorter of the two periods.
	ExpiresAt string `json:"expires_at"`
	// Server-relative path of the authoritative assessment, GET /v1/assessments/{assessment_id}.
	AssessmentURL *string                `json:"assessment_url,omitzero"`
	Timing        AssessmentDetailTiming `json:"timing"`
	// The identity fields of the auth context the assessment was made with, echoed for the tenant's own correlation: the request's auth_context, or the one a hosted provider hook (POST /hooks/{integration_id}) built, with transport provider_hook and its hook_point. Present exactly when the request carried an auth_context, and the same on every revision of the assessment. Values are as the request or the hook sent them; tokens, timestamps, roles and scopes are not echoed, and the echo does not change the score.
	AuthContext *AssessmentDetailAuthContext `json:"auth_context,omitzero"`
}

func (a AssessmentDetail) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentDetail) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentDetail) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentDetail) GetAssessmentID() string {
	if a == nil {
		return ""
	}
	return a.AssessmentID
}

func (a *AssessmentDetail) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AssessmentDetail) GetRequestID() *string {
	if a == nil {
		return nil
	}
	return a.RequestID
}

func (a *AssessmentDetail) GetStatus() Status {
	if a == nil {
		return Status("")
	}
	return a.Status
}

func (a *AssessmentDetail) GetRecommendation() *Recommendation {
	if a == nil {
		return nil
	}
	return a.Recommendation
}

func (a *AssessmentDetail) GetConfidence() Confidence {
	if a == nil {
		return Confidence("")
	}
	return a.Confidence
}

func (a *AssessmentDetail) GetRisk() *Risk {
	if a == nil {
		return nil
	}
	return a.Risk
}

func (a *AssessmentDetail) GetSignals() []Signal {
	if a == nil {
		return nil
	}
	return a.Signals
}

func (a *AssessmentDetail) GetIncompleteSources() []IncompleteSource {
	if a == nil {
		return []IncompleteSource{}
	}
	return a.IncompleteSources
}

func (a *AssessmentDetail) GetPolicyVersion() string {
	if a == nil {
		return ""
	}
	return a.PolicyVersion
}

func (a *AssessmentDetail) GetPolicyBasis() string {
	return "heuristic_unvalidated"
}

func (a *AssessmentDetail) GetAssessmentVersion() int64 {
	if a == nil {
		return 0
	}
	return a.AssessmentVersion
}

func (a *AssessmentDetail) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AssessmentDetail) GetUpdatedAt() string {
	if a == nil {
		return ""
	}
	return a.UpdatedAt
}

func (a *AssessmentDetail) GetExpiresAt() string {
	if a == nil {
		return ""
	}
	return a.ExpiresAt
}

func (a *AssessmentDetail) GetAssessmentURL() *string {
	if a == nil {
		return nil
	}
	return a.AssessmentURL
}

func (a *AssessmentDetail) GetTiming() AssessmentDetailTiming {
	if a == nil {
		return AssessmentDetailTiming{}
	}
	return a.Timing
}

func (a *AssessmentDetail) GetAuthContext() *AssessmentDetailAuthContext {
	if a == nil {
		return nil
	}
	return a.AuthContext
}
