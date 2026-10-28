package components

// AccountBinding - Binding appends an assessment revision and preserves earlier versions. Continuations use the same assessment. Erasing either the provisional or canonical name erases the canonical account, including the bound assessment.
type AccountBinding struct {
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the workspace; another workspace's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	ProvisionalAccountID string `json:"provisional_account_id"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// Monotonic revision of this assessment. 1 on the synchronous answer; each material change while the assessment is served (an async continuation settling a source, a change of recommendation, confidence or status) increments it. Expiry does not increment it: an expired assessment is gone, not revised, and the expiry webhook carries the last served version.
	AssessmentVersion int64               `json:"assessment_version"`
	CreatedAt         string              `json:"created_at"`
	Actor             AccountBindingActor `json:"actor"`
}

func (a *AccountBinding) GetAssessmentID() string {
	if a == nil {
		return ""
	}
	return a.AssessmentID
}

func (a *AccountBinding) GetProvisionalAccountID() string {
	if a == nil {
		return ""
	}
	return a.ProvisionalAccountID
}

func (a *AccountBinding) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountBinding) GetAssessmentVersion() int64 {
	if a == nil {
		return 0
	}
	return a.AssessmentVersion
}

func (a *AccountBinding) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AccountBinding) GetActor() AccountBindingActor {
	if a == nil {
		return AccountBindingActor{}
	}
	return a.Actor
}
