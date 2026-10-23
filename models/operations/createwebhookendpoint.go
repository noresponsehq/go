package operations

import (
	"github.com/noresponsehq/go/v4/models/components"
)

type CreateWebhookEndpointResponse struct {
	Headers map[string][]string
	Result  components.WebhookEndpointCreated
}

func (c *CreateWebhookEndpointResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateWebhookEndpointResponse) GetResult() components.WebhookEndpointCreated {
	if c == nil {
		return components.WebhookEndpointCreated{}
	}
	return c.Result
}
