package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

type CredentialAssociationRetiredBy struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"api_key" json:"type"`
	KeyID string `json:"key_id"`
}

func (c CredentialAssociationRetiredBy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CredentialAssociationRetiredBy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CredentialAssociationRetiredBy) GetType() string {
	return "api_key"
}

func (c *CredentialAssociationRetiredBy) GetKeyID() string {
	if c == nil {
		return ""
	}
	return c.KeyID
}
