package components

// AttendeeBatchItemOutcome - unchanged: the row matched an attendee with the same details.
type AttendeeBatchItemOutcome string

const (
	AttendeeBatchItemOutcomeCreated   AttendeeBatchItemOutcome = "created"
	AttendeeBatchItemOutcomeUpdated   AttendeeBatchItemOutcome = "updated"
	AttendeeBatchItemOutcomeUnchanged AttendeeBatchItemOutcome = "unchanged"
)

func (e AttendeeBatchItemOutcome) ToPointer() *AttendeeBatchItemOutcome {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AttendeeBatchItemOutcome) IsExact() bool {
	if e != nil {
		switch *e {
		case "created", "updated", "unchanged":
			return true
		}
	}
	return false
}
