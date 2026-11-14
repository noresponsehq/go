package components

// PersonBatchItemOutcome - unchanged: the row matched a person with the same details.
type PersonBatchItemOutcome string

const (
	PersonBatchItemOutcomeCreated   PersonBatchItemOutcome = "created"
	PersonBatchItemOutcomeUpdated   PersonBatchItemOutcome = "updated"
	PersonBatchItemOutcomeUnchanged PersonBatchItemOutcome = "unchanged"
)

func (e PersonBatchItemOutcome) ToPointer() *PersonBatchItemOutcome {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PersonBatchItemOutcome) IsExact() bool {
	if e != nil {
		switch *e {
		case "created", "updated", "unchanged":
			return true
		}
	}
	return false
}
