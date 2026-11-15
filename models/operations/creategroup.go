package operations

import (
	"github.com/noresponsehq/go/v13/models/components"
)

type CreateGroupRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	// Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person.
	IdempotencyKey *string                 `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.GroupRequest `request:"mediaType=application/json"`
}

func (c *CreateGroupRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateGroupRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CreateGroupRequest) GetIdempotencyKey() *string {
	if c == nil {
		return nil
	}
	return c.IdempotencyKey
}

func (c *CreateGroupRequest) GetBody() components.GroupRequest {
	if c == nil {
		return components.GroupRequest{}
	}
	return c.Body
}

type CreateGroupResponse struct {
	Headers map[string][]string
	Result  components.Group
}

func (c *CreateGroupResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateGroupResponse) GetResult() components.Group {
	if c == nil {
		return components.Group{}
	}
	return c.Result
}
