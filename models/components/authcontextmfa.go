package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AuthContextMfa - Normalized multi-factor evidence from a verified provider session.
type AuthContextMfa struct {
	Present    *bool    `json:"present,omitzero"`
	Methods    []string `json:"methods,omitzero"`
	AgeSeconds *int64   `json:"age_seconds,omitzero"`
}

func (a AuthContextMfa) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthContextMfa) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuthContextMfa) GetPresent() *bool {
	if a == nil {
		return nil
	}
	return a.Present
}

func (a *AuthContextMfa) GetMethods() []string {
	if a == nil {
		return nil
	}
	return a.Methods
}

func (a *AuthContextMfa) GetAgeSeconds() *int64 {
	if a == nil {
		return nil
	}
	return a.AgeSeconds
}
