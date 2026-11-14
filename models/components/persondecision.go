package components

type PersonDecision string

const (
	PersonDecisionCleared   PersonDecision = "cleared"
	PersonDecisionWatch     PersonDecision = "watch"
	PersonDecisionEscalated PersonDecision = "escalated"
)

func (e PersonDecision) ToPointer() *PersonDecision {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PersonDecision) IsExact() bool {
	if e != nil {
		switch *e {
		case "cleared", "watch", "escalated":
			return true
		}
	}
	return false
}
