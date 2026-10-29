package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type CredentialAssociationActor struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"api_key" json:"type"`
	KeyID string `json:"key_id"`
}

func (c CredentialAssociationActor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CredentialAssociationActor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CredentialAssociationActor) GetType() string {
	return "api_key"
}

func (c *CredentialAssociationActor) GetKeyID() string {
	if c == nil {
		return ""
	}
	return c.KeyID
}
