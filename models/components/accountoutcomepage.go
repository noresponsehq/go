package components

// AccountOutcomePage - List retained account reports, newest first. The page-level effective outcome id names the current report. Each receipt preserves the effective report and risk version from its original write. Page with the returned cursor.
type AccountOutcomePage struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string           `json:"account_id"`
	Items     []AccountOutcome `json:"items"`
	// Opaque, expiring cursor bound to this environment, operation and every filter. Keep the filters unchanged.
	NextCursor *string `json:"next_cursor,omitzero"`
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	EffectiveOutcomeID *string `json:"effective_outcome_id,omitzero"`
}

func (a *AccountOutcomePage) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountOutcomePage) GetItems() []AccountOutcome {
	if a == nil {
		return []AccountOutcome{}
	}
	return a.Items
}

func (a *AccountOutcomePage) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}

func (a *AccountOutcomePage) GetEffectiveOutcomeID() *string {
	if a == nil {
		return nil
	}
	return a.EffectiveOutcomeID
}
