package operations

import (
	"github.com/noresponsehq/go/v19/models/components"
)

type CancelScreeningRequest struct {
	GroupID     string                            `pathParam:"style=simple,explode=false,name=group_id"`
	ScreeningID string                            `pathParam:"style=simple,explode=false,name=screening_id"`
	Body        components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (c *CancelScreeningRequest) GetGroupID() string {
	if c == nil {
		return ""
	}
	return c.GroupID
}

func (c *CancelScreeningRequest) GetScreeningID() string {
	if c == nil {
		return ""
	}
	return c.ScreeningID
}

func (c *CancelScreeningRequest) GetBody() components.ScreeningCancelRequest {
	if c == nil {
		return components.ScreeningCancelRequest{}
	}
	return c.Body
}

type CancelScreeningResponse struct {
	Headers map[string][]string
	Result  components.Screening
}

func (c *CancelScreeningResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CancelScreeningResponse) GetResult() components.Screening {
	if c == nil {
		return components.Screening{}
	}
	return c.Result
}
