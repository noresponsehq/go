package components

// FlowSettings - One flow's policy in force. A score below step_up is allow, below review step_up, and from review review; from block, when present, the decision is block. admission says when the flow's decision is final.
type FlowSettings struct {
	StepUp int64 `json:"step_up"`
	Review int64 `json:"review"`
	// The score from which the decision is block. Absent while blocking is off, the default.
	Block *int64 `json:"block,omitzero"`
	// The depth a request for this flow that names none is served at. Defaults: full for account_creation and purchase, quick for login.
	DefaultDepth DefaultDepth `json:"default_depth"`
	// true while this flow has no thresholds of its own and uses account_creation's, the default for login and purchase. Always false for signup.
	InheritsAccountCreation bool `json:"inherits_account_creation"`
	// immediate (the default): an assessment of this flow is decided from the evidence its answer has, and later evidence may change the decision. held: while the account's current assessment of this flow has not completed, the account's decision for the flow is held, and the final decision arrives as account.risk_changed when the investigation completes or hold_seconds pass, whichever is first. Each flow has its own admission.
	Admission Admission `json:"admission"`
	// How long a held flow's assessment is held after it was created, 30 to 86400 seconds (default 300). Kept while the flow is immediate.
	HoldSeconds int64 `json:"hold_seconds"`
	// The decision a held assessment takes when hold_seconds pass before its investigation completes, unless the evidence so far gives a stricter one (default review). It is announced as account.risk_changed like any other change, so an expired hold is always announced. A hosted provider hook, which cannot wait, answers a held account by it at once: deny for block, allow otherwise. Kept while the flow is immediate.
	OnExpiry OnExpiry `json:"on_expiry"`
}

func (f *FlowSettings) GetStepUp() int64 {
	if f == nil {
		return 0
	}
	return f.StepUp
}

func (f *FlowSettings) GetReview() int64 {
	if f == nil {
		return 0
	}
	return f.Review
}

func (f *FlowSettings) GetBlock() *int64 {
	if f == nil {
		return nil
	}
	return f.Block
}

func (f *FlowSettings) GetDefaultDepth() DefaultDepth {
	if f == nil {
		return DefaultDepth("")
	}
	return f.DefaultDepth
}

func (f *FlowSettings) GetInheritsAccountCreation() bool {
	if f == nil {
		return false
	}
	return f.InheritsAccountCreation
}

func (f *FlowSettings) GetAdmission() Admission {
	if f == nil {
		return Admission("")
	}
	return f.Admission
}

func (f *FlowSettings) GetHoldSeconds() int64 {
	if f == nil {
		return 0
	}
	return f.HoldSeconds
}

func (f *FlowSettings) GetOnExpiry() OnExpiry {
	if f == nil {
		return OnExpiry("")
	}
	return f.OnExpiry
}
