package components

// Confidence - Ordinal confidence in the recommendation under the current policy: low, medium, high or unavailable.
type Confidence string

const (
	ConfidenceLow         Confidence = "low"
	ConfidenceMedium      Confidence = "medium"
	ConfidenceHigh        Confidence = "high"
	ConfidenceUnavailable Confidence = "unavailable"
)

func (e Confidence) ToPointer() *Confidence {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Confidence) IsExact() bool {
	if e != nil {
		switch *e {
		case "low", "medium", "high", "unavailable":
			return true
		}
	}
	return false
}
