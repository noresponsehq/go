package components

type Scope string

const (
	ScopeAssess       Scope = "assess"
	ScopeHistory      Scope = "history"
	ScopeWebhooks     Scope = "webhooks"
	ScopeDetailed     Scope = "detailed"
	ScopeIntegrations Scope = "integrations"
)

func (e Scope) ToPointer() *Scope {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Scope) IsExact() bool {
	if e != nil {
		switch *e {
		case "assess", "history", "webhooks", "detailed", "integrations":
			return true
		}
	}
	return false
}

// KeyRotation - The 201 answer of POST /v1/keys/rotate: the new key, shown once, and when the key that was presented stops working. The new key has the scopes of the presented key.
type KeyRotation struct {
	// The new credential, nrk_<key id>_<secret>. It is shown only in this answer; store it now.
	Key string `json:"key"`
	// The public id of the new key.
	KeyID string `json:"key_id"`
	// The scopes of the new key: those of the presented key.
	Scopes    []Scope `json:"scopes"`
	CreatedAt string  `json:"created_at"`
	// The id of the key that signed this request.
	PreviousKeyID string `json:"previous_key_id"`
	// When the previous key stops authenticating. Until then both keys work.
	PreviousKeyExpiresAt string `json:"previous_key_expires_at"`
}

func (k *KeyRotation) GetKey() string {
	if k == nil {
		return ""
	}
	return k.Key
}

func (k *KeyRotation) GetKeyID() string {
	if k == nil {
		return ""
	}
	return k.KeyID
}

func (k *KeyRotation) GetScopes() []Scope {
	if k == nil {
		return []Scope{}
	}
	return k.Scopes
}

func (k *KeyRotation) GetCreatedAt() string {
	if k == nil {
		return ""
	}
	return k.CreatedAt
}

func (k *KeyRotation) GetPreviousKeyID() string {
	if k == nil {
		return ""
	}
	return k.PreviousKeyID
}

func (k *KeyRotation) GetPreviousKeyExpiresAt() string {
	if k == nil {
		return ""
	}
	return k.PreviousKeyExpiresAt
}
