package components

// Enforcement - The answer of POST /v1/accounts/{account_id}/enforcements: one enforcement of one of the calling workspace's accounts, kept once per (account, risk_version, flow, action). The first report answers 201 with the new record. A repeat answers 200 with the record as first stored, whatever the repeat's state_before, status, provider or error_code. The one exception is a stored failed record followed by an applied report: the record becomes applied, keeps its id and created_at, takes the applied report's state_before and provider, drops error_code and answers 201, so a successful retry after a failure is recorded. Only an applied record at the account's current risk_version takes the account out of GET /v1/accounts?unenforced=true.
type Enforcement struct {
	// Opaque, server-assigned id of one recorded enforcement: 'enf_' and 26 Crockford base32 characters. Unique within the workspace.
	ID string `json:"id"`
	// The account_id of the request path.
	AccountID string `json:"account_id"`
	// The account's risk_version the workspace acted on. It can be older than the account's current risk_version: such a record is kept and does not count for unenforced.
	RiskVersion int64 `json:"risk_version"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// What the workspace did in its identity provider: revoke_sessions (signed the account out everywhere), restrict (limited what the account can do), unrestrict (lifted an earlier restriction), fail_onboarding (stopped an account that was still onboarding), reject_activation (refused to activate a provisional account), activate (let a held onboarding or provisional account in) or none (decided to take no action).
	Action EnforcementAction `json:"action"`
	// The account's state in the workspace's identity provider just before the action: onboarding, provisional (created, not yet activated), active, restricted or rejected.
	StateBefore EnforcementState `json:"state_before"`
	// applied when the identity provider accepted the action, failed when it refused it or could not be reached. Only applied counts as enforced.
	Status EnforcementStatus `json:"status"`
	// The identity provider the action ran in; custom for the workspace's own user store.
	Provider EnforcementProvider `json:"provider"`
	// Why a failed action failed, as a short code of the workspace's choosing (lower-case letters, digits, '_', '.' and '-'), such as the identity provider's error code. Allowed only with status failed: with status applied it answers 400 invalid_request. Keep personal data out of it.
	ErrorCode *string `json:"error_code,omitzero"`
	// When NoResponse first recorded this enforcement, in UTC. A failed record that becomes applied keeps it.
	CreatedAt string `json:"created_at"`
}

func (e *Enforcement) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *Enforcement) GetAccountID() string {
	if e == nil {
		return ""
	}
	return e.AccountID
}

func (e *Enforcement) GetRiskVersion() int64 {
	if e == nil {
		return 0
	}
	return e.RiskVersion
}

func (e *Enforcement) GetFlow() Flow {
	if e == nil {
		return Flow("")
	}
	return e.Flow
}

func (e *Enforcement) GetAction() EnforcementAction {
	if e == nil {
		return EnforcementAction("")
	}
	return e.Action
}

func (e *Enforcement) GetStateBefore() EnforcementState {
	if e == nil {
		return EnforcementState("")
	}
	return e.StateBefore
}

func (e *Enforcement) GetStatus() EnforcementStatus {
	if e == nil {
		return EnforcementStatus("")
	}
	return e.Status
}

func (e *Enforcement) GetProvider() EnforcementProvider {
	if e == nil {
		return EnforcementProvider("")
	}
	return e.Provider
}

func (e *Enforcement) GetErrorCode() *string {
	if e == nil {
		return nil
	}
	return e.ErrorCode
}

func (e *Enforcement) GetCreatedAt() string {
	if e == nil {
		return ""
	}
	return e.CreatedAt
}
