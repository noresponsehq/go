package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AccountDecisionSnapshot - Retrieve the account decision at one recorded risk version. The answer uses the original assessment and account-specific risk inputs. Earlier versions that were not captured return 404. The assessment version, policy version and account calculation version identify the decision.
type AccountDecisionSnapshot struct {
	// Body of the 200 answer to GET /v1/accounts/{account_id}: the one current risk of one of the calling workspace's accounts. It is the risk of the account's latest assessment that has a decision (status complete or partial), by created_at and then assessment_id, scored exactly as GET /v1/assessments/{assessment_id} scores that assessment, plus the account's own reason shared_api_key when a key issued to the account was used from many IP addresses or countries (POST /v1/key_usage/batches). Its assessments are the account's history (GET /v1/accounts/{account_id}/assessments).
	Account Account `json:"account"`
	// Monotonic revision of this assessment. 1 on the synchronous answer; each material change while the assessment is served (an async continuation settling a source, a change of recommendation, confidence or status) increments it. Expiry does not increment it: an expired assessment is gone, not revised, and the expiry webhook carries the last served version.
	AssessmentVersion int64 `json:"assessment_version"`
	// Version of the hand-authored heuristic policy that produced the recommendation, for example 'risk-2026-10.7'. A workspace threshold override adds a +t<N> suffix, e.g. 'risk-2026-10.7+t3'.
	PolicyVersion string `json:"policy_version"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	overlayVersion int64  `const:"1" json:"overlay_version"`
	ChangedAt      string `json:"changed_at"`
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	FeedbackOutcomeID *string `json:"feedback_outcome_id,omitzero"`
}

func (a AccountDecisionSnapshot) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDecisionSnapshot) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDecisionSnapshot) GetAccount() Account {
	if a == nil {
		return Account{}
	}
	return a.Account
}

func (a *AccountDecisionSnapshot) GetAssessmentVersion() int64 {
	if a == nil {
		return 0
	}
	return a.AssessmentVersion
}

func (a *AccountDecisionSnapshot) GetPolicyVersion() string {
	if a == nil {
		return ""
	}
	return a.PolicyVersion
}

func (a *AccountDecisionSnapshot) GetOverlayVersion() int64 {
	return 1
}

func (a *AccountDecisionSnapshot) GetChangedAt() string {
	if a == nil {
		return ""
	}
	return a.ChangedAt
}

func (a *AccountDecisionSnapshot) GetFeedbackOutcomeID() *string {
	if a == nil {
		return nil
	}
	return a.FeedbackOutcomeID
}
