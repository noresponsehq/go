package components

// PolicyRestoreRequest - Publish a captured configuration as a new version. Version checks and publication are atomic. Longer retention does not revive expired records, and restored policy does not undo customer enforcement.
type PolicyRestoreRequest struct {
	Version               int64  `json:"version"`
	ExpectedPolicyVersion int64  `json:"expected_policy_version"`
	ChangeReason          string `json:"change_reason"`
}

func (p *PolicyRestoreRequest) GetVersion() int64 {
	if p == nil {
		return 0
	}
	return p.Version
}

func (p *PolicyRestoreRequest) GetExpectedPolicyVersion() int64 {
	if p == nil {
		return 0
	}
	return p.ExpectedPolicyVersion
}

func (p *PolicyRestoreRequest) GetChangeReason() string {
	if p == nil {
		return ""
	}
	return p.ChangeReason
}
