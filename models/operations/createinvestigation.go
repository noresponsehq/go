package operations

import (
	"github.com/noresponsehq/go/v14/models/components"
)

type CreateInvestigationRequest struct {
	// Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person.
	IdempotencyKey string                          `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.InvestigationRequest `request:"mediaType=application/json"`
}

func (c *CreateInvestigationRequest) GetIdempotencyKey() string {
	if c == nil {
		return ""
	}
	return c.IdempotencyKey
}

func (c *CreateInvestigationRequest) GetBody() components.InvestigationRequest {
	if c == nil {
		return components.InvestigationRequest{}
	}
	return c.Body
}

type CreateInvestigationResponse struct {
	Headers map[string][]string
	Result  components.Investigation
}

func (c *CreateInvestigationResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateInvestigationResponse) GetResult() components.Investigation {
	if c == nil {
		return components.Investigation{}
	}
	return c.Result
}
