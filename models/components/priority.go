package components

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityNormal Priority = "normal"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

func (e Priority) ToPointer() *Priority {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Priority) IsExact() bool {
	if e != nil {
		switch *e {
		case "low", "normal", "high", "urgent":
			return true
		}
	}
	return false
}
