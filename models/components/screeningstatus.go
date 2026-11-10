package components

// ScreeningStatus - partial finished with some attendees or platforms not covered; attendees_failed and attendees_partial count them.
type ScreeningStatus string

const (
	ScreeningStatusQueued    ScreeningStatus = "queued"
	ScreeningStatusRunning   ScreeningStatus = "running"
	ScreeningStatusCompleted ScreeningStatus = "completed"
	ScreeningStatusPartial   ScreeningStatus = "partial"
	ScreeningStatusFailed    ScreeningStatus = "failed"
	ScreeningStatusCancelled ScreeningStatus = "cancelled"
)

func (e ScreeningStatus) ToPointer() *ScreeningStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "queued", "running", "completed", "partial", "failed", "cancelled":
			return true
		}
	}
	return false
}
