package components

import (
	"encoding/json"
	"fmt"
)

// Flow - The provider flow the assessment gates. Use the top-level flow instead; this member is ignored when the top-level flow is present. Absent means signup.
type Flow string

const (
	FlowSignup        Flow = "signup"
	FlowLogin         Flow = "login"
	FlowRecovery      Flow = "recovery"
	FlowProfileChange Flow = "profile_change"
)

func (e Flow) ToPointer() *Flow {
	return &e
}
func (e *Flow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "signup":
		fallthrough
	case "login":
		fallthrough
	case "recovery":
		fallthrough
	case "profile_change":
		*e = Flow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Flow: %v", v)
	}
}
