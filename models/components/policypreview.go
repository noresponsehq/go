package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// PolicyPreview - Exact threshold transitions among eligible retained first answers, newest first within the limit. Partial older revisions without a frozen hold snapshot are reported as unsupported. Preview leaves account state, outcomes and source evidence unchanged.
type PolicyPreview struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	scope         string                    `const:"retained_first_answer_thresholds" json:"scope"`
	PolicyVersion int64                     `json:"policy_version"`
	From          string                    `json:"from"`
	To            string                    `json:"to"`
	RetentionFrom string                    `json:"retention_from"`
	Coverage      PolicyPreviewCoverage     `json:"coverage"`
	Changed       int64                     `json:"changed"`
	Unchanged     int64                     `json:"unchanged"`
	Transitions   []PolicyPreviewTransition `json:"transitions"`
}

func (p PolicyPreview) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyPreview) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PolicyPreview) GetScope() string {
	return "retained_first_answer_thresholds"
}

func (p *PolicyPreview) GetPolicyVersion() int64 {
	if p == nil {
		return 0
	}
	return p.PolicyVersion
}

func (p *PolicyPreview) GetFrom() string {
	if p == nil {
		return ""
	}
	return p.From
}

func (p *PolicyPreview) GetTo() string {
	if p == nil {
		return ""
	}
	return p.To
}

func (p *PolicyPreview) GetRetentionFrom() string {
	if p == nil {
		return ""
	}
	return p.RetentionFrom
}

func (p *PolicyPreview) GetCoverage() PolicyPreviewCoverage {
	if p == nil {
		return PolicyPreviewCoverage{}
	}
	return p.Coverage
}

func (p *PolicyPreview) GetChanged() int64 {
	if p == nil {
		return 0
	}
	return p.Changed
}

func (p *PolicyPreview) GetUnchanged() int64 {
	if p == nil {
		return 0
	}
	return p.Unchanged
}

func (p *PolicyPreview) GetTransitions() []PolicyPreviewTransition {
	if p == nil {
		return []PolicyPreviewTransition{}
	}
	return p.Transitions
}
