package components

import (
	"encoding/json"
	"fmt"
	"github.com/noresponsehq/go/v8/internal/utils"
)

type AssessmentAggregateRequestMetric string

const (
	AssessmentAggregateRequestMetricCount          AssessmentAggregateRequestMetric = "count"
	AssessmentAggregateRequestMetricScoreHistogram AssessmentAggregateRequestMetric = "score_histogram"
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
	case "late_flips":
		*e = AssessmentAggregateRequestMetric(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssessmentAggregateRequestMetric: %v", v)
	}
}

// AssessmentAggregateRequest - Count this environment's assessments created in [from,to) by bucket, from counts kept as each assessment is first answered, not by scanning assessments. Each assessment counts once, by its first answer (the answer served when it was first answered); later revisions and late results do not count it again, and late_flips counts its later account decision changes. from and to start a bucket (an hour, a UTC day, or a week from Monday 00:00 UTC), cover at most 366 days (31 for hour buckets) and from is inside the retention period. The filters are POST /v1/assessments/search's and describe the first answer, so a cell drills down to search with the same filters, revision first and the cell's bucket as from and to. Counts hold no identifier and stay when an assessment or account is deleted or erased; they expire with retention and go with the workspace. Reasons are counted by code and impact band, not by weight. No Idempotency-Key.
type AssessmentAggregateRequest struct {
	// The first bucket's start.
	From string `json:"from"`
	// The end, excluded: the start of the bucket after the last.
	To string `json:"to"`
	// The period each count covers, in UTC; a week starts on Monday.
	Bucket Bucket `json:"bucket"`
	// At most two dimensions to count by, of which at most one is reason, lane, country or asn. An assessment counts once in each of its reasons and incomplete lanes, so those groups can add up to more than the bucket's count.
	GroupBy []AggregateDimension `json:"group_by,omitzero"`
	// count: each bucket's count and groups. score_histogram: first answers by flow and 5-point score bucket. late_flips: changes of an assessment's account decision for its flow after its first answer, by from, to and delay.
	Metrics []AssessmentAggregateRequestMetric `json:"metrics,omitzero"`
	// The most groups answered: the largest over the whole range, then by value. groups_truncated says when there were more.
	Limit *int64 `default:"20" json:"limit"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow *Flow `json:"flow,omitzero"`
	// complete: final; nothing is still running. A lane that failed or missed its 10 s deadline stays listed in incomplete_sources, and confidence is lower. partial: a recommendation exists and at least one lane is still running or did not settle (see incomplete_sources). pending: nothing usable before the deadline; no recommendation yet, poll or wait for the webhook.
	Status *Status `json:"status,omitzero"`
	// The recommended action: allow, step_up or review. Your application applies it.
	Recommendation *Recommendation `json:"recommendation,omitzero"`
	// An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.
	Decision *Decision `json:"decision,omitzero"`
	// Which answer of each assessment the filters describe: latest, the latest stored revision, or first, the first answer, as served when the assessment was first answered (POST /v1/assessments/aggregate counts first answers). A cell of an aggregate drills down to POST /v1/assessments/search with the same filters, revision first and the cell's bucket as from and to.
	Revision *AnswerRevision `json:"revision,omitzero"`
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
