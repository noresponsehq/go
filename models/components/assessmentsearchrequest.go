package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentSearchRequest - Search retained assessments created in [from,to), over at most 366 days. Filters use each assessment decision, rather than the current account decision. Results can change between pages. Keep the assessment version to retrieve the exact decision later. This read does not require an idempotency key.
type AssessmentSearchRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID *string `json:"account_id,omitzero"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow *Flow `json:"flow,omitzero"`
	// complete: final; nothing is still running. A lane that failed or missed its 10 s deadline stays listed in incomplete_sources, and confidence is lower. partial: a recommendation exists and at least one lane is still running or did not settle (see incomplete_sources). pending: nothing usable before the deadline; no recommendation yet, poll or wait for the webhook.
	Status *Status `json:"status,omitzero"`
	// The recommended action: allow, step_up or review. Your application applies it.
	Recommendation *Recommendation `json:"recommendation,omitzero"`
	Limit          *int64          `default:"20" json:"limit"`
	// Opaque, expiring cursor bound to this environment, operation and every filter. Keep the filters unchanged.
	Cursor *string `json:"cursor,omitzero"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode *ResponseMode `json:"response_mode,omitzero"`
	// An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.
	Decision *Decision `json:"decision,omitzero"`
	// Which answer of each assessment the filters describe: latest, the latest stored revision, or first, the first answer, as served when the assessment was first answered (POST /v1/assessments/aggregate counts first answers). A cell of an aggregate drills down to POST /v1/assessments/search with the same filters, revision first and the cell's bucket as from and to.
	Revision *AnswerRevision `json:"revision,omitzero"`
	Band     *Band           `json:"band,omitzero"`
	// A stable reason code. New codes may appear; ignore any you do not know.
	Reason *ReasonCode `json:"reason,omitzero"`
	// The selected reason's served impact band. Requires reason.
	Impact *Impact `json:"impact,omitzero"`
	// An evidence lane: accounts (account presence for email and username), phone, breach, domain (the email domain's registration) and ip.
	Lane    *Lane   `json:"lane,omitzero"`
	Country *string `json:"country,omitzero"`
	Asn     *int64  `json:"asn,omitzero"`
	// Inclusive lower score bound. Pending answers without a score do not match.
	ScoreMin *int64 `json:"score_min,omitzero"`
	// Inclusive upper score bound, at least score_min when both are supplied.
	ScoreMax *int64 `json:"score_max,omitzero"`
	// The list action captured at the first answer. Requires revision first. Older answers without captured list state do not match, including none.
	ListMatch *ListMatch `json:"list_match,omitzero"`
}

func (a AssessmentSearchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentSearchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentSearchRequest) GetFrom() string {
	if a == nil {
		return ""
	}
	return a.From
}

func (a *AssessmentSearchRequest) GetTo() string {
	if a == nil {
		return ""
	}
	return a.To
}

func (a *AssessmentSearchRequest) GetAccountID() *string {
	if a == nil {
		return nil
	}
	return a.AccountID
}

func (a *AssessmentSearchRequest) GetFlow() *Flow {
	if a == nil {
		return nil
	}
	return a.Flow
}

func (a *AssessmentSearchRequest) GetStatus() *Status {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *AssessmentSearchRequest) GetRecommendation() *Recommendation {
	if a == nil {
		return nil
	}
	return a.Recommendation
}

func (a *AssessmentSearchRequest) GetLimit() *int64 {
	if a == nil {
		return nil
	}
	return a.Limit
}

func (a *AssessmentSearchRequest) GetCursor() *string {
	if a == nil {
		return nil
	}
	return a.Cursor
}

func (a *AssessmentSearchRequest) GetResponseMode() *ResponseMode {
	if a == nil {
		return nil
	}
	return a.ResponseMode
}

func (a *AssessmentSearchRequest) GetDecision() *Decision {
	if a == nil {
		return nil
	}
	return a.Decision
}

func (a *AssessmentSearchRequest) GetRevision() *AnswerRevision {
	if a == nil {
		return nil
	}
	return a.Revision
}

func (a *AssessmentSearchRequest) GetBand() *Band {
	if a == nil {
		return nil
	}
	return a.Band
}

func (a *AssessmentSearchRequest) GetReason() *ReasonCode {
	if a == nil {
		return nil
	}
	return a.Reason
}

func (a *AssessmentSearchRequest) GetImpact() *Impact {
	if a == nil {
		return nil
	}
	return a.Impact
}

func (a *AssessmentSearchRequest) GetLane() *Lane {
	if a == nil {
		return nil
	}
	return a.Lane
}

func (a *AssessmentSearchRequest) GetCountry() *string {
	if a == nil {
		return nil
	}
	return a.Country
}

func (a *AssessmentSearchRequest) GetAsn() *int64 {
	if a == nil {
		return nil
	}
	return a.Asn
}

func (a *AssessmentSearchRequest) GetScoreMin() *int64 {
	if a == nil {
		return nil
	}
	return a.ScoreMin
}

func (a *AssessmentSearchRequest) GetScoreMax() *int64 {
	if a == nil {
		return nil
	}
	return a.ScoreMax
}

func (a *AssessmentSearchRequest) GetListMatch() *ListMatch {
	if a == nil {
		return nil
	}
	return a.ListMatch
}
