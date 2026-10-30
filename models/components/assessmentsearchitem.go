package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentSearchItem - One stored assessment revision, exactly as GET /v1/assessments/{assessment_id}/versions/{version} returns it, with the decision of the assessment's first answer. It is not re-scored and carries no current account risk.
type AssessmentSearchItem struct {
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
	// An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.
	FirstDecision *Decision `json:"first_decision,omitzero"`
	// Why this result matched a typed identifier search. value is the original value submitted in assessment_id. Provenance records a customer submission, not a source discovery or verified ownership. Equality is deterministic and has no inferred confidence score. An account result names its newest retained matching assessment, which can differ from latest.
	MatchedIdentifier *IdentifierMatch `json:"matched_identifier,omitzero"`
	// Present only for identifier_expression. All satisfied unique terms, ordered by term_index, with original submitted values and retained assessment witnesses. A duplicate normalized term appears once at its first raw leaf index. The legacy single identifier filter instead returns matched_identifier.
	MatchedIdentifiers []IdentifierTermMatch `json:"matched_identifiers,omitzero"`
}

func (a AssessmentSearchItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentSearchItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentSearchItem) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentSearchItem) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AssessmentSearchItem) GetStatus() Status {
	if a == nil {
		return Status("")
	}
	return a.Status
}

func (a *AssessmentSearchItem) GetFlow() Flow {
	if a == nil {
		return Flow("")
	}
	return a.Flow
}

func (a *AssessmentSearchItem) GetDepth() Depth {
	if a == nil {
		return Depth("")
	}
	return a.Depth
}

func (a *AssessmentSearchItem) GetDetail() AssessmentDetail {
	if a == nil {
		return AssessmentDetail{}
	}
	return a.Detail
}

func (a *AssessmentSearchItem) GetFirstDecision() *Decision {
	if a == nil {
		return nil
	}
	return a.FirstDecision
}

func (a *AssessmentSearchItem) GetMatchedIdentifier() *IdentifierMatch {
	if a == nil {
		return nil
	}
	return a.MatchedIdentifier
}

func (a *AssessmentSearchItem) GetMatchedIdentifiers() []IdentifierTermMatch {
	if a == nil {
		return nil
	}
	return a.MatchedIdentifiers
}
