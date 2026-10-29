package components

// Provenance - Customer-submitted input. It does not mean the customer or NoResponse verified ownership.
type Provenance string

const (
	ProvenanceSubmitted Provenance = "submitted"
)

func (e Provenance) ToPointer() *Provenance {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Provenance) IsExact() bool {
	if e != nil {
		switch *e {
		case "submitted":
			return true
		}
	}
	return false
}
