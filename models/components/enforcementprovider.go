package components

// EnforcementProvider - The identity provider the action ran in; custom for the workspace's own user store.
type EnforcementProvider string

const (
	EnforcementProviderClerk  EnforcementProvider = "clerk"
	EnforcementProviderWorkos EnforcementProvider = "workos"
	EnforcementProviderAuth0  EnforcementProvider = "auth0"
	EnforcementProviderCustom EnforcementProvider = "custom"
)

func (e EnforcementProvider) ToPointer() *EnforcementProvider {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EnforcementProvider) IsExact() bool {
	if e != nil {
		switch *e {
		case "clerk", "workos", "auth0", "custom":
			return true
		}
	}
	return false
}
