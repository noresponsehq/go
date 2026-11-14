package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type CreateGroupRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.GroupRequest   `request:"mediaType=application/json"`
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
