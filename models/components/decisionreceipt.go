package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// DecisionReceipt - The decision prepared for one HTTP attempt. Its account field preserves the account decision returned for that attempt, even after subsequent changes.
type DecisionReceipt struct {
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the workspace; another workspace's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// Monotonic revision of this assessment. 1 on the synchronous answer; each material change while the assessment is served (an async continuation settling a source, a change of recommendation, confidence or status) increments it. Expiry does not increment it: an expired assessment is gone, not revised, and the expiry webhook carries the last served version.
	AssessmentVersion int64 `json:"assessment_version"`
	// Version of the hand-authored heuristic policy that produced the recommendation, for example 'risk-2026-10.7'. A workspace threshold override adds a +t<N> suffix, e.g. 'risk-2026-10.7+t3'.
	PolicyVersion *string `json:"policy_version,omitzero"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// Processing state, separate from evidence coverage. complete: final; nothing is still running, but unsupported, failed or unanswered lanes remain in incomplete_sources. partial: a recommendation exists and evidence is incomplete; assessment_url identifies continuing work when present. pending: no recommendation yet; poll or wait for the webhook. Finished processing does not mean all sources answered.
	Status Status `json:"status"`
	Score  *int64 `json:"score,omitzero"`
	// Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
	Account *Account `json:"account,omitzero"`
}

func (d DecisionReceipt) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DecisionReceipt) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (d *DecisionReceipt) GetAssessmentID() string {
	if d == nil {
		return ""
	}
	return d.AssessmentID
}

func (d *DecisionReceipt) GetAssessmentVersion() int64 {
	if d == nil {
		return 0
	}
	return d.AssessmentVersion
}

func (d *DecisionReceipt) GetPolicyVersion() *string {
	if d == nil {
		return nil
	}
	return d.PolicyVersion
}

func (d *DecisionReceipt) GetFlow() Flow {
	if d == nil {
		return Flow("")
	}
	return d.Flow
}

func (d *DecisionReceipt) GetStatus() Status {
	if d == nil {
		return Status("")
	}
	return d.Status
}

func (d *DecisionReceipt) GetScore() *int64 {
	if d == nil {
		return nil
	}
	return d.Score
}

func (d *DecisionReceipt) GetAccount() *Account {
	if d == nil {
		return nil
	}
	return d.Account
}
