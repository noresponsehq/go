package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// CredentialAssociation - An association between an account and an issued key. The response excludes the key and its hash. Retirement removes this association from risk calculations; revoke the actual key through your own service.
type CredentialAssociation struct {
	ID string `json:"id"`
	// Your own opaque identifier for the account being checked, scoped to your workspace. Use an id you mint, with no direct identifier (email, phone, name or national id) in it. '|' is accepted so an Auth0 user id such as 'auth0|64abc' passes unchanged; '@', '+', '$' and whitespace are refused by the pattern, as is an exact-shape JWT; a hex digest such as a SHA-256 is allowed, because hashing an internal id is a common pseudonymous scheme. A value such as '07700900123' or 'jane.doe.gmail.com' cannot be told apart from an opaque id. It travels in bodies, and in a URL only as the path segment of GET /v1/accounts/{account_id}, its /assessments and its /linked_accounts. An assessment requested by a provider hook at pre_registration, which has no account yet, carries the server-issued account_id pre: followed by its assessment_id.
	AccountID string                          `json:"account_id"`
	CreatedAt string                          `json:"created_at"`
	RetiredAt *string                         `json:"retired_at,omitzero"`
	Actor     CredentialAssociationActor      `json:"actor"`
	RetiredBy *CredentialAssociationRetiredBy `json:"retired_by,omitzero"`
}

func (c CredentialAssociation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CredentialAssociation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CredentialAssociation) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *CredentialAssociation) GetAccountID() string {
	if c == nil {
		return ""
	}
	return c.AccountID
}

func (c *CredentialAssociation) GetCreatedAt() string {
	if c == nil {
		return ""
	}
	return c.CreatedAt
}

func (c *CredentialAssociation) GetRetiredAt() *string {
	if c == nil {
		return nil
	}
	return c.RetiredAt
}

func (c *CredentialAssociation) GetActor() CredentialAssociationActor {
	if c == nil {
		return CredentialAssociationActor{}
	}
	return c.Actor
}

func (c *CredentialAssociation) GetRetiredBy() *CredentialAssociationRetiredBy {
	if c == nil {
		return nil
	}
	return c.RetiredBy
}
