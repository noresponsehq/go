package components

// EnforcementStatus - applied when the identity provider accepted the action, failed when it refused it or could not be reached. Only applied counts as enforced.
type EnforcementStatus string

const (
	EnforcementStatusApplied EnforcementStatus = "applied"
	EnforcementStatusFailed  EnforcementStatus = "failed"
)

func (e EnforcementStatus) ToPointer() *EnforcementStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EnforcementStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "applied", "failed":
			return true
		}
	}
	return false
}
