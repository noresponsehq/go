package operations

import (
	"github.com/noresponsehq/go/v15/models/components"
)

type UpdateKeyAccessRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The key's id. A key of another workspace answers 404, as does a key of the other environment to an API key.
	KeyID string                      `pathParam:"style=simple,explode=false,name=key_id"`
	Body  components.KeyAccessRequest `request:"mediaType=application/json"`
}

func (u *UpdateKeyAccessRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateKeyAccessRequest) GetKeyID() string {
	if u == nil {
		return ""
	}
	return u.KeyID
}

func (u *UpdateKeyAccessRequest) GetBody() components.KeyAccessRequest {
	if u == nil {
		return components.KeyAccessRequest{}
	}
	return u.Body
}

type UpdateKeyAccessResponse struct {
	Headers map[string][]string
	Result  components.Key
}

func (u *UpdateKeyAccessResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateKeyAccessResponse) GetResult() components.Key {
	if u == nil {
		return components.Key{}
	}
	return u.Result
}
