package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

type AssessmentAggregateMetric string

const (
	AssessmentAggregateMetricCount          AssessmentAggregateMetric = "count"
	AssessmentAggregateMetricScoreHistogram AssessmentAggregateMetric = "score_histogram"
	AssessmentAggregateMetricScoreOutcomes  AssessmentAggregateMetric = "score_outcomes"
	AssessmentAggregateMetricLateFlips      AssessmentAggregateMetric = "late_flips"
)

func (e AssessmentAggregateMetric) ToPointer() *AssessmentAggregateMetric {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AssessmentAggregateMetric) IsExact() bool {
	if e != nil {
		switch *e {
		case "count", "score_histogram", "score_outcomes", "late_flips":
			return true
		}
	}
	return false
}

// AssessmentAggregate - Body of the 200 answer to POST /v1/assessments/aggregate: one entry per bucket of [from,to), empty buckets included. Counts are of first answers and stay when an assessment or account is deleted or erased; search returns only assessments still retained.
type AssessmentAggregate struct {
	// Version of the body shape.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	schemaVersion string `const:"1" json:"schema_version"`
	From          string `json:"from"`
	To            string `json:"to"`
	// The period each count covers, in UTC; a week starts on Monday.
	Bucket  Bucket                      `json:"bucket"`
	GroupBy []AggregateDimension        `json:"group_by"`
	Metrics []AssessmentAggregateMetric `json:"metrics"`
	// The calling workspace's configured retention period in days. Default 365 (12 months), which is also the maximum: a workspace may only shorten it, to at least 1 day. Each record expires from its own created_at; see assessment expires_at.
	RetentionDays *int64 `default:"365" json:"retention_days"`
	// true when more groups had counts than limit allowed.
	GroupsTruncated bool              `json:"groups_truncated"`
	Buckets         []AggregateBucket `json:"buckets"`
}

func (a AssessmentAggregate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentAggregate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentAggregate) GetSchemaVersion() string {
	return "1"
}

func (a *AssessmentAggregate) GetFrom() string {
	if a == nil {
		return ""
	}
	return a.From
}

func (a *AssessmentAggregate) GetTo() string {
	if a == nil {
		return ""
	}
	return a.To
}

func (a *AssessmentAggregate) GetBucket() Bucket {
	if a == nil {
		return Bucket("")
	}
	return a.Bucket
}

func (a *AssessmentAggregate) GetGroupBy() []AggregateDimension {
	if a == nil {
		return []AggregateDimension{}
	}
	return a.GroupBy
}

func (a *AssessmentAggregate) GetMetrics() []AssessmentAggregateMetric {
	if a == nil {
		return []AssessmentAggregateMetric{}
	}
	return a.Metrics
}

func (a *AssessmentAggregate) GetRetentionDays() *int64 {
	if a == nil {
		return nil
	}
	return a.RetentionDays
}

func (a *AssessmentAggregate) GetGroupsTruncated() bool {
	if a == nil {
		return false
	}
	return a.GroupsTruncated
}

func (a *AssessmentAggregate) GetBuckets() []AggregateBucket {
	if a == nil {
		return []AggregateBucket{}
	}
	return a.Buckets
}
