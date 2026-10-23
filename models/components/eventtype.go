package components

type EventType string

const (
	EventTypeAssessmentUpdated EventType = "assessment.updated"
)

func (e EventType) ToPointer() *EventType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EventType) IsExact() bool {
	if e != nil {
		switch *e {
		case "assessment.updated":
			return true
		}
	}
	return false
}
