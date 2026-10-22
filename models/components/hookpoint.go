package components

import (
	"encoding/json"
	"fmt"
)

// HookPoint - The provider event a provider_hook context was sent for. pre_registration runs before the user exists, so subject and the request's account_id may be absent there. Sent only with transport provider_hook.
type HookPoint string

const (
	HookPointPreRegistration  HookPoint = "pre_registration"
	HookPointPostRegistration HookPoint = "post_registration"
	HookPointPreLogin         HookPoint = "pre_login"
	HookPointPostLogin        HookPoint = "post_login"
	HookPointTokenIssuance    HookPoint = "token_issuance"
)

func (e HookPoint) ToPointer() *HookPoint {
	return &e
}
func (e *HookPoint) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pre_registration":
		fallthrough
	case "post_registration":
		fallthrough
	case "pre_login":
		fallthrough
	case "post_login":
		fallthrough
	case "token_issuance":
		*e = HookPoint(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HookPoint: %v", v)
	}
}
