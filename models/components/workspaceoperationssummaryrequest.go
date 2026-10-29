package components

// WorkspaceOperationsSummaryRequest - Read a past [from,to) interval of at most 31 days. The complete retained populations must fit a combined 100000-row budget and a 1500ms SQL statement budget. No pagination or sampled totals; choose a shorter interval after a row-budget refusal.
type WorkspaceOperationsSummaryRequest struct {
	Environment AuditSettingsEnvironment `json:"environment"`
	From        string                   `json:"from"`
	To          string                   `json:"to"`
}

func (w *WorkspaceOperationsSummaryRequest) GetEnvironment() AuditSettingsEnvironment {
	if w == nil {
		return AuditSettingsEnvironment("")
	}
	return w.Environment
}

func (w *WorkspaceOperationsSummaryRequest) GetFrom() string {
	if w == nil {
		return ""
	}
	return w.From
}

func (w *WorkspaceOperationsSummaryRequest) GetTo() string {
	if w == nil {
		return ""
	}
	return w.To
}
