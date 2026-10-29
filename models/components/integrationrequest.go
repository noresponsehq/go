package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// IntegrationRequest - POST /v1/integrations body. Registers a hosted endpoint for one authentication provider. The workspace is the API key that signs the request, so the body has no workspace field and refuses one like any other unknown member. A WorkOS integration needs signing_secret, the Actions secret shown in the WorkOS dashboard, which NoResponse uses to verify each Action request and to sign its answer. A Supabase integration needs signing_secret, the hook secret Supabase generates for a Before User Created HTTPS hook (v1,whsec_ and base64), which NoResponse uses to verify each call. An Auth0 integration sends no signing_secret: NoResponse generates the shared secret and returns it once.
type IntegrationRequest struct {
	// The authentication provider that calls the endpoint.
	Platform IntegrationPlatform `json:"platform"`
	// The provider events the endpoint scores. Default every hook point the platform has: pre_registration (WorkOS user registration Action, Auth0 pre-user-registration trigger, Supabase Before User Created hook) and post_login (WorkOS authentication Action, Auth0 post-login trigger). Supabase has only pre_registration, so a Supabase integration that lists post_login is refused. An event that is not listed is answered allow without an assessment.
	HookPoints []IntegrationHookPoint `json:"hook_points,omitzero"`
	// The WorkOS Actions signing secret or the Supabase hook secret, as the provider's dashboard shows it: 16 to 256 printable ASCII characters with no spaces. It is stored sealed and is not returned.
	SigningSecret *string `json:"signing_secret,omitzero"`
}

func (i IntegrationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IntegrationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *IntegrationRequest) GetPlatform() IntegrationPlatform {
	if i == nil {
		return IntegrationPlatform("")
	}
	return i.Platform
}

func (i *IntegrationRequest) GetHookPoints() []IntegrationHookPoint {
	if i == nil {
		return nil
	}
	return i.HookPoints
}

func (i *IntegrationRequest) GetSigningSecret() *string {
	if i == nil {
		return nil
	}
	return i.SigningSecret
}
