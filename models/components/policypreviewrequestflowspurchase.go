package components

import (
	"github.com/noresponsehq/go/v10/optionalnullable"
)

type PolicyPreviewRequestFlowsPurchase struct {
	StepUp *int64                                   `json:"step_up,omitzero"`
	Review *int64                                   `json:"review,omitzero"`
	Block  optionalnullable.OptionalNullable[int64] `json:"block,omitzero"`
	// Send true; false is refused.
	InheritsAccountCreation *bool `json:"inherits_account_creation,omitzero"`
}

func (p *PolicyPreviewRequestFlowsPurchase) GetStepUp() *int64 {
	if p == nil {
		return nil
	}
	return p.StepUp
}

func (p *PolicyPreviewRequestFlowsPurchase) GetReview() *int64 {
	if p == nil {
		return nil
	}
	return p.Review
}

func (p *PolicyPreviewRequestFlowsPurchase) GetBlock() optionalnullable.OptionalNullable[int64] {
	if p == nil {
		return nil
	}
	return p.Block
}

func (p *PolicyPreviewRequestFlowsPurchase) GetInheritsAccountCreation() *bool {
	if p == nil {
		return nil
	}
	return p.InheritsAccountCreation
}
