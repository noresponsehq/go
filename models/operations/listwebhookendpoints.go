package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListWebhookEndpointsResponse struct {
	Headers map[string][]string
	Result  components.WebhookEndpointList
}

func (l *ListWebhookEndpointsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListWebhookEndpointsResponse) GetResult() components.WebhookEndpointList {
	if l == nil {
		return components.WebhookEndpointList{}
	}
	return l.Result
}
