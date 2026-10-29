package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type CustomerExportRequest struct {
	Kind        CustomerExportRequestKind `json:"kind"`
	Environment Environment               `json:"environment"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id} and its /assessments. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID *string `json:"account_id,omitzero"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode *ResponseMode                `json:"response_mode,omitzero"`
	Filter       *CustomerExportRequestFilter `json:"filter,omitzero"`
}

func (c CustomerExportRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomerExportRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CustomerExportRequest) GetKind() CustomerExportRequestKind {
	if c == nil {
		return CustomerExportRequestKind("")
	}
	return c.Kind
}

func (c *CustomerExportRequest) GetEnvironment() Environment {
	if c == nil {
		return Environment("")
	}
	return c.Environment
}

func (c *CustomerExportRequest) GetAccountID() *string {
	if c == nil {
		return nil
	}
	return c.AccountID
}

func (c *CustomerExportRequest) GetResponseMode() *ResponseMode {
	if c == nil {
		return nil
	}
	return c.ResponseMode
}

func (c *CustomerExportRequest) GetFilter() *CustomerExportRequestFilter {
	if c == nil {
		return nil
	}
	return c.Filter
}
