package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type KeyActor struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"api_key" json:"type"`
	// The public id of the workspace API key that made the change.
	KeyID string `json:"key_id"`
}

func (k KeyActor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyActor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (k *KeyActor) GetType() string {
	return "api_key"
}

func (k *KeyActor) GetKeyID() string {
	if k == nil {
		return ""
	}
	return k.KeyID
}
