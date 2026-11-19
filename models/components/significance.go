package components

// Significance - One rule ranks every change: identity changes (accounts, breaches, identifiers) are high, profile fields are medium, new posts are low.
type Significance string

const (
	SignificanceHigh   Significance = "high"
	SignificanceMedium Significance = "medium"
	SignificanceLow    Significance = "low"
)

func (e Significance) ToPointer() *Significance {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Significance) IsExact() bool {
	if e != nil {
		switch *e {
		case "high", "medium", "low":
			return true
		}
	}
	return false
}
