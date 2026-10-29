package components

type From string

const (
	FromAllow  From = "allow"
	FromStepUp From = "step_up"
	FromReview From = "review"
	FromBlock  From = "block"
	FromHeld   From = "held"
)

func (e From) ToPointer() *From {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *From) IsExact() bool {
	if e != nil {
		switch *e {
		case "allow", "step_up", "review", "block", "held":
			return true
		}
	}
	return false
}
