package components

// IntegrationPlatform - The authentication provider that calls the endpoint.
type IntegrationPlatform string

const (
	IntegrationPlatformWorkos   IntegrationPlatform = "workos"
	IntegrationPlatformAuth0    IntegrationPlatform = "auth0"
	IntegrationPlatformSupabase IntegrationPlatform = "supabase"
)

func (e IntegrationPlatform) ToPointer() *IntegrationPlatform {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *IntegrationPlatform) IsExact() bool {
	if e != nil {
		switch *e {
		case "workos", "auth0", "supabase":
			return true
		}
	}
	return false
}
