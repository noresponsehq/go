package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type RotateWebhookSigningKeyRequest struct {
	// The endpoint's whe_ id. An id that is unknown, malformed or another workspace's answers 404.
	EndpointID string                               `pathParam:"style=simple,explode=false,name=endpoint_id"`
	Body       components.WebhookKeyRotationRequest `request:"mediaType=application/json"`
}

func (r *RotateWebhookSigningKeyRequest) GetEndpointID() string {
	if r == nil {
		return ""
	}
	return r.EndpointID
}

func (r *RotateWebhookSigningKeyRequest) GetBody() components.WebhookKeyRotationRequest {
	if r == nil {
		return components.WebhookKeyRotationRequest{}
	}
	return r.Body
}

type RotateWebhookSigningKeyResponse struct {
	Headers map[string][]string
	Result  components.WebhookKeyRotation
}

func (r *RotateWebhookSigningKeyResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RotateWebhookSigningKeyResponse) GetResult() components.WebhookKeyRotation {
	if r == nil {
		return components.WebhookKeyRotation{}
	}
	return r.Result
}
