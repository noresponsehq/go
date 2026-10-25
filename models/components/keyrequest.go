package components

import (
	"encoding/json"
	"fmt"
)

type KeyRequestScope string

const (
	KeyRequestScopeAssess       KeyRequestScope = "assess"
	KeyRequestScopeHistory      KeyRequestScope = "history"
	KeyRequestScopeDelete       KeyRequestScope = "delete"
	KeyRequestScopeWebhooks     KeyRequestScope = "webhooks"
	KeyRequestScopeDetailed     KeyRequestScope = "detailed"
	KeyRequestScopeIntegrations KeyRequestScope = "integrations"
	KeyRequestScopeManagement   KeyRequestScope = "management"
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
		*e = KeyRequestScope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyRequestScope: %v", v)
	}
}

// KeyRequest - POST /v1/workspaces/{workspace_id}/keys body. Creates an API key in one environment of the workspace with the scopes named, and returns its secret once. A live key needs the admin or owner role; a test key needs developer or above. An API key with the management scope creates keys in its own environment with scopes it holds itself. Each scope may appear once.
type KeyRequest struct {
	// The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.
	Environment KeyEnvironment `json:"environment"`
	// The key's scopes, each at most once.
	Scopes []KeyRequestScope `json:"scopes"`
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
