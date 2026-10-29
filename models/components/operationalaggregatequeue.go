package components

// OperationalAggregateQueue - Current retained delivery queue, independent of the requested attempt window and status filter. Endpoint filter applies. Does not prove customer enforcement execution.
type OperationalAggregateQueue struct {
	Pending         int64   `json:"pending"`
	Leased          int64   `json:"leased"`
	DeadRetained    int64   `json:"dead_retained"`
	OldestPendingAt *string `json:"oldest_pending_at,omitzero"`
}

func (o *OperationalAggregateQueue) GetPending() int64 {
	if o == nil {
		return 0
	}
	return o.Pending
}

func (o *OperationalAggregateQueue) GetLeased() int64 {
	if o == nil {
		return 0
	}
	return o.Leased
}

func (o *OperationalAggregateQueue) GetDeadRetained() int64 {
	if o == nil {
		return 0
	}
	return o.DeadRetained
}

func (o *OperationalAggregateQueue) GetOldestPendingAt() *string {
	if o == nil {
		return nil
	}
	return o.OldestPendingAt
}
