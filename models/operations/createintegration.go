package operations

import (
	"github.com/noresponsehq/go/v4/models/components"
)

type CreateIntegrationResponse struct {
	Headers map[string][]string
	Result  components.Integration
}

func (c *CreateIntegrationResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateIntegrationResponse) GetResult() components.Integration {
	if c == nil {
		return components.Integration{}
	}
	return c.Result
}
