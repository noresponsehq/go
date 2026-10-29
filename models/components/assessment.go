package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// Assessment - One assessment, as POST /v1/assessments, GET /v1/assessments/{assessment_id} and the assessment.updated webhook return it. account is the account's one current risk: its score, the decision for this assessment's flow, the reasons and the decision for each flow. detail holds this assessment's own facts. status says whether the assessment is final (complete), usable with a source still running (partial) or not yet usable (pending). While the account's first assessment is pending, account is absent; wait for the webhook or poll the assessment. Keep the highest detail.assessment_version.
type Assessment struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	// This assessment's id, equal to detail.assessment_id.
	ID string `json:"id"`
	// Processing state, separate from evidence coverage. complete: final; nothing is still running, but unsupported, failed or unanswered lanes remain in incomplete_sources. partial: a recommendation exists and evidence is incomplete; assessment_url identifies continuing work when present. pending: no recommendation yet; poll or wait for the webhook. Finished processing does not mean all sources answered.
	Status Status `json:"status"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// How long the synchronous answer waits for sources. quick waits only for the fast sources, those that answer within 500 ms of ingress (or deadline_ms when lower); full waits for the whole budget, up to 3000 ms. Sources still running at the cut continue asynchronously and update the assessment either way.
	Depth Depth `json:"depth"`
	// Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
	Account *Account `json:"account,omitzero"`
	// One assessment's own facts: its score and recommendation, the signals behind them, the lanes still running and its timing. risk here is this assessment's own score; the account's current risk is the account member of the assessment. risk and recommendation are present for complete and partial; assessment_url is present for pending.
	Detail AssessmentDetail `json:"detail"`
}

func (a Assessment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Assessment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *Assessment) GetSchemaVersion() string {
	return "1"
}

func (a *Assessment) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *Assessment) GetStatus() Status {
	if a == nil {
		return Status("")
	}
	return a.Status
}

func (a *Assessment) GetFlow() Flow {
	if a == nil {
		return Flow("")
	}
	return a.Flow
}

func (a *Assessment) GetDepth() Depth {
	if a == nil {
		return Depth("")
	}
	return a.Depth
}

func (a *Assessment) GetAccount() *Account {
	if a == nil {
		return nil
	}
	return a.Account
}

func (a *Assessment) GetDetail() AssessmentDetail {
	if a == nil {
		return AssessmentDetail{}
	}
	return a.Detail
}
