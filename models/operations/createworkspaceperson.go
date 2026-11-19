package operations

import (
	"github.com/noresponsehq/go/v16/models/components"
)

type CreateWorkspacePersonRequest struct {
	// Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person.
	IdempotencyKey *string                           `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.WorkspacePersonRequest `request:"mediaType=application/json"`
}

func (c *CreateWorkspacePersonRequest) GetIdempotencyKey() *string {
	if c == nil {
		return nil
	}
	return c.IdempotencyKey
}

func (c *CreateWorkspacePersonRequest) GetBody() components.WorkspacePersonRequest {
	if c == nil {
		return components.WorkspacePersonRequest{}
	}
	return c.Body
}

type CreateWorkspacePersonResponse struct {
	Headers map[string][]string
	Result  components.WorkspacePersonResult
}

func (c *CreateWorkspacePersonResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateWorkspacePersonResponse) GetResult() components.WorkspacePersonResult {
	if c == nil {
		return components.WorkspacePersonResult{}
	}
	return c.Result
}
