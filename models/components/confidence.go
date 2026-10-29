package components

// Confidence - Evidence coverage: how much of the planned source evidence gave a definite answer. This is not confidence that an identity is genuine, agreement between identifiers, a risk score or a probability. unavailable means no usable evidence answered; a terminal uncovered assessment can have this value and recommend review.
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
