package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

type AccountRisk struct {
	// 0 (safe) to 100 (risky), as risk.score of the assessment.
	Score int64 `json:"score"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. Your application applies it.
	Decision AccountDecision `json:"decision"`
	// Every reason that fired, highest weight first. The weights sum to score before clamping.
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
