package components

// EnforcementState - The account's state in the workspace's identity provider just before the action: onboarding, provisional (created, not yet activated), active, restricted or rejected.
type EnforcementState string

const (
	EnforcementStateOnboarding  EnforcementState = "onboarding"
	EnforcementStateProvisional EnforcementState = "provisional"
	EnforcementStateActive      EnforcementState = "active"
	EnforcementStateRestricted  EnforcementState = "restricted"
	EnforcementStateRejected    EnforcementState = "rejected"
)

func (e EnforcementState) ToPointer() *EnforcementState {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EnforcementState) IsExact() bool {
	if e != nil {
		switch *e {
		case "onboarding", "provisional", "active", "restricted", "rejected":
			return true
		}
	}
	return false
}
