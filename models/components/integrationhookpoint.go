package components

// IntegrationHookPoint - pre_registration runs before the user exists; post_login runs after a successful sign-in.
type IntegrationHookPoint string

const (
	IntegrationHookPointPreRegistration IntegrationHookPoint = "pre_registration"
	IntegrationHookPointPostLogin       IntegrationHookPoint = "post_login"
)

func (e IntegrationHookPoint) ToPointer() *IntegrationHookPoint {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *IntegrationHookPoint) IsExact() bool {
	if e != nil {
		switch *e {
		case "pre_registration", "post_login":
			return true
		}
	}
	return false
}
