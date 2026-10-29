package components

// ReviewSubject - A tenant-local investigation link. risk_version is the account risk version captured when the link was added. It does not assert a relationship or change feedback or enforcement.
type ReviewSubject struct {
	ID                string            `json:"id"`
	Kind              ReviewSubjectKind `json:"kind"`
	AccountID         string            `json:"account_id"`
	AssessmentID      *string           `json:"assessment_id,omitzero"`
	AssessmentVersion *int64            `json:"assessment_version,omitzero"`
	CredentialID      *string           `json:"credential_id,omitzero"`
	RiskVersion       int64             `json:"risk_version"`
	Reason            string            `json:"reason"`
	CreatedBy         string            `json:"created_by"`
	CreatedAt         string            `json:"created_at"`
}

func (r *ReviewSubject) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *ReviewSubject) GetKind() ReviewSubjectKind {
	if r == nil {
		return ReviewSubjectKind("")
	}
	return r.Kind
}

func (r *ReviewSubject) GetAccountID() string {
	if r == nil {
		return ""
	}
	return r.AccountID
}

func (r *ReviewSubject) GetAssessmentID() *string {
	if r == nil {
		return nil
	}
	return r.AssessmentID
}

func (r *ReviewSubject) GetAssessmentVersion() *int64 {
	if r == nil {
		return nil
	}
	return r.AssessmentVersion
}

func (r *ReviewSubject) GetCredentialID() *string {
	if r == nil {
		return nil
	}
	return r.CredentialID
}

func (r *ReviewSubject) GetRiskVersion() int64 {
	if r == nil {
		return 0
	}
	return r.RiskVersion
}

func (r *ReviewSubject) GetReason() string {
	if r == nil {
		return ""
	}
	return r.Reason
}

func (r *ReviewSubject) GetCreatedBy() string {
	if r == nil {
		return ""
	}
	return r.CreatedBy
}

func (r *ReviewSubject) GetCreatedAt() string {
	if r == nil {
		return ""
	}
	return r.CreatedAt
}
