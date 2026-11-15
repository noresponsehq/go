package components

import (
	"encoding/json"
	"fmt"
)

type KeyAccessRequestDeniedAction string

const (
	KeyAccessRequestDeniedActionEvidenceRead  KeyAccessRequestDeniedAction = "evidence.read"
	KeyAccessRequestDeniedActionExportsRead   KeyAccessRequestDeniedAction = "exports.read"
	KeyAccessRequestDeniedActionPolicyPublish KeyAccessRequestDeniedAction = "policy.publish"
)

func (e KeyAccessRequestDeniedAction) ToPointer() *KeyAccessRequestDeniedAction {
	return &e
}
func (e *KeyAccessRequestDeniedAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "evidence.read":
		fallthrough
	case "exports.read":
		fallthrough
	case "policy.publish":
		*e = KeyAccessRequestDeniedAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyAccessRequestDeniedAction: %v", v)
	}
}

type KeyAccessRequest struct {
	// Explicit restrictions that override scope grants. An empty list restores the access supplied by the scopes. Rotation preserves these restrictions.
	DeniedActions []KeyAccessRequestDeniedAction `json:"denied_actions"`
}

func (k *KeyAccessRequest) GetDeniedActions() []KeyAccessRequestDeniedAction {
	if k == nil {
		return []KeyAccessRequestDeniedAction{}
	}
	return k.DeniedActions
}
