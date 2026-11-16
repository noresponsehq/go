package operations

import (
	"github.com/noresponsehq/go/v15/models/components"
)

type CancelInvestigationRequest struct {
	InvestigationID string                            `pathParam:"style=simple,explode=false,name=investigation_id"`
	Body            components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (c *CancelInvestigationRequest) GetInvestigationID() string {
	if c == nil {
		return ""
	}
	return c.InvestigationID
}

func (c *CancelInvestigationRequest) GetBody() components.ScreeningCancelRequest {
	if c == nil {
		return components.ScreeningCancelRequest{}
	}
	return c.Body
}

type CancelInvestigationResponse struct {
	Headers map[string][]string
	Result  components.Investigation
}

func (c *CancelInvestigationResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CancelInvestigationResponse) GetResult() components.Investigation {
	if c == nil {
		return components.Investigation{}
	}
	return c.Result
}
