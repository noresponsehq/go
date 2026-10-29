package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type SendWebhookTestEventRequest struct {
	// The endpoint's whe_ id. An id that is unknown, malformed or another workspace's answers 404.
	EndpointID string                        `pathParam:"style=simple,explode=false,name=endpoint_id"`
	Body       components.WebhookTestRequest `request:"mediaType=application/json"`
}

func (s *SendWebhookTestEventRequest) GetEndpointID() string {
	if s == nil {
		return ""
	}
	return s.EndpointID
}

func (s *SendWebhookTestEventRequest) GetBody() components.WebhookTestRequest {
	if s == nil {
		return components.WebhookTestRequest{}
	}
	return s.Body
}

type SendWebhookTestEventResponse struct {
	Headers map[string][]string
	Result  components.WebhookTest
}

func (s *SendWebhookTestEventResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *SendWebhookTestEventResponse) GetResult() components.WebhookTest {
	if s == nil {
		return components.WebhookTest{}
	}
	return s.Result
}
