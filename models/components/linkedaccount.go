package components

type LinkedAccount struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	// One entry per shared kind, in the order email, phone, username. Each entry uses the strictest basis that links the two accounts: exact, then canonical, then mailbox.
	SharedIdentifiers []SharedIdentifier `json:"shared_identifiers"`
}

func (l *LinkedAccount) GetAccountID() string {
	if l == nil {
		return ""
	}
	return l.AccountID
}

func (l *LinkedAccount) GetSharedIdentifiers() []SharedIdentifier {
	if l == nil {
		return []SharedIdentifier{}
	}
	return l.SharedIdentifiers
}
