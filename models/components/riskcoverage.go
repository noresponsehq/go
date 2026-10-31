package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// RiskCoverage - Counts of planned source-identifier pairs by evidence outcome, deduplicated across signals and incomplete entries. These counts do not measure identity agreement or prove that processing is still running. A lane that was not searched is not an answered source.
type RiskCoverage struct {
	// Source-identifier pairs that gave a definite found or not_found answer.
	Answered int64 `json:"answered"`
	// Source-identifier pairs that could not answer, including unsupported coverage markers and source failures. Read detailed signals to distinguish Not covered from Failed.
	Failed int64 `json:"failed"`
	// Source-identifier pairs with deferred or timed_out evidence. A terminal timeout can remain here after processing finishes; use assessment status and continuation metadata to determine whether work is still running.
	Pending int64 `json:"pending"`
}

func (r RiskCoverage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskCoverage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskCoverage) GetAnswered() int64 {
	if r == nil {
		return 0
	}
	return r.Answered
}

func (r *RiskCoverage) GetFailed() int64 {
	if r == nil {
		return 0
	}
	return r.Failed
}

func (r *RiskCoverage) GetPending() int64 {
	if r == nil {
		return 0
	}
	return r.Pending
}
