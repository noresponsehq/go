package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentRevision - Retrieve a retained assessment version with its original decision, policy, flow and depth. This response excludes the current account decision. Current retention and response-mode permissions still apply.
type AssessmentRevision struct {
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
	// One assessment's own facts: its score and recommendation, the signals behind them, the lanes still running and its timing. risk here is this assessment's own score; the account's current risk is the account member of the assessment. risk and recommendation are present for complete and partial; assessment_url is present for pending.
	Detail AssessmentDetail `json:"detail"`
}

func (a AssessmentRevision) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentRevision) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentRevision) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentRevision) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AssessmentRevision) GetStatus() Status {
	if a == nil {
		return Status("")
	}
	return a.Status
}

func (a *AssessmentRevision) GetFlow() Flow {
	if a == nil {
		return Flow("")
	}
	return a.Flow
}

func (a *AssessmentRevision) GetDepth() Depth {
	if a == nil {
		return Depth("")
	}
	return a.Depth
}

func (a *AssessmentRevision) GetDetail() AssessmentDetail {
	if a == nil {
		return AssessmentDetail{}
	}
	return a.Detail
}
