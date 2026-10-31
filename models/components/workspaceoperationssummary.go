package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// WorkspaceOperationsSummary - Three separately defined populations with independent denominators. Interpret each section within its own population; their totals do not form a conversion funnel. All counts and retention come from one snapshot. Requests above the population or query budget are refused without truncation.
type WorkspaceOperationsSummary struct {
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID   string                   `json:"workspace_id"`
	Environment   AuditSettingsEnvironment `json:"environment"`
	From          string                   `json:"from"`
	To            string                   `json:"to"`
	AsOf          string                   `json:"as_of"`
	RetainedFrom  string                   `json:"retained_from"`
	RetentionDays int64                    `json:"retention_days"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	rowBudget   int64 `const:"100000" json:"row_budget"`
	RecordsRead int64 `json:"records_read"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	truncated bool `const:"false" json:"truncated"`
	// Accounts whose current completed assessment was created in the selected window and remains retained at as_of. The effective account feedback head is observed at as_of, including explicit unknown reports. Customer-reported labels and unlabelled accounts have separate counts. Superseded labels are excluded.
	Outcomes Outcomes `json:"outcomes"`
	// Retained customer enforcement report rows created in [from,to), one row per account/risk_version/flow/action. A later applied report can replace failed at the original created_at; retries are not additional attempts here. applied_coverage divides reported_applied by report rows, including explicit action none; it is not independently verified execution, all requested actions, or prevented fraud. Unreported actions are unknown.
	Enforcements Enforcements `json:"enforcements"`
	// Reviews created in [from,to) whose review and current account assessment remain retained. Status and assignment are current at as_of, not a transition history. label_coverage numerator is resolved reviews labeled benign or abuse; denominator is resolved reviews. Inconclusive is not a known outcome. Review conclusions do not change account feedback or enforcement.
	Reviews Reviews `json:"reviews"`
}

func (w WorkspaceOperationsSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspaceOperationsSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WorkspaceOperationsSummary) GetWorkspaceID() string {
	if w == nil {
		return ""
	}
	return w.WorkspaceID
}

func (w *WorkspaceOperationsSummary) GetEnvironment() AuditSettingsEnvironment {
	if w == nil {
		return AuditSettingsEnvironment("")
	}
	return w.Environment
}

func (w *WorkspaceOperationsSummary) GetFrom() string {
	if w == nil {
		return ""
	}
	return w.From
}

func (w *WorkspaceOperationsSummary) GetTo() string {
	if w == nil {
		return ""
	}
	return w.To
}

func (w *WorkspaceOperationsSummary) GetAsOf() string {
	if w == nil {
		return ""
	}
	return w.AsOf
}

func (w *WorkspaceOperationsSummary) GetRetainedFrom() string {
	if w == nil {
		return ""
	}
	return w.RetainedFrom
}

func (w *WorkspaceOperationsSummary) GetRetentionDays() int64 {
	if w == nil {
		return 0
	}
	return w.RetentionDays
}

func (w *WorkspaceOperationsSummary) GetRowBudget() int64 {
	return 100000
}

func (w *WorkspaceOperationsSummary) GetRecordsRead() int64 {
	if w == nil {
		return 0
	}
	return w.RecordsRead
}

func (w *WorkspaceOperationsSummary) GetTruncated() bool {
	return false
}

func (w *WorkspaceOperationsSummary) GetOutcomes() Outcomes {
	if w == nil {
		return Outcomes{}
	}
	return w.Outcomes
}

func (w *WorkspaceOperationsSummary) GetEnforcements() Enforcements {
	if w == nil {
		return Enforcements{}
	}
	return w.Enforcements
}

func (w *WorkspaceOperationsSummary) GetReviews() Reviews {
	if w == nil {
		return Reviews{}
	}
	return w.Reviews
}
