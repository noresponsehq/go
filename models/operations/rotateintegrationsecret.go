package operations

import (
	"github.com/noresponsehq/go/v4/models/components"
)

type RotateIntegrationSecretRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the integration_id shape, or another tenant's id, answers 404.
	IntegrationID string                                `pathParam:"style=simple,explode=false,name=integration_id"`
	Body          components.IntegrationRotationRequest `request:"mediaType=application/json"`
}

func (r *RotateIntegrationSecretRequest) GetIntegrationID() string {
	if r == nil {
		return ""
	}
	return r.IntegrationID
}

func (r *RotateIntegrationSecretRequest) GetBody() components.IntegrationRotationRequest {
	if r == nil {
		return components.IntegrationRotationRequest{}
	}
	return r.Body
}

type RotateIntegrationSecretResponse struct {
	Headers map[string][]string
	Result  components.Integration
}

func (r *RotateIntegrationSecretResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RotateIntegrationSecretResponse) GetResult() components.Integration {
	if r == nil {
		return components.Integration{}
	}
	return r.Result
}
