package components

// PermissionDeniedErrorCode - Machine codes an HTTP 403 answer can carry.
type PermissionDeniedErrorCode string

const (
	PermissionDeniedErrorCodeWorkspaceForbidden PermissionDeniedErrorCode = "workspace_forbidden"
	PermissionDeniedErrorCodeRoleForbidden      PermissionDeniedErrorCode = "role_forbidden"
	PermissionDeniedErrorCodeEmailUnverified    PermissionDeniedErrorCode = "email_unverified"
	PermissionDeniedErrorCodeOnboardingInReview PermissionDeniedErrorCode = "onboarding_in_review"
	PermissionDeniedErrorCodeOnboardingRefused  PermissionDeniedErrorCode = "onboarding_refused"
)

func (e PermissionDeniedErrorCode) ToPointer() *PermissionDeniedErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PermissionDeniedErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "workspace_forbidden", "role_forbidden", "email_unverified", "onboarding_in_review", "onboarding_refused":
			return true
		}
	}
	return false
}
