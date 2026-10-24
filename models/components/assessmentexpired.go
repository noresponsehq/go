package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

// AssessmentExpired - Terminal. Nothing to fetch; the consumer deletes its state for assessment_id.
type AssessmentExpired struct {
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the workspace; another workspace's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// The last served version: expiry is not a revision.
	AssessmentVersion int64 `json:"assessment_version"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	status string `const:"expired" json:"status"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	changeReason string `const:"expired" json:"change_reason"`
}

func (a AssessmentExpired) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentExpired) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentExpired) GetAssessmentID() string {
	if a == nil {
		return ""
	}
	return a.AssessmentID
}

func (a *AssessmentExpired) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AssessmentExpired) GetAssessmentVersion() int64 {
	if a == nil {
		return 0
	}
	return a.AssessmentVersion
}

func (a *AssessmentExpired) GetStatus() string {
	return "expired"
}

func (a *AssessmentExpired) GetChangeReason() string {
	return "expired"
}
