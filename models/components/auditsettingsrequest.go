package components

type AuditSettingsRequest struct {
	RetentionDays int64 `json:"retention_days"`
}

func (a *AuditSettingsRequest) GetRetentionDays() int64 {
	if a == nil {
		return 0
	}
	return a.RetentionDays
}
