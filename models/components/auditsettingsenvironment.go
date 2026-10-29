package components

type AuditSettingsEnvironment string

const (
	AuditSettingsEnvironmentLive AuditSettingsEnvironment = "live"
	AuditSettingsEnvironmentTest AuditSettingsEnvironment = "test"
)

func (e AuditSettingsEnvironment) ToPointer() *AuditSettingsEnvironment {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AuditSettingsEnvironment) IsExact() bool {
	if e != nil {
		switch *e {
		case "live", "test":
			return true
		}
	}
	return false
}
