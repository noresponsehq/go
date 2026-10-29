package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentAggregateCoverage - Captured in the same database snapshot as the counts. Capture start times describe when counting began, not the earliest creation time counted: an older assessment may first answer later. Older outcomes are not inferred to be absent. Erased detail cannot be returned by drill-down even while its anonymous count remains.
type AssessmentAggregateCoverage struct {
	AsOf                   string `json:"as_of"`
	FirstAnswersStartedAt  string `json:"first_answers_started_at"`
	ScoreOutcomesStartedAt string `json:"score_outcomes_started_at"`
	RetainedFrom           string `json:"retained_from"`
	// first_answers uses anonymous counters that can include erased detail. retained_first_answers computes exact intersections over retained first answers within 31 days. Requests exceeding the 100000-record work budget are refused.
	Population            Population `json:"population"`
	ErasedRecordsIncluded bool       `json:"erased_records_included"`
	OverlappingGroups     bool       `json:"overlapping_groups"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	lateFlipUnit string `const:"decision_transitions" json:"late_flip_unit"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	historicalBackfill bool `const:"false" json:"historical_backfill"`
}

func (a AssessmentAggregateCoverage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentAggregateCoverage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentAggregateCoverage) GetAsOf() string {
	if a == nil {
		return ""
	}
	return a.AsOf
}

func (a *AssessmentAggregateCoverage) GetFirstAnswersStartedAt() string {
	if a == nil {
		return ""
	}
	return a.FirstAnswersStartedAt
}

func (a *AssessmentAggregateCoverage) GetScoreOutcomesStartedAt() string {
	if a == nil {
		return ""
	}
	return a.ScoreOutcomesStartedAt
}

func (a *AssessmentAggregateCoverage) GetRetainedFrom() string {
	if a == nil {
		return ""
	}
	return a.RetainedFrom
}

func (a *AssessmentAggregateCoverage) GetPopulation() Population {
	if a == nil {
		return Population("")
	}
	return a.Population
}

func (a *AssessmentAggregateCoverage) GetErasedRecordsIncluded() bool {
	if a == nil {
		return false
	}
	return a.ErasedRecordsIncluded
}

func (a *AssessmentAggregateCoverage) GetOverlappingGroups() bool {
	if a == nil {
		return false
	}
	return a.OverlappingGroups
}

func (a *AssessmentAggregateCoverage) GetLateFlipUnit() string {
	return "decision_transitions"
}

func (a *AssessmentAggregateCoverage) GetHistoricalBackfill() bool {
	return false
}
