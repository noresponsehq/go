package components

// Integration - A hosted provider integration. Point the provider at url: a WorkOS user registration or authentication Action, or the NoResponse Auth0 Action, which sends secret as a bearer token. The answer allows or denies the event: deny only when the assessment's risk.block is true, and allow on every failure, so a fault on the NoResponse side lets the user through. secret is present only in the answers of POST /v1/integrations and POST /v1/integrations/{integration_id}/rotate for an Auth0 integration, and is not returned again.
type Integration struct {
	// Opaque, server-assigned integration id: int_ and a 26-character Crockford base32 ULID. Another workspace's integration id answers 404.
	IntegrationID string `json:"integration_id"`
	// The authentication provider that calls the endpoint.
	Platform IntegrationPlatform `json:"platform"`
	// Each hook point at most once.
	HookPoints []IntegrationHookPoint `json:"hook_points"`
	// The hosted endpoint, https://api.noresponse.org/hooks/{integration_id} in production.
	URL string `json:"url"`
	// Auth0 only, and only when it is created or rotated: the shared secret the Auth0 Action sends as Authorization: Bearer <secret>. nri_ and 43 unpadded base64url characters (32 random bytes). Store it as an Auth0 Action secret.
	Secret    *string `json:"secret,omitzero"`
	CreatedAt string  `json:"created_at"`
	// When the secret was last rotated; absent until the first rotation.
	RotatedAt *string `json:"rotated_at,omitzero"`
	// After a rotation, when the replaced secret stops being accepted; absent when no replaced secret is still accepted.
	PreviousSecretExpiresAt *string `json:"previous_secret_expires_at,omitzero"`
}

func (i *Integration) GetIntegrationID() string {
	if i == nil {
		return ""
	}
	return i.IntegrationID
}

func (i *Integration) GetPlatform() IntegrationPlatform {
	if i == nil {
		return IntegrationPlatform("")
	}
	return i.Platform
}

func (i *Integration) GetHookPoints() []IntegrationHookPoint {
	if i == nil {
		return []IntegrationHookPoint{}
	}
	return i.HookPoints
}

func (i *Integration) GetURL() string {
	if i == nil {
		return ""
	}
	return i.URL
}

func (i *Integration) GetSecret() *string {
	if i == nil {
		return nil
	}
	return i.Secret
}

func (i *Integration) GetCreatedAt() string {
	if i == nil {
		return ""
	}
	return i.CreatedAt
}

func (i *Integration) GetRotatedAt() *string {
	if i == nil {
		return nil
	}
	return i.RotatedAt
}

func (i *Integration) GetPreviousSecretExpiresAt() *string {
	if i == nil {
		return nil
	}
	return i.PreviousSecretExpiresAt
}
