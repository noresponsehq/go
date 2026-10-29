package components

type PolicyEvaluationRequestLimits struct {
	ReviewRate            float64 `json:"review_rate"`
	ServerErrorRate       float64 `json:"server_error_rate"`
	UnavailableResultRate float64 `json:"unavailable_result_rate"`
	P95Ms                 float64 `json:"p95_ms"`
}

func (p *PolicyEvaluationRequestLimits) GetReviewRate() float64 {
	if p == nil {
		return 0.0
	}
	return p.ReviewRate
}

func (p *PolicyEvaluationRequestLimits) GetServerErrorRate() float64 {
	if p == nil {
		return 0.0
	}
	return p.ServerErrorRate
}

func (p *PolicyEvaluationRequestLimits) GetUnavailableResultRate() float64 {
	if p == nil {
		return 0.0
	}
	return p.UnavailableResultRate
}

func (p *PolicyEvaluationRequestLimits) GetP95Ms() float64 {
	if p == nil {
		return 0.0
	}
	return p.P95Ms
}
