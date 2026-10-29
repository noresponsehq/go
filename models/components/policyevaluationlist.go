package components

// PolicyEvaluationList - Retained jobs, newest first. Follow next_before for another page.
type PolicyEvaluationList struct {
	Items      []PolicyEvaluation `json:"items"`
	NextBefore *string            `json:"next_before,omitzero"`
}

func (p *PolicyEvaluationList) GetItems() []PolicyEvaluation {
	if p == nil {
		return []PolicyEvaluation{}
	}
	return p.Items
}

func (p *PolicyEvaluationList) GetNextBefore() *string {
	if p == nil {
		return nil
	}
	return p.NextBefore
}
