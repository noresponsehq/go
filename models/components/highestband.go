package components

type HighestBand string

const (
	HighestBandCritical HighestBand = "critical"
	HighestBandHigh     HighestBand = "high"
	HighestBandModerate HighestBand = "moderate"
	HighestBandLow      HighestBand = "low"
	HighestBandInfo     HighestBand = "info"
)

func (e HighestBand) ToPointer() *HighestBand {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *HighestBand) IsExact() bool {
	if e != nil {
		switch *e {
		case "critical", "high", "moderate", "low", "info":
			return true
		}
	}
	return false
}
