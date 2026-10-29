package components

type Resolution struct {
	ReviewVersion     int64      `json:"review_version"`
	RiskVersion       int64      `json:"risk_version"`
	AssessmentVersion int64      `json:"assessment_version"`
	AssessmentID      string     `json:"assessment_id"`
	Conclusion        Conclusion `json:"conclusion"`
	Note              string     `json:"note"`
	CreatedBy         string     `json:"created_by"`
	CreatedAt         string     `json:"created_at"`
}

func (r *Resolution) GetReviewVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ReviewVersion
}

func (r *Resolution) GetRiskVersion() int64 {
	if r == nil {
		return 0
	}
	return r.RiskVersion
}

func (r *Resolution) GetAssessmentVersion() int64 {
	if r == nil {
		return 0
	}
	return r.AssessmentVersion
}

func (r *Resolution) GetAssessmentID() string {
	if r == nil {
		return ""
	}
	return r.AssessmentID
}

func (r *Resolution) GetConclusion() Conclusion {
	if r == nil {
		return Conclusion("")
	}
	return r.Conclusion
}

func (r *Resolution) GetNote() string {
	if r == nil {
		return ""
	}
	return r.Note
}

func (r *Resolution) GetCreatedBy() string {
	if r == nil {
		return ""
	}
	return r.CreatedBy
}

func (r *Resolution) GetCreatedAt() string {
	if r == nil {
		return ""
	}
	return r.CreatedAt
}
