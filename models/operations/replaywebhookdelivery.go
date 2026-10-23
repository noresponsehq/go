package operations

import (
	"github.com/noresponsehq/go/v2/models/components"
)

type ReplayWebhookDeliveryRequest struct {
	// The endpoint's whe_ id. An id that is unknown, malformed or another tenant's answers 404.
	EndpointID string `pathParam:"style=simple,explode=false,name=endpoint_id"`
	// The evt_ id of a delivery the endpoint's log holds, as GET /v1/webhooks/{endpoint_id}/deliveries lists it. An id that is malformed, not sent to this endpoint, past the log's retention or another tenant's answers 404.
	EventID string                          `pathParam:"style=simple,explode=false,name=event_id"`
	Body    components.WebhookReplayRequest `request:"mediaType=application/json"`
}

func (r *ReplayWebhookDeliveryRequest) GetEndpointID() string {
	if r == nil {
		return ""
	}
	return r.EndpointID
}

func (r *ReplayWebhookDeliveryRequest) GetEventID() string {
	if r == nil {
		return ""
	}
	return r.EventID
}

func (r *ReplayWebhookDeliveryRequest) GetBody() components.WebhookReplayRequest {
	if r == nil {
		return components.WebhookReplayRequest{}
	}
	return r.Body
}

type ReplayWebhookDeliveryResponse struct {
	Headers map[string][]string
	Result  components.WebhookReplay
}

func (r *ReplayWebhookDeliveryResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReplayWebhookDeliveryResponse) GetResult() components.WebhookReplay {
	if r == nil {
		return components.WebhookReplay{}
	}
	return r.Result
}
