package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CreateAssessmentRequest struct {
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                       `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.AssessmentRequest `request:"mediaType=application/json"`
}

func (c *CreateAssessmentRequest) GetIdempotencyKey() string {
	if c == nil {
		return ""
	}
	return c.IdempotencyKey
}

func (c *CreateAssessmentRequest) GetBody() components.AssessmentRequest {
	if c == nil {
		return components.AssessmentRequest{}
	}
	return c.Body
}

type CreateAssessmentResponse struct {
	Headers map[string][]string
	Result  components.Assessment
}

func (c *CreateAssessmentResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateAssessmentResponse) GetResult() components.Assessment {
	if c == nil {
		return components.Assessment{}
	}
	return c.Result
}
