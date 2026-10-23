package components

// AccountDecisionsByFlow - The decision for each flow: the account's score under that flow's thresholds. A flow the workspace has not set uses its signup thresholds, so the flows differ only where it set them.
type AccountDecisionsByFlow struct {
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. Your application applies it.
	Signup AccountDecision `json:"signup"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. Your application applies it.
	Login AccountDecision `json:"login"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. Your application applies it.
	Purchase AccountDecision `json:"purchase"`
}

func (a *AccountDecisionsByFlow) GetSignup() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.Signup
}

func (a *AccountDecisionsByFlow) GetLogin() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.Login
}

func (a *AccountDecisionsByFlow) GetPurchase() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.Purchase
}
