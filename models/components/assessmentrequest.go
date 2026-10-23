package components

import (
	"github.com/noresponsehq/go/v5/internal/utils"
)

// AssessmentRequest - POST /v1/assessments body. At least one of email, username or phone is required; name and ip only refine an assessment. No field carries a password, hash, token or breach value, and unknown members are refused. The email local part is kept as received and only the domain is lower-cased; the username is compared case-insensitively; the phone is strict E.164 and is assessed asynchronously. The whole body is capped at 32768 bytes. client_context holds bounded hints from the client of the user; auth_context holds normalised claims from a session your backend has verified. Raw tokens are refused.
type AssessmentRequest struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// How long the synchronous answer waits for sources. quick waits only for the fast sources, those that answer within 500 ms of ingress (or deadline_ms when lower); full waits for the whole budget, up to 3000 ms. Sources still running at the cut continue asynchronously and update the assessment either way.
	Depth *Depth `json:"depth,omitzero"`
	// At least one of email, username or phone. name and ip refine the match and travel beside a lookup identifier; a body with only name or ip is refused with 400.
	Identifiers AssessmentRequestIdentifiers `json:"identifiers"`
	// Optional policy inputs: region.
	PolicyContext *AssessmentRequestPolicyContext `json:"policy_context,omitzero"`
	// An optional provider correlation id, echoed back. It is not an idempotency key. Refuses an exact-shape JWT or a bare 40, 64 or 128 character hex digest; a UUID is accepted.
	RequestID *string `json:"request_id,omitzero"`
	// Optional tighter budget, 50 to 3000 ms, measured from the moment the request reaches NoResponse. Absent means the 3000 ms ceiling; a value above it is refused with 400 rather than clamped.
	DeadlineMs *int64 `json:"deadline_ms,omitzero"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode *ResponseMode `json:"response_mode,omitzero"`
	// Optional context observed by the client of the user. The server parses these hints into coarse features, leaves raw user_agent and opaque ids out of every answer, and weighs a mismatch with the other evidence.
	ClientContext *AssessmentRequestClientContext `json:"client_context,omitzero"`
	// Normalised claims from a session your backend has already verified with WorkOS, Clerk, Auth0 or any OIDC provider. Pass the output of the SDK's auth adapters, not a raw token.
	AuthContext *AuthContext `json:"auth_context,omitzero"`
	// Optional. When your platform issues an API key to the new account, send SHA-256 of that key, so later usage of the key, or a leaked copy, ties back to this assessment. Each SDK has a helper that computes it. Send only the hash: customer keys are high-entropy, so the hash reveals nothing about them. The assessment stores it, and no response, webhook or log carries it.
	IssuedCredential *AssessmentRequestIssuedCredential `json:"issued_credential,omitzero"`
}

func (a AssessmentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentRequest) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AssessmentRequest) GetFlow() Flow {
	if a == nil {
		return Flow("")
	}
	return a.Flow
}

func (a *AssessmentRequest) GetDepth() *Depth {
	if a == nil {
		return nil
	}
	return a.Depth
}

func (a *AssessmentRequest) GetIdentifiers() AssessmentRequestIdentifiers {
	if a == nil {
		return AssessmentRequestIdentifiers{}
	}
	return a.Identifiers
}

func (a *AssessmentRequest) GetPolicyContext() *AssessmentRequestPolicyContext {
	if a == nil {
		return nil
	}
	return a.PolicyContext
}

func (a *AssessmentRequest) GetRequestID() *string {
	if a == nil {
		return nil
	}
	return a.RequestID
}

func (a *AssessmentRequest) GetDeadlineMs() *int64 {
	if a == nil {
		return nil
	}
	return a.DeadlineMs
}

func (a *AssessmentRequest) GetResponseMode() *ResponseMode {
	if a == nil {
		return nil
	}
	return a.ResponseMode
}

func (a *AssessmentRequest) GetClientContext() *AssessmentRequestClientContext {
	if a == nil {
		return nil
	}
	return a.ClientContext
}

func (a *AssessmentRequest) GetAuthContext() *AuthContext {
	if a == nil {
		return nil
	}
	return a.AuthContext
}

func (a *AssessmentRequest) GetIssuedCredential() *AssessmentRequestIssuedCredential {
	if a == nil {
		return nil
	}
	return a.IssuedCredential
}
