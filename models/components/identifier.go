package components

type Identifier string

const (
	IdentifierEmail    Identifier = "email"
	IdentifierUsername Identifier = "username"
)

func (e Identifier) ToPointer() *Identifier {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Identifier) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "username":
			return true
		}
	}
	return false
}
