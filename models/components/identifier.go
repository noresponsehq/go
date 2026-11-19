package components

type Identifier string

const (
	IdentifierEmail    Identifier = "email"
	IdentifierPhone    Identifier = "phone"
	IdentifierUsername Identifier = "username"
	IdentifierName     Identifier = "name"
)

func (e Identifier) ToPointer() *Identifier {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Identifier) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "username", "name":
			return true
		}
	}
	return false
}
