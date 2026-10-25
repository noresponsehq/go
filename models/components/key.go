package components

type KeyScope string

const (
	KeyScopeAssess       KeyScope = "assess"
	KeyScopeHistory      KeyScope = "history"
	KeyScopeDelete       KeyScope = "delete"
	KeyScopeWebhooks     KeyScope = "webhooks"
	KeyScopeDetailed     KeyScope = "detailed"
	KeyScopeIntegrations KeyScope = "integrations"
	KeyScopeManagement   KeyScope = "management"
)

func (e KeyScope) ToPointer() *KeyScope {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *KeyScope) IsExact() bool {
	if e != nil {
		switch *e {
		case "assess", "history", "delete", "webhooks", "detailed", "integrations", "management":
			return true
		}
	}
	return false
}

// Key - One API key of your workspace. A key belongs to exactly one environment. key, the secret credential, is present only in the answer that created the key: NoResponse keeps a peppered digest of it, so store it now. Every other field is safe to show.
type Key struct {
	// The key's public id, a 26-character ULID.
	ID string `json:"id"`
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string `json:"workspace_id"`
	// The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.
	Environment KeyEnvironment `json:"environment"`
	// The public start of the credential, nrk_<environment>_<key id>, which identifies the key in logs and secret stores without its secret.
	Prefix string `json:"prefix"`
	// The key's scopes, each at most once.
	Scopes []KeyScope `json:"scopes"`
	// active: the key authenticates. expired: a rotation's overlap window has ended (expires_at). revoked: the key was revoked (revoked_at) and answers 401.
	Status    KeyStatus `json:"status"`
	CreatedAt string    `json:"created_at"`
	// The external user id of the member who created the key through the console, or the prefix (nrk_<environment>_<key id>) of the management key that created it. Absent for a key NoResponse issued or a rotation made.
	CreatedBy *string `json:"created_by,omitzero"`
	// When a rotated key stops authenticating (POST /v1/keys/rotate). Absent for a key with no end.
	ExpiresAt *string `json:"expires_at,omitzero"`
	RevokedAt *string `json:"revoked_at,omitzero"`
	// The external user id of the member who revoked the key through the console, or the prefix of the management key that revoked it. Absent when the key is not revoked or NoResponse revoked it.
	RevokedBy *string `json:"revoked_by,omitzero"`
	// The credential, nrk_<environment>_<key id>_<secret>. Present only in the 201 answer that created the key.
	Key *string `json:"key,omitzero"`
}

func (k *Key) GetID() string {
	if k == nil {
		return ""
	}
	return k.ID
}

func (k *Key) GetWorkspaceID() string {
	if k == nil {
		return ""
	}
	return k.WorkspaceID
}

func (k *Key) GetEnvironment() KeyEnvironment {
	if k == nil {
		return KeyEnvironment("")
	}
	return k.Environment
}

func (k *Key) GetPrefix() string {
	if k == nil {
		return ""
	}
	return k.Prefix
}

func (k *Key) GetScopes() []KeyScope {
	if k == nil {
		return []KeyScope{}
	}
	return k.Scopes
}

func (k *Key) GetStatus() KeyStatus {
	if k == nil {
		return KeyStatus("")
	}
	return k.Status
}

func (k *Key) GetCreatedAt() string {
	if k == nil {
		return ""
	}
	return k.CreatedAt
}

func (k *Key) GetCreatedBy() *string {
	if k == nil {
		return nil
	}
	return k.CreatedBy
}

func (k *Key) GetExpiresAt() *string {
	if k == nil {
		return nil
	}
	return k.ExpiresAt
}

func (k *Key) GetRevokedAt() *string {
	if k == nil {
		return nil
	}
	return k.RevokedAt
}

func (k *Key) GetRevokedBy() *string {
	if k == nil {
		return nil
	}
	return k.RevokedBy
}

func (k *Key) GetKey() *string {
	if k == nil {
		return nil
	}
	return k.Key
}
