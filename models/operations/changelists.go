package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ChangeListsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                     `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.ListBulkRequest `request:"mediaType=application/json"`
}

func (c *ChangeListsRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *ChangeListsRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *ChangeListsRequest) GetIdempotencyKey() string {
	if c == nil {
		return ""
	}
	return c.IdempotencyKey
}

func (c *ChangeListsRequest) GetBody() components.ListBulkRequest {
	if c == nil {
		return components.ListBulkRequest{}
	}
	return c.Body
}

type ChangeListsResponse struct {
	Headers map[string][]string
	Result  components.ListChange
}

func (c *ChangeListsResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *ChangeListsResponse) GetResult() components.ListChange {
	if c == nil {
		return components.ListChange{}
	}
	return c.Result
}
