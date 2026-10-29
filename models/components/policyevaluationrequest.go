package components

type PolicyEvaluationRequest struct {
	ExpectedPolicyVersion int64                         `json:"expected_policy_version"`
	From                  string                        `json:"from"`
	To                    string                        `json:"to"`
	Candidates            []Thresholds                  `json:"candidates"`
	Costs                 PolicyEvaluationRequestCosts  `json:"costs"`
	Limits                PolicyEvaluationRequestLimits `json:"limits"`
	LabelCollection       string                        `json:"label_collection"`
	SplitSeed             string                        `json:"split_seed"`
}

func (p *PolicyEvaluationRequest) GetExpectedPolicyVersion() int64 {
	if p == nil {
		return 0
	}
	return p.ExpectedPolicyVersion
}

func (p *PolicyEvaluationRequest) GetFrom() string {
	if p == nil {
		return ""
	}
	return p.From
}

func (p *PolicyEvaluationRequest) GetTo() string {
	if p == nil {
		return ""
	}
	return p.To
}

func (p *PolicyEvaluationRequest) GetCandidates() []Thresholds {
	if p == nil {
		return []Thresholds{}
	}
	return p.Candidates
}

func (p *PolicyEvaluationRequest) GetCosts() PolicyEvaluationRequestCosts {
	if p == nil {
		return PolicyEvaluationRequestCosts{}
	}
	return p.Costs
}

func (p *PolicyEvaluationRequest) GetLimits() PolicyEvaluationRequestLimits {
	if p == nil {
		return PolicyEvaluationRequestLimits{}
	}
	return p.Limits
}

func (p *PolicyEvaluationRequest) GetLabelCollection() string {
	if p == nil {
		return ""
	}
	return p.LabelCollection
}

func (p *PolicyEvaluationRequest) GetSplitSeed() string {
	if p == nil {
		return ""
	}
	return p.SplitSeed
}
