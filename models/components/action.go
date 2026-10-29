package components

type Action string

const (
	ActionAllow Action = "allow"
	ActionBlock Action = "block"
)

func (e Action) ToPointer() *Action {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Action) IsExact() bool {
	if e != nil {
		switch *e {
		case "allow", "block":
			return true
		}
	}
	return false
}
