package components

type Band string

const (
	BandLow      Band = "low"
	BandMedium   Band = "medium"
	BandHigh     Band = "high"
	BandCritical Band = "critical"
)

func (e Band) ToPointer() *Band {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Band) IsExact() bool {
	if e != nil {
		switch *e {
		case "low", "medium", "high", "critical":
			return true
		}
	}
	return false
}
