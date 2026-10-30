package components

// AccountOutcome - Immutable acceptance receipt. risk_version is the account version after this report committed, not the version when a retry reads it. effective_outcome_id names the effective head at that write. The effective expiration may shorten under retention changes. A legitimate report or correction removes only the account-feedback contribution, not unrelated risk.
type AccountOutcome struct {
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	ID string `json:"id"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this workspace for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
	Outcome Outcome `json:"outcome"`
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	Supersedes *string `json:"supersedes,omitzero"`
	ObservedAt string  `json:"observed_at"`
	ReceivedAt string  `json:"received_at"`
	ExpiresAt  string  `json:"expires_at"`
	// Your own opaque reference for the outcome, such as a case or ticket id. Stored as given and not interpreted; keep direct identifiers out of it.
	ExternalRef       *string             `json:"external_ref,omitzero"`
	Actor             AccountOutcomeActor `json:"actor"`
	RiskVersionBefore int64               `json:"risk_version_before"`
	RiskVersion       int64               `json:"risk_version"`
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	EffectiveOutcomeID *string `json:"effective_outcome_id,omitzero"`
}

func (a *AccountOutcome) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AccountOutcome) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountOutcome) GetOutcome() Outcome {
	if a == nil {
		return Outcome("")
	}
	return a.Outcome
}

func (a *AccountOutcome) GetSupersedes() *string {
	if a == nil {
		return nil
	}
	return a.Supersedes
}

func (a *AccountOutcome) GetObservedAt() string {
	if a == nil {
		return ""
	}
	return a.ObservedAt
}

func (a *AccountOutcome) GetReceivedAt() string {
	if a == nil {
		return ""
	}
	return a.ReceivedAt
}

func (a *AccountOutcome) GetExpiresAt() string {
	if a == nil {
		return ""
	}
	return a.ExpiresAt
}

func (a *AccountOutcome) GetExternalRef() *string {
	if a == nil {
		return nil
	}
	return a.ExternalRef
}

func (a *AccountOutcome) GetActor() AccountOutcomeActor {
	if a == nil {
		return AccountOutcomeActor{}
	}
	return a.Actor
}

func (a *AccountOutcome) GetRiskVersionBefore() int64 {
	if a == nil {
		return 0
	}
	return a.RiskVersionBefore
}

func (a *AccountOutcome) GetRiskVersion() int64 {
	if a == nil {
		return 0
	}
	return a.RiskVersion
}

func (a *AccountOutcome) GetEffectiveOutcomeID() *string {
	if a == nil {
		return nil
	}
	return a.EffectiveOutcomeID
}
