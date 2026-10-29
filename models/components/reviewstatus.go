package components

type ReviewStatus string

const (
	ReviewStatusOpen       ReviewStatus = "open"
	ReviewStatusInProgress ReviewStatus = "in_progress"
	ReviewStatusResolved   ReviewStatus = "resolved"
)

func (e ReviewStatus) ToPointer() *ReviewStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ReviewStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "open", "in_progress", "resolved":
			return true
		}
	}
	return false
}
