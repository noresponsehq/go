package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type AggregateBucket struct {
	// The bucket's first instant, in UTC. A day or week bucket in a time zone starts at local midnight there, so in Asia/Tokyo a day starts at 15:00Z the day before.
	Start string `json:"start"`
	// Assessments first answered whose creation falls in the bucket. With metric count.
	Count *int64 `json:"count,omitzero"`
	// With metric count and group_by: the bucket's counts of the groups answered, largest group first.
	Groups         []AggregateGroup      `json:"groups,omitzero"`
	ScoreHistogram []AggregateScoreCount `json:"score_histogram,omitzero"`
	// With metric score_outcomes: the score_histogram's first answers whose assessment has a reported outcome, by its latest. A bin's count less its score_outcomes is the first answers with none reported.
	ScoreOutcomes []AggregateScoreOutcome `json:"score_outcomes,omitzero"`
	LateFlips     []AggregateLateFlip     `json:"late_flips,omitzero"`
}

func (a AggregateBucket) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AggregateBucket) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AggregateBucket) GetStart() string {
	if a == nil {
		return ""
	}
	return a.Start
}

func (a *AggregateBucket) GetCount() *int64 {
	if a == nil {
		return nil
	}
	return a.Count
}

func (a *AggregateBucket) GetGroups() []AggregateGroup {
	if a == nil {
		return nil
	}
	return a.Groups
}

func (a *AggregateBucket) GetScoreHistogram() []AggregateScoreCount {
	if a == nil {
		return nil
	}
	return a.ScoreHistogram
}

func (a *AggregateBucket) GetScoreOutcomes() []AggregateScoreOutcome {
	if a == nil {
		return nil
	}
	return a.ScoreOutcomes
}

func (a *AggregateBucket) GetLateFlips() []AggregateLateFlip {
	if a == nil {
		return nil
	}
	return a.LateFlips
}
