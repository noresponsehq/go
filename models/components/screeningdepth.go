package components

type ScreeningDepth string

const (
	ScreeningDepthStandard ScreeningDepth = "standard"
)

func (e ScreeningDepth) ToPointer() *ScreeningDepth {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningDepth) IsExact() bool {
	if e != nil {
		switch *e {
		case "standard":
			return true
		}
	}
	return false
}
