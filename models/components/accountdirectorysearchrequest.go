package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AccountDirectorySearchRequest - Bounded distinct-account search for one authenticated workspace/environment. identifier matches any retained admitted assessment of an account, including historical submitted values. email preserves the legacy canonical-mailbox filter. identifier and email cannot be combined. All other filters combine using AND. Typed identifier search uses stable opaque account order; ordinary and legacy email searches retain newest-seen ordering. Pages are live, not a cross-page snapshot. Requires history or investigate authority and identifier-read permission for identifier, identifier_expression or email filtering. identifier_expression explicitly chooses scope assessment (one whole matching submission) or account (terms may match different retained submissions of one account). At most 8 raw terms and 3 group levels. Mutually exclusive with identifier and email. Expression search uses the same stable opaque account order as typed identifier search.
type AccountDirectorySearchRequest struct {
	// Select accounts with a retained open or in-progress review, or a current review/block account decision with no retained effective account outcome. The server applies this filter before counting and pagination. Requires reviews.read as well as history.read for console members; API keys additionally require management.
	NeedsAction *bool `default:"false" json:"needs_action"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID *string `json:"account_id,omitzero"`
	// Assessments of every account that one of its retained admitted assessments sent this mailbox as identifiers.email, matched as lists match an email entry: case-insensitive, a +tag removed and, at gmail.com and googlemail.com, the dots of the local part ignored. Only assessments stored since email list matching began carry the mailbox. An address that is not a valid mailbox is refused with 400.
	Email *string `json:"email,omitzero"`
	Limit *int64  `default:"25" json:"limit"`
	// Opaque, expiring cursor bound to workspace, environment, operation and all filters, including expression scope, ordered structure and original term values. Keep grouping, operand order, matching modes and values unchanged; an omitted match is equivalent to exact. The page limit may change.
	Cursor *string `json:"cursor,omitzero"`
	// One customer-submitted identifier. Defaults to exact comparison of the original submitted value. canonical uses the supported identifier normalisation without an ownership claim: email preserves local-part case, NFC-normalises it and lowercases the domain without removing dots or tags; phone uses international E.164 without extensions or an inferred region; username uses NFC and Unicode case folding within the customer namespace. mailbox is email-only and explicitly selects existing list rules: case-insensitive, remove +tag, and ignore local-part dots at gmail.com and googlemail.com; googlemail.com maps to gmail.com. Username requires namespace customer. Email and phone forbid namespace. Source-discovered identifiers and platform handles are not searched. Unsupported kinds, namespaces, modes and invalid values return 400. Search equality does not establish verified ownership or merge accounts.
	Identifier *IdentifierSearch `json:"identifier,omitzero"`
	// A bounded AND/OR expression over submitted identifiers. Required scope assessment means the entire expression must match one retained assessment. Scope account permits each term to match a different retained assessment of the same account and is accepted only by account search. Exactly one of term, and or or is allowed at every node. Groups contain 2 to 8 operands, with at most 3 nested group levels and 8 raw term occurrences over the entire expression. A root term is allowed. Duplicate terms with the same kind, match rule, namespace and normalized comparison are idempotent, but count toward the raw limit. Match rules default to exact and preserve the existing identifier contract. Only retained submitted identifiers can satisfy a term. Unknown fields, unsupported scopes and malformed or oversized expressions return 400 invalid_request.
	IdentifierExpression *IdentifierExpression `json:"identifier_expression,omitzero"`
}

func (a AccountDirectorySearchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountDirectorySearchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountDirectorySearchRequest) GetNeedsAction() *bool {
	if a == nil {
		return nil
	}
	return a.NeedsAction
}

func (a *AccountDirectorySearchRequest) GetAccountID() *string {
	if a == nil {
		return nil
	}
	return a.AccountID
}

func (a *AccountDirectorySearchRequest) GetEmail() *string {
	if a == nil {
		return nil
	}
	return a.Email
}

func (a *AccountDirectorySearchRequest) GetLimit() *int64 {
	if a == nil {
		return nil
	}
	return a.Limit
}

func (a *AccountDirectorySearchRequest) GetCursor() *string {
	if a == nil {
		return nil
	}
	return a.Cursor
}

func (a *AccountDirectorySearchRequest) GetIdentifier() *IdentifierSearch {
	if a == nil {
		return nil
	}
	return a.Identifier
}

func (a *AccountDirectorySearchRequest) GetIdentifierExpression() *IdentifierExpression {
	if a == nil {
		return nil
	}
	return a.IdentifierExpression
}
