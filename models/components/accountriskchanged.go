package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AccountRiskChanged - The account's decision for at least one flow changed. account is the account as GET /v1/accounts/{account_id} answers it for the revision that made the change: its new score, decision, reasons and decisions_by_flow, and the assessment the risk comes from. previous_decisions_by_flow are the decisions it replaced, so account.decisions_by_flow differs from them for at least one flow, and previous_risk_version is the risk_version they were announced at. On the account's first event no decision was announced before: previous_risk_version is absent and previous_decisions_by_flow is allow for every flow, the decisions an account has before its first event. Events of one account are not ordered: account.risk_version orders them, so keep the account with the highest risk_version and ignore an event whose account is older than the one held.
type AccountRiskChanged struct {
	// Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
	Account Account `json:"account"`
	// The decision for each flow: the account's score under that flow's thresholds. A flow the workspace has not set uses its account_creation thresholds, so the flows differ only where it set them. The flow of the current assessment is held while that flow's admission is held and the assessment's investigation has not completed.
	PreviousDecisionsByFlow AccountDecisionsByFlow `json:"previous_decisions_by_flow"`
	// The account's risk_version when previous_decisions_by_flow were announced, always lower than account.risk_version. Absent on the account's first event, when no decision was announced before, and on the first event after the account was deleted.
	PreviousRiskVersion *int64 `json:"previous_risk_version,omitzero"`
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

func (a *AccountRiskChanged) GetPreviousRiskVersion() *int64 {
	if a == nil {
		return nil
	}
	return a.PreviousRiskVersion
}
