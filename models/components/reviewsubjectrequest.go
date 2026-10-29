package components

// ReviewSubjectRequest - An assessment link also requires assessment_id and assessment_version; an issued_credential link also requires credential_id. Every reference must belong to account_id in this environment. Other kinds reject those fields.
type ReviewSubjectRequest struct {
	Kind              ReviewSubjectKind `json:"kind"`
	AccountID         string            `json:"account_id"`
	AssessmentID      *string           `json:"assessment_id,omitzero"`
	AssessmentVersion *int64            `json:"assessment_version,omitzero"`
	CredentialID      *string           `json:"credential_id,omitzero"`
	Reason            string            `json:"reason"`
	ExpectedVersion   int64             `json:"expected_version"`
}

func (r *ReviewSubjectRequest) GetKind() ReviewSubjectKind {
	if r == nil {
		return ReviewSubjectKind("")
	}
	return r.Kind
}

func (r *ReviewSubjectRequest) GetAccountID() string {
	if r == nil {
		return ""
	}
	return r.AccountID
}

func (r *ReviewSubjectRequest) GetAssessmentID() *string {
	if r == nil {
		return nil
	}
	return r.AssessmentID
}

func (r *ReviewSubjectRequest) GetAssessmentVersion() *int64 {
	if r == nil {
		return nil
	}
	return r.AssessmentVersion
}

func (r *ReviewSubjectRequest) GetCredentialID() *string {
	if r == nil {
		return nil
	}
	return r.CredentialID
}

func (r *ReviewSubjectRequest) GetReason() string {
	if r == nil {
		return ""
	}
	return r.Reason
}

func (r *ReviewSubjectRequest) GetExpectedVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ExpectedVersion
}
