package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type CreateEventRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.EventRequest   `request:"mediaType=application/json"`
}

func (c *CreateEventRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateEventRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CreateEventRequest) GetBody() components.EventRequest {
	if c == nil {
		return components.EventRequest{}
	}
	return c.Body
}

type CreateEventResponse struct {
	Headers map[string][]string
	Result  components.Event
}

func (c *CreateEventResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateEventResponse) GetResult() components.Event {
	if c == nil {
		return components.Event{}
	}
	return c.Result
}
