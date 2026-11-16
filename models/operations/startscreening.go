package operations

import (
	"github.com/noresponsehq/go/v15/models/components"
)

type StartScreeningRequest struct {
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                      `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.ScreeningRequest `request:"mediaType=application/json"`
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
