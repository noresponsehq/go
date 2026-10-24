package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

// AccountDecisionsByFlow - The decision for each flow: the account's score under that flow's thresholds. A flow the workspace has not set uses its account_creation thresholds, so the flows differ only where it set them. The flow of the current assessment is held while that flow's admission is held and the assessment's investigation has not completed.
type AccountDecisionsByFlow struct {
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.
	AccountCreation AccountDecision `json:"account_creation"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.
	Login AccountDecision `json:"login"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.
	Purchase AccountDecision `json:"purchase"`
}

func (a AccountDecisionsByFlow) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDecisionsByFlow) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDecisionsByFlow) GetAccountCreation() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.AccountCreation
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
