package components

// Decision - An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.
type Decision string

const (
	DecisionAllow  Decision = "allow"
	DecisionStepUp Decision = "step_up"
	DecisionReview Decision = "review"
	DecisionBlock  Decision = "block"
)

func (e Decision) ToPointer() *Decision {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Decision) IsExact() bool {
	if e != nil {
		switch *e {
		case "allow", "step_up", "review", "block":
			return true
		}
	}
	return false
}
