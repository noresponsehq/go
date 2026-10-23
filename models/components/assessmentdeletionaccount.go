package components

type AssessmentDeletionAccount struct {
	// The account_id of the deleted assessment.
	ID string `json:"id"`
	// The assessment the account's current risk now comes from. Absent when the account has no assessment with a decision left, and GET /v1/accounts/{account_id} then answers 404.
	CurrentAssessmentID *string `json:"current_assessment_id,omitzero"`
}

func (a *AssessmentDeletionAccount) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AssessmentDeletionAccount) GetCurrentAssessmentID() *string {
	if a == nil {
		return nil
	}
	return a.CurrentAssessmentID
}
