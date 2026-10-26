package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// AuthContext - Normalised claims from a session your backend has already verified with WorkOS, Clerk, Auth0 or any OIDC provider. Pass the output of the SDK's auth adapters, not a raw token.
type AuthContext struct {
	// Lowercase provider slug. WorkOS, Clerk and Auth0 have first-class adapters; every other provider uses a customer-specific OIDC slug. Provider brand is metadata; it does not change risk by itself.
	Provider string `json:"provider"`
	// The verified token issuer, normally the provider's HTTPS issuer URL.
	Issuer string `json:"issuer"`
	// The verified audience/resource identifier used by the customer's API. Use audiences when the provider emits multiple values.
	Audience *string `json:"audience,omitzero"`
	// All verified audience/resource identifiers from an OAuth or OIDC token. The first value may also be copied to audience for simple integrations.
	Audiences []string `json:"audiences,omitzero"`
	// The verified stable subject identifier (`sub`), opaque to NoResponse. Required, except from a provider hook at pre_registration, where no user exists yet.
	Subject string `json:"subject"`
	// The verified session identifier (`sid` or provider equivalent).
	SessionID *string `json:"session_id,omitzero"`
	// The active organization or workspace identifier, when the provider issued one.
	OrganizationID *string `json:"organization_id,omitzero"`
	// The provider workspace identifier, when distinct from organization_id.
	TenantID *string `json:"tenant_id,omitzero"`
	// The verified token identifier (`jti`), used for replay and rotation correlation.
	TokenID *string `json:"token_id,omitzero"`
	// Verified token issue time (`iat`).
	IssuedAt *string `json:"issued_at,omitzero"`
	// Verified token expiry (`exp`).
	ExpiresAt *string `json:"expires_at,omitzero"`
	// Creation time of the provider session from the verified backend session record (for example Clerk Session.createdAt or WorkOS Session.created_at). This is separate from token issuance and auth_time.
	SessionCreatedAt *string `json:"session_created_at,omitzero"`
	// Verified end-user authentication time (`auth_time`), when present.
	AuthTime *string `json:"auth_time,omitzero"`
	// Account creation time from the customer's verified provider record, when available.
	AccountCreatedAt *string `json:"account_created_at,omitzero"`
	// Whether the customer's provider asserts that the email identifier was verified.
	EmailVerified *bool `json:"email_verified,omitzero"`
	// Whether the customer's provider asserts that the phone identifier was verified.
	PhoneVerified *bool `json:"phone_verified,omitzero"`
	// Provider or OIDC assurance level, when the customer has mapped one.
	AssuranceLevel *AssuranceLevel `json:"assurance_level,omitzero"`
	// Whether the verified session is acting on behalf of another subject.
	Impersonated *bool `json:"impersonated,omitzero"`
	// Where the context was derived: backend (also when the request named no transport), browser, or provider_hook for a hosted provider hook.
	Transport *Transport `json:"transport,omitzero"`
	// The provider event a provider_hook context was sent for; absent for any other transport.
	HookPoint *HookPoint `json:"hook_point,omitzero"`
	// Whether the customer completed signature, issuer, audience and expiry validation before calling NoResponse. False contexts are accepted only as weak hints.
	Verified bool `json:"verified"`
	// Normalized multi-factor evidence from a verified provider session.
	Mfa *AuthContextMfa `json:"mfa,omitzero"`
	// Bounded role slugs from the verified provider; roles do not change risk by themselves.
	Roles []string `json:"roles,omitzero"`
	// Bounded scopes from the verified access token; scopes do not change risk by themselves.
	Scopes []string `json:"scopes,omitzero"`
}

func (a AuthContext) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthContext) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuthContext) GetProvider() string {
	if a == nil {
		return ""
	}
	return a.Provider
}

func (a *AuthContext) GetIssuer() string {
	if a == nil {
		return ""
	}
	return a.Issuer
}

func (a *AuthContext) GetAudience() *string {
	if a == nil {
		return nil
	}
	return a.Audience
}

func (a *AuthContext) GetAudiences() []string {
	if a == nil {
		return nil
	}
	return a.Audiences
}

func (a *AuthContext) GetSubject() string {
	if a == nil {
		return ""
	}
	return a.Subject
}

func (a *AuthContext) GetSessionID() *string {
	if a == nil {
		return nil
	}
	return a.SessionID
}

func (a *AuthContext) GetOrganizationID() *string {
	if a == nil {
		return nil
	}
	return a.OrganizationID
}

func (a *AuthContext) GetTenantID() *string {
	if a == nil {
		return nil
	}
	return a.TenantID
}

func (a *AuthContext) GetTokenID() *string {
	if a == nil {
		return nil
	}
	return a.TokenID
}

func (a *AuthContext) GetIssuedAt() *string {
	if a == nil {
		return nil
	}
	return a.IssuedAt
}

func (a *AuthContext) GetExpiresAt() *string {
	if a == nil {
		return nil
	}
	return a.ExpiresAt
}

func (a *AuthContext) GetSessionCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.SessionCreatedAt
}

func (a *AuthContext) GetAuthTime() *string {
	if a == nil {
		return nil
	}
	return a.AuthTime
}

func (a *AuthContext) GetAccountCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.AccountCreatedAt
}

func (a *AuthContext) GetEmailVerified() *bool {
	if a == nil {
		return nil
	}
	return a.EmailVerified
}

func (a *AuthContext) GetPhoneVerified() *bool {
	if a == nil {
		return nil
	}
	return a.PhoneVerified
}

func (a *AuthContext) GetAssuranceLevel() *AssuranceLevel {
	if a == nil {
		return nil
	}
	return a.AssuranceLevel
}

func (a *AuthContext) GetImpersonated() *bool {
	if a == nil {
		return nil
	}
	return a.Impersonated
}

func (a *AuthContext) GetTransport() *Transport {
	if a == nil {
		return nil
	}
	return a.Transport
}

func (a *AuthContext) GetHookPoint() *HookPoint {
	if a == nil {
		return nil
	}
	return a.HookPoint
}

func (a *AuthContext) GetVerified() bool {
	if a == nil {
		return false
	}
	return a.Verified
}

func (a *AuthContext) GetMfa() *AuthContextMfa {
	if a == nil {
		return nil
	}
	return a.Mfa
}

func (a *AuthContext) GetRoles() []string {
	if a == nil {
		return nil
	}
	return a.Roles
}

func (a *AuthContext) GetScopes() []string {
	if a == nil {
		return nil
	}
	return a.Scopes
}
