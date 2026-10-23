package components

// HookPoint - The provider event a provider_hook context was sent for; absent for any other transport.
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

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *HookPoint) IsExact() bool {
	if e != nil {
		switch *e {
		case "pre_registration", "post_registration", "pre_login", "post_login", "token_issuance":
			return true
		}
	}
	return false
}
