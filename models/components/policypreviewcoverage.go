package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type PolicyPreviewCoverage struct {
	Scanned                int64 `json:"scanned"`
	Eligible               int64 `json:"eligible"`
	MissingScore           int64 `json:"missing_score"`
	Unsupported            int64 `json:"unsupported"`
	Truncated              bool  `json:"truncated"`
	CompleteRetainedWindow bool  `json:"complete_retained_window"`
	// Deleted, expired and uncaptured records cannot be reconstructed. Counts describe retained evidence and provide no causal estimate of fraud prevented.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	historicalPopulationKnown bool `const:"false" json:"historical_population_known"`
}

func (p PolicyPreviewCoverage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyPreviewCoverage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PolicyPreviewCoverage) GetScanned() int64 {
	if p == nil {
		return 0
	}
	return p.Scanned
}

func (p *PolicyPreviewCoverage) GetEligible() int64 {
	if p == nil {
		return 0
	}
	return p.Eligible
}

func (p *PolicyPreviewCoverage) GetMissingScore() int64 {
	if p == nil {
		return 0
	}
	return p.MissingScore
}

func (p *PolicyPreviewCoverage) GetUnsupported() int64 {
	if p == nil {
		return 0
	}
	return p.Unsupported
}

func (p *PolicyPreviewCoverage) GetTruncated() bool {
	if p == nil {
		return false
	}
	return p.Truncated
}

func (p *PolicyPreviewCoverage) GetCompleteRetainedWindow() bool {
	if p == nil {
		return false
	}
	return p.CompleteRetainedWindow
}

func (p *PolicyPreviewCoverage) GetHistoricalPopulationKnown() bool {
	return false
}
