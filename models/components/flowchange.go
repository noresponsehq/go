package components

import (
	"github.com/noresponsehq/go/v9/optionalnullable"
)

// FlowChange - A change to one flow. A field left out keeps the value in force. Thresholds set on login or purchase start from the values in force, account_creation's while it inherits them. After the change, 0 < step_up < review <= 100 and block is off or review <= block <= 100.
type FlowChange struct {
	StepUp *int64 `json:"step_up,omitzero"`
	Review *int64 `json:"review,omitzero"`
	// The score from which the decision is block; null turns blocking off.
	Block optionalnullable.OptionalNullable[int64] `json:"block,omitzero"`
	// assessments only.
	DefaultDepth *UsageCountDepth `json:"default_depth,omitzero"`
	// Login or purchase only: drop the flow's own thresholds so it uses account_creation's again. Not combined with step_up, review or block. Send true; false is refused.
	InheritsAccountCreation *bool `json:"inherits_account_creation,omitzero"`
	// Alias of inherits_account_creation, still accepted. Send true; false is refused.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	InheritsSignup *bool `json:"inherits_signup,omitzero"`
	// immediate (the default): an assessment of this flow is decided from the evidence its answer has, and later evidence may change the decision. held: while the account's current assessment of this flow has not completed, the account's decision for the flow is held, and the final decision arrives as account.risk_changed when the investigation completes or hold_seconds pass, whichever is first. Each flow has its own admission.
	Admission *Admission `json:"admission,omitzero"`
	// How long a held flow's assessment is held after it was created, 30 to 86400 seconds (default 300). Kept while the flow is immediate.
	HoldSeconds *int64 `json:"hold_seconds,omitzero"`
	// The decision a held assessment takes when hold_seconds pass before its investigation completes, unless the evidence so far gives a stricter one (default review). It is announced as account.risk_changed like any other change, so an expired hold is always announced. A hosted provider hook, which cannot wait, answers a held account by it at once: deny for block, allow otherwise. Kept while the flow is immediate.
	OnExpiry *OnExpiry `json:"on_expiry,omitzero"`
}

func (f *FlowChange) GetStepUp() *int64 {
	if f == nil {
		return nil
	}
	return f.StepUp
}

func (f *FlowChange) GetReview() *int64 {
	if f == nil {
		return nil
	}
	return f.Review
}

func (f *FlowChange) GetBlock() optionalnullable.OptionalNullable[int64] {
	if f == nil {
		return nil
	}
	return f.Block
}

func (f *FlowChange) GetDefaultDepth() *UsageCountDepth {
	if f == nil {
		return nil
	}
	return f.DefaultDepth
}

func (f *FlowChange) GetInheritsAccountCreation() *bool {
	if f == nil {
		return nil
	}
	return f.InheritsAccountCreation
}

func (f *FlowChange) GetInheritsSignup() *bool {
	if f == nil {
		return nil
	}
	return f.InheritsSignup
}

func (f *FlowChange) GetAdmission() *Admission {
	if f == nil {
		return nil
	}
	return f.Admission
}

func (f *FlowChange) GetHoldSeconds() *int64 {
	if f == nil {
		return nil
	}
	return f.HoldSeconds
}

func (f *FlowChange) GetOnExpiry() *OnExpiry {
	if f == nil {
		return nil
	}
	return f.OnExpiry
}
