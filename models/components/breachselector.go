package components

// BreachSelector - The submitted identifier a breach table matched on: email or username. Phone and name are not breach selectors.
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
