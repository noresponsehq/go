package components

// FindingBand - How severe a finding is, most severe first.
type FindingBand string

const (
	FindingBandCritical FindingBand = "critical"
	FindingBandHigh     FindingBand = "high"
	FindingBandModerate FindingBand = "moderate"
	FindingBandLow      FindingBand = "low"
	FindingBandInfo     FindingBand = "info"
)

func (e FindingBand) ToPointer() *FindingBand {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *FindingBand) IsExact() bool {
	if e != nil {
		switch *e {
		case "critical", "high", "moderate", "low", "info":
			return true
		}
	}
	return false
}
