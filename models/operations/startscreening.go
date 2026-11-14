package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type StartScreeningRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                      `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.ScreeningRequest `request:"mediaType=application/json"`
}

func (s *StartScreeningRequest) GetWorkspaceID() string {
	if s == nil {
		return ""
	}
	return s.WorkspaceID
}

func (s *StartScreeningRequest) GetEnvironment() components.KeyEnvironment {
	if s == nil {
		return components.KeyEnvironment("")
	}
	return s.Environment
}

func (s *StartScreeningRequest) GetGroupID() string {
	if s == nil {
		return ""
	}
	return s.GroupID
}

func (s *StartScreeningRequest) GetIdempotencyKey() string {
	if s == nil {
		return ""
	}
	return s.IdempotencyKey
}

func (s *StartScreeningRequest) GetBody() components.ScreeningRequest {
	if s == nil {
		return components.ScreeningRequest{}
	}
	return s.Body
}

type StartScreeningResponse struct {
	Headers map[string][]string
	Result  components.Screening
}

func (s *StartScreeningResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *StartScreeningResponse) GetResult() components.Screening {
	if s == nil {
		return components.Screening{}
	}
	return s.Result
}
