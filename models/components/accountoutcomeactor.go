package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

type AccountOutcomeActor struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"api_key" json:"type"`
	KeyID string `json:"key_id"`
}

func (a AccountOutcomeActor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountOutcomeActor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountOutcomeActor) GetType() string {
	return "api_key"
}

func (a *AccountOutcomeActor) GetKeyID() string {
	if a == nil {
		return ""
	}
	return a.KeyID
}
