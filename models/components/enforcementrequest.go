package components

// EnforcementRequest - POST /v1/accounts/{account_id}/enforcements body: what the workspace did to the account in its identity provider after reading the account's risk at risk_version. The account is named by the path and the workspace by the API key, so neither is a body member. An enforcement is the workspace's own action, not an outcome: what the workspace later learned about an assessment goes to POST /v1/assessments/{assessment_id}/outcomes. The record's key is (account, risk_version, flow, action), so a retry of the same report is safe and needs no Idempotency-Key.
type EnforcementRequest struct {
	// The account's risk_version the workspace acted on, as read from GET /v1/accounts/{account_id}, GET /v1/accounts, an assessment answer or account.risk_changed. A version older than the account's current one is accepted and stored; a version newer than the current one answers 409 risk_version_ahead.
	RiskVersion int64 `json:"risk_version"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// What the workspace did in its identity provider: revoke_sessions (signed the account out everywhere), restrict (limited what the account can do), unrestrict (lifted an earlier restriction), fail_onboarding (stopped an account that was still onboarding), reject_activation (refused to activate a provisional account) or none (decided to take no action).
	Action EnforcementAction `json:"action"`
	// The account's state in the workspace's identity provider just before the action: onboarding, provisional (created, not yet activated), active, restricted or rejected.
	StateBefore EnforcementState `json:"state_before"`
	// applied when the identity provider accepted the action, failed when it refused it or could not be reached. Only applied counts as enforced.
	Status EnforcementStatus `json:"status"`
	// The identity provider the action ran in; custom for the workspace's own user store.
	Provider EnforcementProvider `json:"provider"`
	// Why a failed action failed, as a short code of the workspace's choosing (lower-case letters, digits, '_', '.' and '-'), such as the identity provider's error code. Allowed only with status failed: with status applied it answers 400 invalid_request. Keep personal data out of it.
	ErrorCode *string `json:"error_code,omitzero"`
}

func (e *EnforcementRequest) GetRiskVersion() int64 {
	if e == nil {
		return 0
	}
	return e.RiskVersion
}

func (e *EnforcementRequest) GetFlow() Flow {
	if e == nil {
		return Flow("")
	}
	return e.Flow
}

func (e *EnforcementRequest) GetAction() EnforcementAction {
	if e == nil {
		return EnforcementAction("")
	}
	return e.Action
}

func (e *EnforcementRequest) GetStateBefore() EnforcementState {
	if e == nil {
		return EnforcementState("")
	}
	return e.StateBefore
}

func (e *EnforcementRequest) GetStatus() EnforcementStatus {
	if e == nil {
		return EnforcementStatus("")
	}
	return e.Status
}

func (e *EnforcementRequest) GetProvider() EnforcementProvider {
	if e == nil {
		return EnforcementProvider("")
	}
	return e.Provider
}

func (e *EnforcementRequest) GetErrorCode() *string {
	if e == nil {
		return nil
	}
	return e.ErrorCode
}
