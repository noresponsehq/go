package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type AccountBindingActor struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"api_key" json:"type"`
	KeyID string `json:"key_id"`
}

func (a AccountBindingActor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountBindingActor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountBindingActor) GetType() string {
	return "api_key"
}

func (a *AccountBindingActor) GetKeyID() string {
	if a == nil {
		return ""
	}
	return a.KeyID
}
