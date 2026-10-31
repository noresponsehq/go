package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ConfigureAuditStreamRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                        `pathParam:"style=simple,explode=false,name=workspace_id"`
	Environment components.Environment        `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.AuditStreamRequest `request:"mediaType=application/json"`
}

func (c *ConfigureAuditStreamRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *ConfigureAuditStreamRequest) GetEnvironment() components.Environment {
	if c == nil {
		return components.Environment("")
	}
	return c.Environment
}

func (c *ConfigureAuditStreamRequest) GetBody() components.AuditStreamRequest {
	if c == nil {
		return components.AuditStreamRequest{}
	}
	return c.Body
}

type ConfigureAuditStreamResponse struct {
	Headers map[string][]string
	Result  components.AuditStreamList
}

func (c *ConfigureAuditStreamResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *ConfigureAuditStreamResponse) GetResult() components.AuditStreamList {
	if c == nil {
		return components.AuditStreamList{}
	}
	return c.Result
}
