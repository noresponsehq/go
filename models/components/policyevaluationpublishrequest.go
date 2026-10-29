package components

type PolicyEvaluationPublishRequest struct {
	ExpectedPolicyVersion int64      `json:"expected_policy_version"`
	ManifestSha256        string     `json:"manifest_sha256"`
	InputSha256           string     `json:"input_sha256"`
	ResultSha256          string     `json:"result_sha256"`
	Candidate             Thresholds `json:"candidate"`
	ChangeReason          string     `json:"change_reason"`
}

func (p *PolicyEvaluationPublishRequest) GetExpectedPolicyVersion() int64 {
	if p == nil {
		return 0
	}
	return p.ExpectedPolicyVersion
}

func (p *PolicyEvaluationPublishRequest) GetManifestSha256() string {
	if p == nil {
		return ""
	}
	return p.ManifestSha256
}

func (p *PolicyEvaluationPublishRequest) GetInputSha256() string {
	if p == nil {
		return ""
	}
	return p.InputSha256
}

func (p *PolicyEvaluationPublishRequest) GetResultSha256() string {
	if p == nil {
		return ""
	}
	return p.ResultSha256
}

func (p *PolicyEvaluationPublishRequest) GetCandidate() Thresholds {
	if p == nil {
		return Thresholds{}
	}
	return p.Candidate
}

func (p *PolicyEvaluationPublishRequest) GetChangeReason() string {
	if p == nil {
		return ""
	}
	return p.ChangeReason
}
