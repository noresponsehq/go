package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type CancelScreeningRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment         `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                            `pathParam:"style=simple,explode=false,name=group_id"`
	ScreeningID string                            `pathParam:"style=simple,explode=false,name=screening_id"`
	Body        components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (c *CancelScreeningRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CancelScreeningRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CancelScreeningRequest) GetGroupID() string {
	if c == nil {
		return ""
	}
	return c.GroupID
}

func (c *CancelScreeningRequest) GetScreeningID() string {
	if c == nil {
		return ""
	}
	return c.ScreeningID
}

func (c *CancelScreeningRequest) GetBody() components.ScreeningCancelRequest {
	if c == nil {
		return components.ScreeningCancelRequest{}
	}
	return c.Body
}

type CancelScreeningResponse struct {
	Headers map[string][]string
	Result  components.Screening
}

func (c *CancelScreeningResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CancelScreeningResponse) GetResult() components.Screening {
	if c == nil {
		return components.Screening{}
	}
	return c.Result
}
