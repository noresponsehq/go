package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AuditSettings - Read audit retention and available history.
type AuditSettings struct {
	Environment       AuditSettingsEnvironment `json:"environment"`
	RetentionDays     int64                    `json:"retention_days"`
	OldestAvailableAt *string                  `json:"oldest_available_at,omitzero"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	capture string `const:"available_events_only" json:"capture"`
}

func (a AuditSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuditSettings) GetEnvironment() AuditSettingsEnvironment {
	if a == nil {
		return AuditSettingsEnvironment("")
	}
	return a.Environment
}

func (a *AuditSettings) GetRetentionDays() int64 {
	if a == nil {
		return 0
	}
	return a.RetentionDays
}

func (a *AuditSettings) GetOldestAvailableAt() *string {
	if a == nil {
		return nil
	}
	return a.OldestAvailableAt
}

func (a *AuditSettings) GetCapture() string {
	return "available_events_only"
}
