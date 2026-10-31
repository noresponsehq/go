package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type UserActor struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"user" json:"type"`
	// The user's id at the identity provider the console signs people in with: 1 to 256 printable ASCII characters with no space, compared exactly.
	ExternalUserID string `json:"external_user_id"`
	// The public id of the console credential the user acted through.
	ConsoleCredentialID string `json:"console_credential_id"`
	// Stable workspace principal id, retained across identity-provider subject migration. Absent for older events.
	PrincipalID *string `json:"principal_id,omitzero"`
}

func (u UserActor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserActor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (u *UserActor) GetType() string {
	return "user"
}

func (u *UserActor) GetExternalUserID() string {
	if u == nil {
		return ""
	}
	return u.ExternalUserID
}

func (u *UserActor) GetConsoleCredentialID() string {
	if u == nil {
		return ""
	}
	return u.ConsoleCredentialID
}

func (u *UserActor) GetPrincipalID() *string {
	if u == nil {
		return nil
	}
	return u.PrincipalID
}
