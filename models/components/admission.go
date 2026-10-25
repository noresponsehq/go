package components

// Admission - immediate (the default): an assessment of this flow is decided from the evidence its answer has, and later evidence may change the decision. held: while the account's current assessment of this flow has not completed, the account's decision for the flow is held, and the final decision arrives as account.risk_changed when the investigation completes or hold_seconds pass, whichever is first. Each flow has its own admission.
type Admission string

const (
	AdmissionImmediate Admission = "immediate"
	AdmissionHeld      Admission = "held"
)

func (e Admission) ToPointer() *Admission {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Admission) IsExact() bool {
	if e != nil {
		switch *e {
		case "immediate", "held":
			return true
		}
	}
	return false
}
