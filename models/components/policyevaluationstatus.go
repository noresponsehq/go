package components

type PolicyEvaluationStatus string

const (
	PolicyEvaluationStatusQueued    PolicyEvaluationStatus = "queued"
	PolicyEvaluationStatusRunning   PolicyEvaluationStatus = "running"
	PolicyEvaluationStatusCompleted PolicyEvaluationStatus = "completed"
	PolicyEvaluationStatusFailed    PolicyEvaluationStatus = "failed"
	PolicyEvaluationStatusCancelled PolicyEvaluationStatus = "cancelled"
)

func (e PolicyEvaluationStatus) ToPointer() *PolicyEvaluationStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PolicyEvaluationStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "queued", "running", "completed", "failed", "cancelled":
			return true
		}
	}
	return false
}
