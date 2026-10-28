package components

import (
	"encoding/json"
	"fmt"
)

// Decision - The decision of the assessment's latest stored revision: block when its stored risk.block is true, otherwise its recommendation. This is not the account's current decision or held-admission state. Combined with recommendation using AND.
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
func (e *Decision) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		fallthrough
	case "step_up":
		fallthrough
	case "review":
		fallthrough
	case "block":
		*e = Decision(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Decision: %v", v)
	}
}
