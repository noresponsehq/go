package components

type ReviewResolveRequest struct {
	ExpectedVersion     int64      `json:"expected_version"`
	ExpectedRiskVersion int64      `json:"expected_risk_version"`
	Conclusion          Conclusion `json:"conclusion"`
	Note                string     `json:"note"`
}

func (r *ReviewResolveRequest) GetExpectedVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ExpectedVersion
}

func (r *ReviewResolveRequest) GetExpectedRiskVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ExpectedRiskVersion
}

func (r *ReviewResolveRequest) GetConclusion() Conclusion {
	if r == nil {
		return Conclusion("")
	}
	return r.Conclusion
}

func (r *ReviewResolveRequest) GetNote() string {
	if r == nil {
		return ""
	}
	return r.Note
}
