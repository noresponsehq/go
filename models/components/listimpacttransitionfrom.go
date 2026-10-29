package components

type ListImpactTransitionFrom string

const (
	ListImpactTransitionFromNone  ListImpactTransitionFrom = "none"
	ListImpactTransitionFromAllow ListImpactTransitionFrom = "allow"
	ListImpactTransitionFromBlock ListImpactTransitionFrom = "block"
)

func (e ListImpactTransitionFrom) ToPointer() *ListImpactTransitionFrom {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ListImpactTransitionFrom) IsExact() bool {
	if e != nil {
		switch *e {
		case "none", "allow", "block":
			return true
		}
	}
	return false
}
