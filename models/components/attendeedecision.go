package components

type AttendeeDecision string

const (
	AttendeeDecisionCleared   AttendeeDecision = "cleared"
	AttendeeDecisionWatch     AttendeeDecision = "watch"
	AttendeeDecisionEscalated AttendeeDecision = "escalated"
)

func (e AttendeeDecision) ToPointer() *AttendeeDecision {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AttendeeDecision) IsExact() bool {
	if e != nil {
		switch *e {
		case "cleared", "watch", "escalated":
			return true
		}
	}
	return false
}
