package components

// AccountDecision - block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. Your application applies it.
type AccountDecision string

const (
	AccountDecisionAllow  AccountDecision = "allow"
	AccountDecisionStepUp AccountDecision = "step_up"
	AccountDecisionReview AccountDecision = "review"
	AccountDecisionBlock  AccountDecision = "block"
)

func (e AccountDecision) ToPointer() *AccountDecision {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AccountDecision) IsExact() bool {
	if e != nil {
		switch *e {
		case "allow", "step_up", "review", "block":
			return true
		}
	}
	return false
}
