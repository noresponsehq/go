package components

type KeyScope string

const (
	KeyScopeAssess        KeyScope = "assess"
	KeyScopeHistory       KeyScope = "history"
	KeyScopeDelete        KeyScope = "delete"
	KeyScopeWebhooks      KeyScope = "webhooks"
	KeyScopeDetailed      KeyScope = "detailed"
	KeyScopeIntegrations  KeyScope = "integrations"
	KeyScopeManagement    KeyScope = "management"
	KeyScopeAudit         KeyScope = "audit"
	KeyScopeUsage         KeyScope = "usage"
	KeyScopeInvestigate   KeyScope = "investigate"
	KeyScopeEvidence      KeyScope = "evidence"
	KeyScopeExports       KeyScope = "exports"
	KeyScopeAuditManage   KeyScope = "audit_manage"
	KeyScopeSettings      KeyScope = "settings"
	KeyScopePolicyPublish KeyScope = "policy_publish"
	KeyScopeMembers       KeyScope = "members"
	KeyScopeMembersManage KeyScope = "members_manage"
)

func (e KeyScope) ToPointer() *KeyScope {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *KeyScope) IsExact() bool {
	if e != nil {
		switch *e {
		case "assess", "history", "delete", "webhooks", "detailed", "integrations", "management", "audit", "usage", "investigate", "evidence", "exports", "audit_manage", "settings", "policy_publish", "members", "members_manage":
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
	// The earlier of the explicit lifetime and any rotation overlap. Rotation preserves the original lifetime.
	ExpiresAt *string `json:"expires_at,omitzero"`
	RevokedAt *string `json:"revoked_at,omitzero"`
	// The external user id of the member who revoked the key through the console, or the prefix of the management key that revoked it. Absent when the key is not revoked or NoResponse revoked it.
	RevokedBy *string `json:"revoked_by,omitzero"`
	// The credential, nrk_<environment>_<key id>_<secret>. Present only in the 201 answer that created the key.
	Key *string `json:"key,omitzero"`
	// The optional display name supplied for this API key. Preserved on rotation.
	Name *string `json:"name,omitzero"`
	// Optional operational purpose, at most 200 printable characters without leading or trailing whitespace. Do not put secrets or personal data here. Preserved on rotation and included in the creation request fingerprint; omitted for older credentials.
	Purpose *string  `json:"purpose,omitzero"`
	Kind    *KeyKind `json:"kind,omitzero"`
	// The issuing member's external identity for a personal key. Personal ownership is fixed and the key is revoked when that member loses authority. Service keys belong to the workspace and omit this field.
	OwnerUserID *string `json:"owner_user_id,omitzero"`
	// Replacement public key identifier, or legacy when the successor predates recorded lineage. This field excludes the secret.
	RotatedTo *string `json:"rotated_to,omitzero"`
	// Latest recorded authenticated request observation, updated at most once per minute. Capture is best effort after the response and includes authenticated requests refused by authorization or rate limits. Absence means unknown and does not prove the key was unused. It is independent of diagnostic retention and resets on the replacement key after rotation. This field retains no raw IP or user-agent.
	LastUsedAt *string `json:"last_used_at,omitzero"`
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

func (k *Key) GetName() *string {
	if k == nil {
		return nil
	}
	return k.Name
}

func (k *Key) GetPurpose() *string {
	if k == nil {
		return nil
	}
	return k.Purpose
}

func (k *Key) GetKind() *KeyKind {
	if k == nil {
		return nil
	}
	return k.Kind
}

func (k *Key) GetOwnerUserID() *string {
	if k == nil {
		return nil
	}
	return k.OwnerUserID
}

func (k *Key) GetRotatedTo() *string {
	if k == nil {
		return nil
	}
	return k.RotatedTo
}

func (k *Key) GetLastUsedAt() *string {
	if k == nil {
		return nil
	}
	return k.LastUsedAt
}
