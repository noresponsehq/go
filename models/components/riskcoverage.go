package components

// RiskCoverage - Sources counted by state, not named.
type RiskCoverage struct {
	// Sources that gave a definite answer.
	Answered int64 `json:"answered"`
	// Sources that could not answer.
	Failed int64 `json:"failed"`
	// Sources still being asked.
	Pending int64 `json:"pending"`
}

func (r *RiskCoverage) GetAnswered() int64 {
	if r == nil {
		return 0
	}
	return r.Answered
}

func (r *RiskCoverage) GetFailed() int64 {
	if r == nil {
		return 0
	}
	return r.Failed
}

func (r *RiskCoverage) GetPending() int64 {
	if r == nil {
		return 0
	}
	return r.Pending
}
