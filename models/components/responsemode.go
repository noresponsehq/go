package components

import (
	"encoding/json"
	"fmt"
)

// ResponseMode - standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
type ResponseMode string

const (
	ResponseModeStandard ResponseMode = "standard"
	ResponseModeDetailed ResponseMode = "detailed"
)

func (e ResponseMode) ToPointer() *ResponseMode {
	return &e
}
func (e *ResponseMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "detailed":
		*e = ResponseMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseMode: %v", v)
	}
}
