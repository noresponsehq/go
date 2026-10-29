package components

type OperationalAggregateBucket struct {
	Start            string `json:"start"`
	Attempts         int64  `json:"attempts"`
	Succeeded        int64  `json:"succeeded"`
	Failed           int64  `json:"failed"`
	ClientErrors     int64  `json:"client_errors"`
	ServerErrors     int64  `json:"server_errors"`
	Replays          int64  `json:"replays"`
	IncompleteWrites int64  `json:"incomplete_writes"`
	UnattributedKeys int64  `json:"unattributed_keys"`
	// Continuous percentile over retained attempts. Omitted when there are no captured samples.
	P50Ms *float64 `json:"p50_ms,omitzero"`
	// Continuous percentile over retained attempts. Omitted when there are no captured samples.
	P95Ms *float64 `json:"p95_ms,omitzero"`
	// Continuous percentile over retained attempts. Omitted when there are no captured samples.
	P99Ms *float64 `json:"p99_ms,omitzero"`
}

func (o *OperationalAggregateBucket) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}

func (o *OperationalAggregateBucket) GetAttempts() int64 {
	if o == nil {
		return 0
	}
	return o.Attempts
}

func (o *OperationalAggregateBucket) GetSucceeded() int64 {
	if o == nil {
		return 0
	}
	return o.Succeeded
}

func (o *OperationalAggregateBucket) GetFailed() int64 {
	if o == nil {
		return 0
	}
	return o.Failed
}

func (o *OperationalAggregateBucket) GetClientErrors() int64 {
	if o == nil {
		return 0
	}
	return o.ClientErrors
}

func (o *OperationalAggregateBucket) GetServerErrors() int64 {
	if o == nil {
		return 0
	}
	return o.ServerErrors
}

func (o *OperationalAggregateBucket) GetReplays() int64 {
	if o == nil {
		return 0
	}
	return o.Replays
}

func (o *OperationalAggregateBucket) GetIncompleteWrites() int64 {
	if o == nil {
		return 0
	}
	return o.IncompleteWrites
}

func (o *OperationalAggregateBucket) GetUnattributedKeys() int64 {
	if o == nil {
		return 0
	}
	return o.UnattributedKeys
}

func (o *OperationalAggregateBucket) GetP50Ms() *float64 {
	if o == nil {
		return nil
	}
	return o.P50Ms
}

func (o *OperationalAggregateBucket) GetP95Ms() *float64 {
	if o == nil {
		return nil
	}
	return o.P95Ms
}

func (o *OperationalAggregateBucket) GetP99Ms() *float64 {
	if o == nil {
		return nil
	}
	return o.P99Ms
}
