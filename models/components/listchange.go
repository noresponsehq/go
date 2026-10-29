package components

// ListChange - A committed bulk list change and bounded notification progress, available while propagation is pending and for 24 hours after all propagation completes. References count an account once per changed list entry and are not unique accounts. No endpoint delivery or customer enforcement is implied by completed.
type ListChange struct {
	ID                         string          `json:"id"`
	CreatedAt                  string          `json:"created_at"`
	State                      ListChangeState `json:"state"`
	Added                      []ListEntry     `json:"added"`
	Removed                    []string        `json:"removed"`
	PendingPropagations        int64           `json:"pending_propagations"`
	ProcessedAccountReferences int64           `json:"processed_account_references"`
}

func (l *ListChange) GetID() string {
	if l == nil {
		return ""
	}
	return l.ID
}

func (l *ListChange) GetCreatedAt() string {
	if l == nil {
		return ""
	}
	return l.CreatedAt
}

func (l *ListChange) GetState() ListChangeState {
	if l == nil {
		return ListChangeState("")
	}
	return l.State
}

func (l *ListChange) GetAdded() []ListEntry {
	if l == nil {
		return []ListEntry{}
	}
	return l.Added
}

func (l *ListChange) GetRemoved() []string {
	if l == nil {
		return []string{}
	}
	return l.Removed
}

func (l *ListChange) GetPendingPropagations() int64 {
	if l == nil {
		return 0
	}
	return l.PendingPropagations
}

func (l *ListChange) GetProcessedAccountReferences() int64 {
	if l == nil {
		return 0
	}
	return l.ProcessedAccountReferences
}
