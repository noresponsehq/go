package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// PolicyPreviewRequest - Side-effect-free threshold preview for at most 31 days. Retained first answers are evaluated at their exact frozen scores using the production threshold and hold rules. Changes to admission, source selection, depth or retention are not simulated.
type PolicyPreviewRequest struct {
	ExpectedPolicyVersion int64                     `json:"expected_policy_version"`
	From                  string                    `json:"from"`
	To                    string                    `json:"to"`
	Limit                 *int64                    `default:"1000" json:"limit"`
	Flows                 PolicyPreviewRequestFlows `json:"flows"`
}

func (p PolicyPreviewRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyPreviewRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PolicyPreviewRequest) GetExpectedPolicyVersion() int64 {
	if p == nil {
		return 0
	}
	return p.ExpectedPolicyVersion
}

func (p *PolicyPreviewRequest) GetFrom() string {
	if p == nil {
		return ""
	}
	return p.From
}

func (p *PolicyPreviewRequest) GetTo() string {
	if p == nil {
		return ""
	}
	return p.To
}

func (p *PolicyPreviewRequest) GetLimit() *int64 {
	if p == nil {
		return nil
	}
	return p.Limit
}

func (p *PolicyPreviewRequest) GetFlows() PolicyPreviewRequestFlows {
	if p == nil {
		return PolicyPreviewRequestFlows{}
	}
	return p.Flows
}
