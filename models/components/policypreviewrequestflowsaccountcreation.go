package components

import (
	"github.com/noresponsehq/go/v9/optionalnullable"
)

type PolicyPreviewRequestFlowsAccountCreation struct {
	StepUp *int64                                   `json:"step_up,omitzero"`
	Review *int64                                   `json:"review,omitzero"`
	Block  optionalnullable.OptionalNullable[int64] `json:"block,omitzero"`
}

func (p *PolicyPreviewRequestFlowsAccountCreation) GetStepUp() *int64 {
	if p == nil {
		return nil
	}
	return p.StepUp
}

func (p *PolicyPreviewRequestFlowsAccountCreation) GetReview() *int64 {
	if p == nil {
		return nil
	}
	return p.Review
}

func (p *PolicyPreviewRequestFlowsAccountCreation) GetBlock() optionalnullable.OptionalNullable[int64] {
	if p == nil {
		return nil
	}
	return p.Block
}
