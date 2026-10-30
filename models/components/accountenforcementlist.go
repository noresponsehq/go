package components

// AccountEnforcementList - Page through account enforcements.
type AccountEnforcementList struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string        `json:"account_id"`
	Items     []Enforcement `json:"items"`
	// Opaque, expiring cursor bound to the workspace, environment, account and history kind. Absent at the end.
	NextCursor *string `json:"next_cursor,omitzero"`
}

func (a *AccountEnforcementList) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountEnforcementList) GetItems() []Enforcement {
	if a == nil {
		return []Enforcement{}
	}
	return a.Items
}

func (a *AccountEnforcementList) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}
