package components

type PolicyEvaluationRequestCosts struct {
	Legitimate Cost `json:"legitimate"`
	Abuse      Cost `json:"abuse"`
}

func (p *PolicyEvaluationRequestCosts) GetLegitimate() Cost {
	if p == nil {
		return Cost{}
	}
	return p.Legitimate
}

func (p *PolicyEvaluationRequestCosts) GetAbuse() Cost {
	if p == nil {
		return Cost{}
	}
	return p.Abuse
}
