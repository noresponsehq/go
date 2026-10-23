package components

import (
	"encoding/json"
	"fmt"
)

// AssessmentFlow - The customer event an assessment is made for: signup (a new account), login (a returning sign-in) or purchase (a payment or credit grant). The account's risk is reported with the decision for this flow.
type AssessmentFlow string

const (
	AssessmentFlowSignup   AssessmentFlow = "signup"
	AssessmentFlowLogin    AssessmentFlow = "login"
	AssessmentFlowPurchase AssessmentFlow = "purchase"
)

func (e AssessmentFlow) ToPointer() *AssessmentFlow {
	return &e
}
func (e *AssessmentFlow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "signup":
		fallthrough
	case "login":
		fallthrough
	case "purchase":
		*e = AssessmentFlow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssessmentFlow: %v", v)
	}
}
