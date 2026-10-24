package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

// AccountRiskChanged - The account's decision for at least one flow changed. account is the account as GET /v1/accounts/{account_id} answers it for the revision that made the change: its new score, decision, reasons and decisions_by_flow, and the assessment the risk comes from. previous_decisions_by_flow are the decisions it replaced, so account.decisions_by_flow differs from them for at least one flow. On the account's first decision, previous_decisions_by_flow is allow for every flow. Events of one account are not ordered: account.risk_version orders them, so keep the account with the highest risk_version and ignore an event whose account is older than the one held.
type AccountRiskChanged struct {
	// Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
	Account Account `json:"account"`
	// The decision for each flow: the account's score under that flow's thresholds. A flow the workspace has not set uses its account_creation thresholds, so the flows differ only where it set them.
	PreviousDecisionsByFlow AccountDecisionsByFlow `json:"previous_decisions_by_flow"`
}

func (a AccountRiskChanged) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountRiskChanged) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountRiskChanged) GetAccount() Account {
	if a == nil {
		return Account{}
	}
	return a.Account
}

func (a *AccountRiskChanged) GetPreviousDecisionsByFlow() AccountDecisionsByFlow {
	if a == nil {
		return AccountDecisionsByFlow{}
	}
	return a.PreviousDecisionsByFlow
}
