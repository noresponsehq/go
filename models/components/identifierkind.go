package components

// IdentifierKind - Which submitted identifier a signal concerns. ip is a selector of breach_exposure only.
type IdentifierKind string

const (
	IdentifierKindEmail    IdentifierKind = "email"
	IdentifierKindUsername IdentifierKind = "username"
	IdentifierKindPhone    IdentifierKind = "phone"
	IdentifierKindIP       IdentifierKind = "ip"
)

func (e IdentifierKind) ToPointer() *IdentifierKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *IdentifierKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "username", "phone", "ip":
			return true
		}
	}
	return false
}
