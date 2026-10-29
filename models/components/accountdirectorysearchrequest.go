package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AccountDirectorySearchRequest - Bounded distinct-account search for a single authenticated workspace/environment. Email matches any retained admitted assessment of an account, using assessment search's canonical mailbox rules. No assessment history download is required.
type AccountDirectorySearchRequest struct {
	// Select accounts with a retained open or in-progress review, or a current review/block account decision with no retained effective account outcome. The server applies this filter before counting and pagination. Requires reviews.read as well as history.read for console members; API keys additionally require management.
	NeedsAction *bool `default:"false" json:"needs_action"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID *string `json:"account_id,omitzero"`
	// Assessments of every account that one of its retained admitted assessments sent this mailbox as identifiers.email, matched as lists match an email entry: case-insensitive, a +tag removed and, at gmail.com and googlemail.com, the dots of the local part ignored. Only assessments stored since email list matching began carry the mailbox. An address that is not a valid mailbox is refused with 400.
	Email *string `json:"email,omitzero"`
	Limit *int64  `default:"25" json:"limit"`
	// Opaque, expiring cursor bound to this environment, operation and every filter. Keep the filters unchanged.
	Cursor *string `json:"cursor,omitzero"`
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
