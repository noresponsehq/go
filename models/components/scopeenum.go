package components

import (
	"encoding/json"
	"fmt"
)

// ScopeEnum - assessment requires one submission to satisfy the complete expression. account permits witnesses across retained submissions of one account and is supported only by POST /v1/accounts/search. No default: select scope explicitly.
type ScopeEnum string

const (
	ScopeEnumAssessment ScopeEnum = "assessment"
	ScopeEnumAccount    ScopeEnum = "account"
)

func (e ScopeEnum) ToPointer() *ScopeEnum {
	return &e
}
func (e *ScopeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "assessment":
		fallthrough
	case "account":
		*e = ScopeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScopeEnum: %v", v)
	}
}
