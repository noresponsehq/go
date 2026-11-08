package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type AddAttendeesRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	EventID     string                    `pathParam:"style=simple,explode=false,name=event_id"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                          `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.AttendeeBatchRequest `request:"mediaType=application/json"`
}

func (a *AddAttendeesRequest) GetWorkspaceID() string {
	if a == nil {
		return ""
	}
	return a.WorkspaceID
}

func (a *AddAttendeesRequest) GetEnvironment() components.KeyEnvironment {
	if a == nil {
		return components.KeyEnvironment("")
	}
	return a.Environment
}

func (a *AddAttendeesRequest) GetEventID() string {
	if a == nil {
		return ""
	}
	return a.EventID
}

func (a *AddAttendeesRequest) GetIdempotencyKey() string {
	if a == nil {
		return ""
	}
	return a.IdempotencyKey
}

func (a *AddAttendeesRequest) GetBody() components.AttendeeBatchRequest {
	if a == nil {
		return components.AttendeeBatchRequest{}
	}
	return a.Body
}

type AddAttendeesResponse struct {
	Headers map[string][]string
	Result  components.AttendeeBatch
}

func (a *AddAttendeesResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AddAttendeesResponse) GetResult() components.AttendeeBatch {
	if a == nil {
		return components.AttendeeBatch{}
	}
	return a.Result
}
