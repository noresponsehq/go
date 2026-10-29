package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// Reviews created in [from,to) whose review and current account assessment remain retained. Status and assignment are current at as_of, not a transition history. label_coverage numerator is resolved reviews labeled benign or abuse; denominator is resolved reviews. Inconclusive is not a known outcome. Review conclusions do not change account feedback or enforcement.
type Reviews struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	population  string `const:"retained_reviews_by_created_at" json:"population"`
	Denominator int64  `json:"denominator"`
	Open        int64  `json:"open"`
	InProgress  int64  `json:"in_progress"`
	Resolved    int64  `json:"resolved"`
	Assigned    int64  `json:"assigned"`
	Unassigned  int64  `json:"unassigned"`
	// Counts only; a zero denominator means no eligible retained records and does not imply a zero rate.
	LabelCoverage Coverage `json:"label_coverage"`
	// Current resolution of each resolved retained review; prior resolutions and reopened reviews do not count. Missing keys mean zero. These investigation conclusions are independent of account feedback.
	Conclusions ReviewsConclusions `json:"conclusions"`
}

func (r Reviews) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Reviews) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *Reviews) GetPopulation() string {
	return "retained_reviews_by_created_at"
}

func (r *Reviews) GetDenominator() int64 {
	if r == nil {
		return 0
	}
	return r.Denominator
}

func (r *Reviews) GetOpen() int64 {
	if r == nil {
		return 0
	}
	return r.Open
}

func (r *Reviews) GetInProgress() int64 {
	if r == nil {
		return 0
	}
	return r.InProgress
}

func (r *Reviews) GetResolved() int64 {
	if r == nil {
		return 0
	}
	return r.Resolved
}

func (r *Reviews) GetAssigned() int64 {
	if r == nil {
		return 0
	}
	return r.Assigned
}

func (r *Reviews) GetUnassigned() int64 {
	if r == nil {
		return 0
	}
	return r.Unassigned
}

func (r *Reviews) GetLabelCoverage() Coverage {
	if r == nil {
		return Coverage{}
	}
	return r.LabelCoverage
}

func (r *Reviews) GetConclusions() ReviewsConclusions {
	if r == nil {
		return ReviewsConclusions{}
	}
	return r.Conclusions
}
