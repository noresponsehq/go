package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AccountDeletion - Body of the 200 answer to DELETE /v1/accounts/{account_id}: the calling workspace's account was erased in the environment of the calling key. Every assessment of the account went, with its revisions, evidence, outcomes and the identity-graph marks its outcomes set, and an assessment that was still running when the call arrived is stopped. Afterwards GET /v1/accounts/{account_id} and GET /v1/assessments/{assessment_id} of each erased assessment answer 404, and the history page is empty. The call is idempotent: a repeat, or an account with nothing to erase, answers 200 with assessments_deleted 0. The account_id can be assessed again later, as a new account.
type AccountDeletion struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// The account_id that was erased.
	ID string `json:"id"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	deleted bool `const:"true" json:"deleted"`
	// How many assessments this call deleted. 0 when the account had none left, so a repeat answers 0.
	AssessmentsDeleted int64 `json:"assessments_deleted"`
	// When the erasure took effect. An assessment of the account begun before this time is refused or stopped.
	DeletedAt string `json:"deleted_at"`
	// Receipt for GET /v1/erasures/{erasure_id}, when durable storage issued one.
	ErasureID *string `json:"erasure_id,omitzero"`
}

func (a AccountDeletion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDeletion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDeletion) GetSchemaVersion() string {
	return "1"
}

func (a *AccountDeletion) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AccountDeletion) GetDeleted() bool {
	return true
}

func (a *AccountDeletion) GetAssessmentsDeleted() int64 {
	if a == nil {
		return 0
	}
	return a.AssessmentsDeleted
}

func (a *AccountDeletion) GetDeletedAt() string {
	if a == nil {
		return ""
	}
	return a.DeletedAt
}

func (a *AccountDeletion) GetErasureID() *string {
	if a == nil {
		return nil
	}
	return a.ErasureID
}
