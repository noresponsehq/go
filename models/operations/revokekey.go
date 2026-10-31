package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type RevokeKeyRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The key's id. A key of another workspace answers 404, as does a key of the other environment to an API key.
	KeyID string `pathParam:"style=simple,explode=false,name=key_id"`
}

func (r *RevokeKeyRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RevokeKeyRequest) GetKeyID() string {
	if r == nil {
		return ""
	}
	return r.KeyID
}

type RevokeKeyResponse struct {
	Headers map[string][]string
	Result  components.Key
}

func (r *RevokeKeyResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RevokeKeyResponse) GetResult() components.Key {
	if r == nil {
		return components.Key{}
	}
	return r.Result
}
