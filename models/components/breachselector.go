package components

// BreachSelector - The submitted email or username whose own search matched the table. A submitted full name is searched as well, and what it finds feeds the identity's reasons rather than this list.
type BreachSelector string

const (
	BreachSelectorEmail    BreachSelector = "email"
	BreachSelectorUsername BreachSelector = "username"
	BreachSelectorIP       BreachSelector = "ip"
)

func (e BreachSelector) ToPointer() *BreachSelector {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *BreachSelector) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "username", "ip":
			return true
		}
	}
	return false
}
