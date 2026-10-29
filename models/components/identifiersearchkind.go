package components

type IdentifierSearchKind string

const (
	IdentifierSearchKindEmail    IdentifierSearchKind = "email"
	IdentifierSearchKindPhone    IdentifierSearchKind = "phone"
	IdentifierSearchKindUsername IdentifierSearchKind = "username"
)

func (e IdentifierSearchKind) ToPointer() *IdentifierSearchKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *IdentifierSearchKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "username":
			return true
		}
	}
	return false
}
