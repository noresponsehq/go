package components

import (
	"github.com/noresponsehq/go/v5/internal/utils"
)

// AssessmentDeletion - Body of the 200 answer to DELETE /v1/assessments/{assessment_id}: one of the calling workspace's assessments was deleted, with its versions, evidence, outcomes and the identity-graph marks its outcomes set. Afterwards GET /v1/assessments/{assessment_id} answers 404 and the assessment leaves the account's history. The account is scored again from what remains: account.current_assessment_id names the assessment GET /v1/accounts/{account_id} now serves its risk from, and is absent when the account has no assessment with a decision left. Deleting an assessment that is already gone answers 404.
type AssessmentDeletion struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// The assessment that was deleted.
	ID string `json:"id"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	deleted bool                      `const:"true" json:"deleted"`
	Account AssessmentDeletionAccount `json:"account"`
	// When the deletion took effect.
	DeletedAt string `json:"deleted_at"`
}

func (a AssessmentDeletion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentDeletion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentDeletion) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentDeletion) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AssessmentDeletion) GetDeleted() bool {
	return true
}

func (a *AssessmentDeletion) GetAccount() AssessmentDeletionAccount {
	if a == nil {
		return AssessmentDeletionAccount{}
	}
	return a.Account
}

func (a *AssessmentDeletion) GetDeletedAt() string {
	if a == nil {
		return ""
	}
	return a.DeletedAt
}
