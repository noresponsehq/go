package components

type InvestigationStatus string

const (
	InvestigationStatusQueued    InvestigationStatus = "queued"
	InvestigationStatusRunning   InvestigationStatus = "running"
	InvestigationStatusCompleted InvestigationStatus = "completed"
	InvestigationStatusPartial   InvestigationStatus = "partial"
	InvestigationStatusFailed    InvestigationStatus = "failed"
	InvestigationStatusCancelled InvestigationStatus = "cancelled"
	InvestigationStatusUnknown   InvestigationStatus = "unknown"
)

func (e InvestigationStatus) ToPointer() *InvestigationStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *InvestigationStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "queued", "running", "completed", "partial", "failed", "cancelled", "unknown":
			return true
		}
	}
	return false
}
