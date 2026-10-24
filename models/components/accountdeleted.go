package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

// AccountDeleted - Terminal for the account's assessments created at or before deleted_at in the receiving endpoint's environment: delete what you hold from them and ignore later-arriving updates of them. deleted_at is no earlier than the creation of any assessment it erased. An assessment created after it belongs to the account assessed again: keep it, even when this event arrives after it.
type AccountDeleted struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// How many assessments the erasure deleted.
	AssessmentsDeleted int64 `json:"assessments_deleted"`
	// When the erasure took effect, equal to deleted_at in the DELETE answer.
	DeletedAt string `json:"deleted_at"`
}

func (a AccountDeleted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDeleted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDeleted) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountDeleted) GetAssessmentsDeleted() int64 {
	if a == nil {
		return 0
	}
	return a.AssessmentsDeleted
}

func (a *AccountDeleted) GetDeletedAt() string {
	if a == nil {
		return ""
	}
	return a.DeletedAt
}
