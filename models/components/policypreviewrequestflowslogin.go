package components

import (
	"github.com/noresponsehq/go/v9/optionalnullable"
)

type PolicyPreviewRequestFlowsLogin struct {
	StepUp *int64                                   `json:"step_up,omitzero"`
	Review *int64                                   `json:"review,omitzero"`
	Block  optionalnullable.OptionalNullable[int64] `json:"block,omitzero"`
	// Send true; false is refused.
	InheritsAccountCreation *bool `json:"inherits_account_creation,omitzero"`
}

func (p *PolicyPreviewRequestFlowsLogin) GetStepUp() *int64 {
	if p == nil {
		return nil
	}
	return p.StepUp
}

func (p *PolicyPreviewRequestFlowsLogin) GetReview() *int64 {
	if p == nil {
		return nil
	}
	return p.Review
}

func (p *PolicyPreviewRequestFlowsLogin) GetBlock() optionalnullable.OptionalNullable[int64] {
	if p == nil {
		return nil
	}
	return p.Block
}

func (p *PolicyPreviewRequestFlowsLogin) GetInheritsAccountCreation() *bool {
	if p == nil {
		return nil
	}
	return p.InheritsAccountCreation
}
