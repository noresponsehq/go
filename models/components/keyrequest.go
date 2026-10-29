package components

import (
	"encoding/json"
	"fmt"
)

type KeyRequestScope string

const (
	KeyRequestScopeAssess        KeyRequestScope = "assess"
	KeyRequestScopeHistory       KeyRequestScope = "history"
	KeyRequestScopeDelete        KeyRequestScope = "delete"
	KeyRequestScopeWebhooks      KeyRequestScope = "webhooks"
	KeyRequestScopeDetailed      KeyRequestScope = "detailed"
	KeyRequestScopeIntegrations  KeyRequestScope = "integrations"
	KeyRequestScopeManagement    KeyRequestScope = "management"
	KeyRequestScopeAudit         KeyRequestScope = "audit"
	KeyRequestScopeUsage         KeyRequestScope = "usage"
	KeyRequestScopeInvestigate   KeyRequestScope = "investigate"
	KeyRequestScopeEvidence      KeyRequestScope = "evidence"
	KeyRequestScopeExports       KeyRequestScope = "exports"
	KeyRequestScopeAuditManage   KeyRequestScope = "audit_manage"
	KeyRequestScopeSettings      KeyRequestScope = "settings"
	KeyRequestScopePolicyPublish KeyRequestScope = "policy_publish"
	KeyRequestScopeMembers       KeyRequestScope = "members"
	KeyRequestScopeMembersManage KeyRequestScope = "members_manage"
)

func (e KeyRequestScope) ToPointer() *KeyRequestScope {
	return &e
}
func (e *KeyRequestScope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "assess":
		fallthrough
	case "history":
		fallthrough
	case "delete":
		fallthrough
	case "webhooks":
		fallthrough
	case "detailed":
		fallthrough
	case "integrations":
		fallthrough
	case "management":
		fallthrough
	case "audit":
		fallthrough
	case "usage":
		fallthrough
	case "investigate":
		fallthrough
	case "evidence":
		fallthrough
	case "exports":
		fallthrough
	case "audit_manage":
		fallthrough
	case "settings":
		fallthrough
	case "policy_publish":
		fallthrough
	case "members":
		fallthrough
	case "members_manage":
		*e = KeyRequestScope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyRequestScope: %v", v)
	}
}

// KeyRequest - Creates a key whose scopes are bounded by the issuer capabilities in this environment. Delete is owner-only. A service key outlives its creator; a personal key belongs to the authenticated member and is revoked on membership or authority changes. request_id deduplicates creation: retry returns 200 metadata without the secret; changed parameters answer 409. Team scopes members and members_manage are explicitly workspace-wide, live-only, and delegated by a current full owner. Personal credentials cannot hold management or members_manage. Legacy management includes policy publication; use settings without policy_publish for configuration preparation and retention-only edits.
type KeyRequest struct {
	// The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.
	Environment KeyEnvironment `json:"environment"`
	// The key's scopes, each at most once.
	Scopes []KeyRequestScope `json:"scopes"`
	// Optional display name for this API key, at most 100 printable characters. Preserved on rotation.
	Name *string `json:"name,omitzero"`
	// Optional operational purpose, at most 200 printable characters without leading or trailing whitespace. Do not put secrets or personal data here. Preserved on rotation and included in the creation request fingerprint; omitted for older credentials.
	Purpose *string  `json:"purpose,omitzero"`
	Kind    *KeyKind `json:"kind,omitzero"`
	// Between one minute and 365 days from now.
	ExpiresAt *string `json:"expires_at,omitzero"`
	// Caller-chosen identifier for one creation attempt, scoped to the issuing actor and environment. Use 1 to 256 printable ASCII characters without spaces. A retry with the same parameters returns the original key metadata with status 200 and no secret; changed parameters return 409.
	RequestID *string `json:"request_id,omitzero"`
}

func (k *KeyRequest) GetEnvironment() KeyEnvironment {
	if k == nil {
		return KeyEnvironment("")
	}
	return k.Environment
}

func (k *KeyRequest) GetScopes() []KeyRequestScope {
	if k == nil {
		return []KeyRequestScope{}
	}
	return k.Scopes
}

func (k *KeyRequest) GetName() *string {
	if k == nil {
		return nil
	}
	return k.Name
}

func (k *KeyRequest) GetPurpose() *string {
	if k == nil {
		return nil
	}
	return k.Purpose
}

func (k *KeyRequest) GetKind() *KeyKind {
	if k == nil {
		return nil
	}
	return k.Kind
}

func (k *KeyRequest) GetExpiresAt() *string {
	if k == nil {
		return nil
	}
	return k.ExpiresAt
}

func (k *KeyRequest) GetRequestID() *string {
	if k == nil {
		return nil
	}
	return k.RequestID
}
