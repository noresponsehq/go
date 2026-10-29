package components

type UsageCount struct {
	Environment AuditSettingsEnvironment `json:"environment"`
	Day         string                   `json:"day"`
	// assessments: assessments admitted. webhook_deliveries: events sent to endpoints, one per event per endpoint. key_usage_events: key-usage events accepted.
	Metric Metric `json:"metric"`
	// assessments only.
	Flow *UsageCountFlow `json:"flow,omitzero"`
	// assessments only.
	Depth *UsageCountDepth `json:"depth,omitzero"`
	Count int64            `json:"count"`
}

func (u *UsageCount) GetEnvironment() AuditSettingsEnvironment {
	if u == nil {
		return AuditSettingsEnvironment("")
	}
	return u.Environment
}

func (u *UsageCount) GetDay() string {
	if u == nil {
		return ""
	}
	return u.Day
}

func (u *UsageCount) GetMetric() Metric {
	if u == nil {
		return Metric("")
	}
	return u.Metric
}

func (u *UsageCount) GetFlow() *UsageCountFlow {
	if u == nil {
		return nil
	}
	return u.Flow
}

func (u *UsageCount) GetDepth() *UsageCountDepth {
	if u == nil {
		return nil
	}
	return u.Depth
}

func (u *UsageCount) GetCount() int64 {
	if u == nil {
		return 0
	}
	return u.Count
}
