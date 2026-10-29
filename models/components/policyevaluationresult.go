package components

// PolicyEvaluationResult - The evaluation report and its provenance. A no_change report cannot authorise publication.
type PolicyEvaluationResult struct {
	Evaluation PolicyEvaluation `json:"evaluation"`
	Report     Report           `json:"report"`
}

func (p *PolicyEvaluationResult) GetEvaluation() PolicyEvaluation {
	if p == nil {
		return PolicyEvaluation{}
	}
	return p.Evaluation
}

func (p *PolicyEvaluationResult) GetReport() Report {
	if p == nil {
		return Report{}
	}
	return p.Report
}
