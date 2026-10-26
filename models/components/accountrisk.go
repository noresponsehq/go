package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

type AccountRisk struct {
	// 0 (safe) to 100 (risky), as risk.score of the assessment.
	Score int64 `json:"score"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.
	Decision AccountDecision `json:"decision"`
	// Every reason that fired: risk-raising reasons first, then neutral, then risk-lowering, each strongest impact band first and then by code. The policy's numeric weights are kept private, so a reason carries only its direction and impact band.
	Reasons []AccountRiskReason `json:"reasons"`
}

func (a AccountRisk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountRisk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountRisk) GetScore() int64 {
	if a == nil {
		return 0
	}
	return a.Score
}

func (a *AccountRisk) GetDecision() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.Decision
}

func (a *AccountRisk) GetReasons() []AccountRiskReason {
	if a == nil {
		return []AccountRiskReason{}
	}
	return a.Reasons
}
