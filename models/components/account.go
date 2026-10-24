package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

// Account - Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
type Account struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// The account_id the workspace sends with its assessments.
	ID   string      `json:"id"`
	Risk AccountRisk `json:"risk"`
	// The decision for each flow: the account's score under that flow's thresholds. A flow the workspace has not set uses its account_creation thresholds, so the flows differ only where it set them.
	DecisionsByFlow AccountDecisionsByFlow `json:"decisions_by_flow"`
	// The assessment the current risk comes from.
	AssessmentID string `json:"assessment_id"`
	// updated_at of that assessment's latest revision.
	UpdatedAt string `json:"updated_at"`
	// The version of the account's current risk. It increases every time the account's current risk or decisions_by_flow changes, and does not decrease. Of two copies of an account, the one with the higher risk_version is newer. account.risk_changed carries the risk_version that its change made.
	RiskVersion int64 `json:"risk_version"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Account) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *Account) GetSchemaVersion() string {
	return "1"
}

func (a *Account) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *Account) GetRisk() AccountRisk {
	if a == nil {
		return AccountRisk{}
	}
	return a.Risk
}

func (a *Account) GetDecisionsByFlow() AccountDecisionsByFlow {
	if a == nil {
		return AccountDecisionsByFlow{}
	}
	return a.DecisionsByFlow
}

func (a *Account) GetAssessmentID() string {
	if a == nil {
		return ""
	}
	return a.AssessmentID
}

func (a *Account) GetUpdatedAt() string {
	if a == nil {
		return ""
	}
	return a.UpdatedAt
}

func (a *Account) GetRiskVersion() int64 {
	if a == nil {
		return 0
	}
	return a.RiskVersion
}
