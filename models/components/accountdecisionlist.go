package components

// AccountDecisionList - Page through recorded account decisions.
type AccountDecisionList struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string                    `json:"account_id"`
	Items     []AccountDecisionSnapshot `json:"items"`
	// Opaque, expiring cursor bound to the workspace, environment, account and history kind. Absent at the end.
	NextCursor *string `json:"next_cursor,omitzero"`
}

func (a *AccountDecisionList) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AccountDecisionList) GetItems() []AccountDecisionSnapshot {
	if a == nil {
		return []AccountDecisionSnapshot{}
	}
	return a.Items
}

func (a *AccountDecisionList) GetNextCursor() *string {
	if a == nil {
		return nil
	}
	return a.NextCursor
}
