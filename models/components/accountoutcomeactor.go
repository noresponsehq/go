package components

type AccountOutcomeActor struct {
	Type  AccountOutcomeActorType `json:"type"`
	KeyID *string                 `json:"key_id,omitzero"`
	// The user's id at the identity provider the console signs people in with: 1 to 256 printable ASCII characters with no space, compared exactly.
	ExternalUserID      *string `json:"external_user_id,omitzero"`
	ConsoleCredentialID *string `json:"console_credential_id,omitzero"`
	// Opaque, server-assigned member id: mem_ and a 26-character ULID. It stays the same when a removed member is invited back.
	PrincipalID *string `json:"principal_id,omitzero"`
}

func (a *AccountOutcomeActor) GetType() AccountOutcomeActorType {
	if a == nil {
		return AccountOutcomeActorType("")
	}
	return a.Type
}

func (a *AccountOutcomeActor) GetKeyID() *string {
	if a == nil {
		return nil
	}
	return a.KeyID
}

func (a *AccountOutcomeActor) GetExternalUserID() *string {
	if a == nil {
		return nil
	}
	return a.ExternalUserID
}

func (a *AccountOutcomeActor) GetConsoleCredentialID() *string {
	if a == nil {
		return nil
	}
	return a.ConsoleCredentialID
}

func (a *AccountOutcomeActor) GetPrincipalID() *string {
	if a == nil {
		return nil
	}
	return a.PrincipalID
}
