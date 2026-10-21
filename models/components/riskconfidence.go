package components

// RiskConfidence - How much of the planned evidence answered; it says nothing about risk.
type RiskConfidence string

const (
	RiskConfidenceLow    RiskConfidence = "low"
	RiskConfidenceMedium RiskConfidence = "medium"
	RiskConfidenceHigh   RiskConfidence = "high"
)

func (e RiskConfidence) ToPointer() *RiskConfidence {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RiskConfidence) IsExact() bool {
	if e != nil {
		switch *e {
		case "low", "medium", "high":
			return true
		}
	}
	return false
}
