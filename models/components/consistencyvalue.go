package components

// ConsistencyValue - How a submitted identifier agrees with what the evidence holds for the other identifiers. unknown is a definite 'cannot tell' (for example a truncated or too-common breach result); absent means the comparison was not made.
type ConsistencyValue string

const (
	ConsistencyValueMatch    ConsistencyValue = "match"
	ConsistencyValuePartial  ConsistencyValue = "partial"
	ConsistencyValueMismatch ConsistencyValue = "mismatch"
	ConsistencyValueUnknown  ConsistencyValue = "unknown"
)

func (e ConsistencyValue) ToPointer() *ConsistencyValue {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ConsistencyValue) IsExact() bool {
	if e != nil {
		switch *e {
		case "match", "partial", "mismatch", "unknown":
			return true
		}
	}
	return false
}
