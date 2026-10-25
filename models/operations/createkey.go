package operations

import (
	"github.com/noresponsehq/go/v7/models/components"
)

type CreateKeyRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                `pathParam:"style=simple,explode=false,name=workspace_id"`
	Body        components.KeyRequest `request:"mediaType=application/json"`
}

func (c *CreateKeyRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateKeyRequest) GetBody() components.KeyRequest {
	if c == nil {
		return components.KeyRequest{}
	}
	return c.Body
}

type CreateKeyResponse struct {
	Headers map[string][]string
	Result  components.Key
}

func (c *CreateKeyResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateKeyResponse) GetResult() components.Key {
	if c == nil {
		return components.Key{}
	}
	return c.Result
}
