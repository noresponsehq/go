package components

// EnforcementAction - What the workspace did in its identity provider: revoke_sessions (signed the account out everywhere), restrict (limited what the account can do), unrestrict (lifted an earlier restriction), fail_onboarding (stopped an account that was still onboarding), reject_activation (refused to activate a provisional account), activate (let a held onboarding or provisional account in) or none (decided to take no action).
type EnforcementAction string

const (
	EnforcementActionRevokeSessions   EnforcementAction = "revoke_sessions"
	EnforcementActionRestrict         EnforcementAction = "restrict"
	EnforcementActionUnrestrict       EnforcementAction = "unrestrict"
	EnforcementActionFailOnboarding   EnforcementAction = "fail_onboarding"
	EnforcementActionRejectActivation EnforcementAction = "reject_activation"
	EnforcementActionActivate         EnforcementAction = "activate"
	EnforcementActionNone             EnforcementAction = "none"
)

func (e EnforcementAction) ToPointer() *EnforcementAction {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *EnforcementAction) IsExact() bool {
	if e != nil {
		switch *e {
		case "revoke_sessions", "restrict", "unrestrict", "fail_onboarding", "reject_activation", "activate", "none":
			return true
		}
	}
	return false
}
