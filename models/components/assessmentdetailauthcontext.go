package components

// AssessmentDetailAuthContext - The identity fields of the auth context the assessment was made with, echoed for the workspace's own correlation: the request's auth_context, or the one a hosted provider hook (POST /hooks/{integration_id}) built, with transport provider_hook and its hook_point. Present exactly when the request carried an auth_context, and the same on every revision of the assessment. Values are as the request or the hook sent them; tokens, timestamps, roles and scopes are not echoed, and the echo does not change the score.
type AssessmentDetailAuthContext struct {
	// The provider slug the context named.
	Provider string `json:"provider"`
	// The issuer the context named.
	Issuer string `json:"issuer"`
	// The subject the context named; absent when it named none (a pre_registration hook).
	Subject *string `json:"subject,omitzero"`
	// The session identifier the context named, when it named one.
	SessionID *string `json:"session_id,omitzero"`
	// The organization identifier the context named, when it named one.
	OrganizationID *string `json:"organization_id,omitzero"`
	// The verified flag the context carried. A provider_hook context is always false.
	Verified bool `json:"verified"`
	// Where the context was derived: backend (also when the request named no transport), browser, or provider_hook for a hosted provider hook.
	Transport Transport `json:"transport"`
	// The provider event a provider_hook context was sent for; absent for any other transport.
	HookPoint *HookPoint `json:"hook_point,omitzero"`
}

func (a *AssessmentDetailAuthContext) GetProvider() string {
	if a == nil {
		return ""
	}
	return a.Provider
}

func (a *AssessmentDetailAuthContext) GetIssuer() string {
	if a == nil {
		return ""
	}
	return a.Issuer
}

func (a *AssessmentDetailAuthContext) GetSubject() *string {
	if a == nil {
		return nil
	}
	return a.Subject
}

func (a *AssessmentDetailAuthContext) GetSessionID() *string {
	if a == nil {
		return nil
	}
	return a.SessionID
}

func (a *AssessmentDetailAuthContext) GetOrganizationID() *string {
	if a == nil {
		return nil
	}
	return a.OrganizationID
}

func (a *AssessmentDetailAuthContext) GetVerified() bool {
	if a == nil {
		return false
	}
	return a.Verified
}

func (a *AssessmentDetailAuthContext) GetTransport() Transport {
	if a == nil {
		return Transport("")
	}
	return a.Transport
}

func (a *AssessmentDetailAuthContext) GetHookPoint() *HookPoint {
	if a == nil {
		return nil
	}
	return a.HookPoint
}
