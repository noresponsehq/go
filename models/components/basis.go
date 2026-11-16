package components

// Basis - stated: read from a record. inferred: reasoned from several facts. absence: every source answered and none found it.
type Basis string

const (
	BasisStated   Basis = "stated"
	BasisInferred Basis = "inferred"
	BasisAbsence  Basis = "absence"
)

func (e Basis) ToPointer() *Basis {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Basis) IsExact() bool {
	if e != nil {
		switch *e {
		case "stated", "inferred", "absence":
			return true
		}
	}
	return false
}
