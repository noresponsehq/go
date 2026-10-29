package components

// Impact - The selected reason's served impact band. Requires reason.
type Impact string

const (
	ImpactNone   Impact = "none"
	ImpactLow    Impact = "low"
	ImpactMedium Impact = "medium"
	ImpactHigh   Impact = "high"
)

func (e Impact) ToPointer() *Impact {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Impact) IsExact() bool {
	if e != nil {
		switch *e {
		case "none", "low", "medium", "high":
			return true
		}
	}
	return false
}
