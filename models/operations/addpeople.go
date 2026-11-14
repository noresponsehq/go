package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type AddPeopleRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                        `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.PersonBatchRequest `request:"mediaType=application/json"`
}

func (a *AddPeopleRequest) GetWorkspaceID() string {
	if a == nil {
		return ""
	}
	return a.WorkspaceID
}

func (a *AddPeopleRequest) GetEnvironment() components.KeyEnvironment {
	if a == nil {
		return components.KeyEnvironment("")
	}
	return a.Environment
}

func (a *AddPeopleRequest) GetGroupID() string {
	if a == nil {
		return ""
	}
	return a.GroupID
}

func (a *AddPeopleRequest) GetIdempotencyKey() string {
	if a == nil {
		return ""
	}
	return a.IdempotencyKey
}

func (a *AddPeopleRequest) GetBody() components.PersonBatchRequest {
	if a == nil {
		return components.PersonBatchRequest{}
	}
	return a.Body
}

type AddPeopleResponse struct {
	Headers map[string][]string
	Result  components.PersonBatch
}

func (a *AddPeopleResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AddPeopleResponse) GetResult() components.PersonBatch {
	if a == nil {
		return components.PersonBatch{}
	}
	return a.Result
}
