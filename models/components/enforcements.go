package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// Enforcements - Retained customer enforcement report rows created in [from,to), one row per account/risk_version/flow/action. A later applied report can replace failed at the original created_at; retries are not additional attempts here. applied_coverage divides reported_applied by report rows, including explicit action none; it is not independently verified execution, all requested actions, or prevented fraud. Unreported actions are unknown.
type Enforcements struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	population      string `const:"retained_enforcement_reports_by_created_at" json:"population"`
	Denominator     int64  `json:"denominator"`
	Accounts        int64  `json:"accounts"`
	ReportedApplied int64  `json:"reported_applied"`
	ReportedFailed  int64  `json:"reported_failed"`
	// Counts only; a zero denominator means no eligible retained records and does not imply a zero rate.
	AppliedCoverage Coverage `json:"applied_coverage"`
}

func (e Enforcements) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *Enforcements) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *Enforcements) GetPopulation() string {
	return "retained_enforcement_reports_by_created_at"
}

func (e *Enforcements) GetDenominator() int64 {
	if e == nil {
		return 0
	}
	return e.Denominator
}

func (e *Enforcements) GetAccounts() int64 {
	if e == nil {
		return 0
	}
	return e.Accounts
}

func (e *Enforcements) GetReportedApplied() int64 {
	if e == nil {
		return 0
	}
	return e.ReportedApplied
}

func (e *Enforcements) GetReportedFailed() int64 {
	if e == nil {
		return 0
	}
	return e.ReportedFailed
}

func (e *Enforcements) GetAppliedCoverage() Coverage {
	if e == nil {
		return Coverage{}
	}
	return e.AppliedCoverage
}
