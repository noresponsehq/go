package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

// AssessmentQuery - POST /v1/assessments/query body. account_id travels in the body rather than the URL path so a provider value that turns out to hold personal data is not written to access logs, proxies or browser history. Read-only: no Idempotency-Key. An account with no unexpired assessments answers 200 with empty items.
type AssessmentQuery struct {
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string `json:"account_id"`
	Limit     *int64 `default:"20" json:"limit"`
	// next_cursor from the previous page, unchanged. Opaque, MAC'd and bound to the workspace and account_id.
	Cursor *string `json:"cursor,omitzero"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode *ResponseMode `json:"response_mode,omitzero"`
}

func (a AssessmentQuery) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentQuery) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentQuery) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AssessmentQuery) GetLimit() *int64 {
	if a == nil {
		return nil
	}
	return a.Limit
}

func (a *AssessmentQuery) GetCursor() *string {
	if a == nil {
		return nil
	}
	return a.Cursor
}

func (a *AssessmentQuery) GetResponseMode() *ResponseMode {
	if a == nil {
		return nil
	}
	return a.ResponseMode
}
