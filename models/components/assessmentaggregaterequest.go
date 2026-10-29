package components

import (
	"encoding/json"
	"fmt"
	"github.com/noresponsehq/go/v9/internal/utils"
)

type AssessmentAggregateRequestMetric string

const (
	AssessmentAggregateRequestMetricCount          AssessmentAggregateRequestMetric = "count"
	AssessmentAggregateRequestMetricScoreHistogram AssessmentAggregateRequestMetric = "score_histogram"
	AssessmentAggregateRequestMetricScoreOutcomes  AssessmentAggregateRequestMetric = "score_outcomes"
	AssessmentAggregateRequestMetricLateFlips      AssessmentAggregateRequestMetric = "late_flips"
)

func (e AssessmentAggregateRequestMetric) ToPointer() *AssessmentAggregateRequestMetric {
	return &e
}
func (e *AssessmentAggregateRequestMetric) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "count":
		fallthrough
	case "score_histogram":
		fallthrough
	case "score_outcomes":
		fallthrough
	case "late_flips":
		*e = AssessmentAggregateRequestMetric(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssessmentAggregateRequestMetric: %v", v)
	}
}

// AssessmentAggregateRequest - Count this environment's assessments created in [from,to) by bucket, from counts kept as each assessment is first answered, not by scanning assessments. Each assessment counts once, by its first answer (the answer served when it was first answered); later revisions and late results do not count it again, and late_flips counts its later account decision changes. from and to start a bucket (an hour, or a day or a week from Monday that starts at midnight in time_zone, UTC by default), cover at most 366 days (31 for hour buckets, 92 for day and week buckets in a time_zone other than UTC) and from is inside the retention period. The filters are POST /v1/assessments/search's and describe the first answer, so a cell drills down to search with the same filters, revision first and the cell's bucket as from and to. Counts hold no identifier and stay when an assessment or account is deleted or erased; they expire with retention and go with the workspace. Reasons are counted by code and impact band, not by weight. No Idempotency-Key. Select retained_first_answers for exact facet filters and pairs over retained detail. This mode supports count, score_histogram and score_outcomes, and refuses late_flips. It has a 31-day and 100000-record work budget.
type AssessmentAggregateRequest struct {
	// The first bucket's start.
	From string `json:"from"`
	// The end, excluded: the start of the bucket after the last.
	To string `json:"to"`
	// The period each count covers: an hour, or a day or a week (from Monday) in time_zone. Hour buckets are whole UTC hours in any time zone.
	Bucket Bucket `json:"bucket"`
	// The IANA time zone whose local midnight starts each day and week bucket, such as Europe/London or Asia/Kolkata; daylight saving changes make those days 23 or 25 hours long. A name the zone database does not know is refused (400 invalid_request naming time_zone). Day and week buckets in a zone other than UTC cover at most 92 days (400 invalid_request naming to); a year of them reads too many hourly counts to answer in time. Hour buckets stay whole UTC hours. Bucket starts, from and to stay instants. With the default population, a zone whose local midnight falls inside a UTC hour (a half-hour offset such as Asia/Kolkata) is refused for day and week buckets, as those counts are kept by the hour; retained_first_answers counts it.
	TimeZone *string `default:"UTC" json:"time_zone"`
	// At most two dimensions to count by, of which at most one is reason, lane, country or asn. An assessment counts once in each of its reasons and incomplete lanes, so those groups can add up to more than the bucket's count. With population retained_first_answers, both dimensions can be facets; counts cover retained first answers only.
	GroupBy []AggregateDimension `json:"group_by,omitzero"`
	// count: each bucket's count and groups. score_histogram: first answers by flow and 5-point score bucket. score_outcomes: the histogram's first answers that have a reported outcome, by flow, 5-point score bucket and the latest outcome reported for the assessment (POST /v1/assessments/{assessment_id}/outcomes); the rest of a bin has none reported. late_flips: changes of an assessment's account decision for its flow after its first answer, by from, to and delay.
	Metrics []AssessmentAggregateRequestMetric `json:"metrics,omitzero"`
	// The most groups answered: the largest over the whole range, then by value. groups_truncated says when there were more.
	Limit *int64 `default:"20" json:"limit"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow *Flow `json:"flow,omitzero"`
	// Processing state, separate from evidence coverage. complete: final; nothing is still running, but unsupported, failed or unanswered lanes remain in incomplete_sources. partial: a recommendation exists and evidence is incomplete; assessment_url identifies continuing work when present. pending: no recommendation yet; poll or wait for the webhook. Finished processing does not mean all sources answered.
	Status *Status `json:"status,omitzero"`
	// The recommended action: allow, step_up or review. Your application applies it.
	Recommendation *Recommendation `json:"recommendation,omitzero"`
	// An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.
	Decision *Decision `json:"decision,omitzero"`
	// Which answer of each assessment the filters describe: latest, the latest stored revision, or first, the first answer, as served when the assessment was first answered (POST /v1/assessments/aggregate counts first answers). A cell of an aggregate drills down to POST /v1/assessments/search with the same filters, revision first and the cell's bucket as from and to.
	Revision *AnswerRevision `json:"revision,omitzero"`
	// first_answers uses anonymous counters that can include erased detail. retained_first_answers computes exact intersections over retained first answers within 31 days. Requests exceeding the 100000-record work budget are refused.
	Population *Population `json:"population,omitzero"`
	Band       *Band       `json:"band,omitzero"`
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

func (a AssessmentAggregateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentAggregateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentAggregateRequest) GetFrom() string {
	if a == nil {
		return ""
	}
	return a.From
}

func (a *AssessmentAggregateRequest) GetTo() string {
	if a == nil {
		return ""
	}
	return a.To
}

func (a *AssessmentAggregateRequest) GetBucket() Bucket {
	if a == nil {
		return Bucket("")
	}
	return a.Bucket
}

func (a *AssessmentAggregateRequest) GetTimeZone() *string {
	if a == nil {
		return nil
	}
	return a.TimeZone
}

func (a *AssessmentAggregateRequest) GetGroupBy() []AggregateDimension {
	if a == nil {
		return nil
	}
	return a.GroupBy
}

func (a *AssessmentAggregateRequest) GetMetrics() []AssessmentAggregateRequestMetric {
	if a == nil {
		return nil
	}
	return a.Metrics
}

func (a *AssessmentAggregateRequest) GetLimit() *int64 {
	if a == nil {
		return nil
	}
	return a.Limit
}

func (a *AssessmentAggregateRequest) GetFlow() *Flow {
	if a == nil {
		return nil
	}
	return a.Flow
}

func (a *AssessmentAggregateRequest) GetStatus() *Status {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *AssessmentAggregateRequest) GetRecommendation() *Recommendation {
	if a == nil {
		return nil
	}
	return a.Recommendation
}

func (a *AssessmentAggregateRequest) GetDecision() *Decision {
	if a == nil {
		return nil
	}
	return a.Decision
}

func (a *AssessmentAggregateRequest) GetRevision() *AnswerRevision {
	if a == nil {
		return nil
	}
	return a.Revision
}

func (a *AssessmentAggregateRequest) GetPopulation() *Population {
	if a == nil {
		return nil
	}
	return a.Population
}

func (a *AssessmentAggregateRequest) GetBand() *Band {
	if a == nil {
		return nil
	}
	return a.Band
}

func (a *AssessmentAggregateRequest) GetReason() *ReasonCode {
	if a == nil {
		return nil
	}
	return a.Reason
}

func (a *AssessmentAggregateRequest) GetImpact() *Impact {
	if a == nil {
		return nil
	}
	return a.Impact
}

func (a *AssessmentAggregateRequest) GetLane() *Lane {
	if a == nil {
		return nil
	}
	return a.Lane
}

func (a *AssessmentAggregateRequest) GetCountry() *string {
	if a == nil {
		return nil
	}
	return a.Country
}

func (a *AssessmentAggregateRequest) GetAsn() *int64 {
	if a == nil {
		return nil
	}
	return a.Asn
}

func (a *AssessmentAggregateRequest) GetScoreMin() *int64 {
	if a == nil {
		return nil
	}
	return a.ScoreMin
}

func (a *AssessmentAggregateRequest) GetScoreMax() *int64 {
	if a == nil {
		return nil
	}
	return a.ScoreMax
}

func (a *AssessmentAggregateRequest) GetListMatch() *ListMatch {
	if a == nil {
		return nil
	}
	return a.ListMatch
}
