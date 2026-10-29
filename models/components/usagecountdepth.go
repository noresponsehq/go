package components

// UsageCountDepth - assessments only.
type UsageCountDepth string

const (
	UsageCountDepthQuick UsageCountDepth = "quick"
	UsageCountDepthFull  UsageCountDepth = "full"
)

func (e UsageCountDepth) ToPointer() *UsageCountDepth {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *UsageCountDepth) IsExact() bool {
	if e != nil {
		switch *e {
		case "quick", "full":
			return true
		}
	}
	return false
}
