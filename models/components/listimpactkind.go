package components

type ListImpactKind string

const (
	ListImpactKindEmail       ListImpactKind = "email"
	ListImpactKindEmailDomain ListImpactKind = "email_domain"
)

func (e ListImpactKind) ToPointer() *ListImpactKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ListImpactKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "email_domain":
			return true
		}
	}
	return false
}
