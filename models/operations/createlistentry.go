package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CreateListEntryRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment   `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.ListEntryRequest `request:"mediaType=application/json"`
}

func (c *CreateListEntryRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateListEntryRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CreateListEntryRequest) GetBody() components.ListEntryRequest {
	if c == nil {
		return components.ListEntryRequest{}
	}
	return c.Body
}

type CreateListEntryResponse struct {
	Headers map[string][]string
	Result  components.ListEntry
}

func (c *CreateListEntryResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateListEntryResponse) GetResult() components.ListEntry {
	if c == nil {
		return components.ListEntry{}
	}
	return c.Result
}
