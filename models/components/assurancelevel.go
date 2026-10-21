package components

import (
	"encoding/json"
	"fmt"
)

// AssuranceLevel - Provider or OIDC assurance level, when the customer has mapped one.
type AssuranceLevel string

const (
	AssuranceLevelUnknown AssuranceLevel = "unknown"
	AssuranceLevelLow     AssuranceLevel = "low"
	AssuranceLevelMedium  AssuranceLevel = "medium"
	AssuranceLevelHigh    AssuranceLevel = "high"
)

func (e AssuranceLevel) ToPointer() *AssuranceLevel {
	return &e
}
func (e *AssuranceLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		fallthrough
	case "low":
		fallthrough
	case "medium":
		fallthrough
	case "high":
		*e = AssuranceLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssuranceLevel: %v", v)
	}
}
