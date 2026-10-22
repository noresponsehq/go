package operations

type DeleteWebhookEndpointRequest struct {
	// The endpoint's whe_ id. An id that is unknown, malformed or another tenant's answers 404.
	EndpointID string `pathParam:"style=simple,explode=false,name=endpoint_id"`
}

func (d *DeleteWebhookEndpointRequest) GetEndpointID() string {
	if d == nil {
		return ""
	}
	return d.EndpointID
}

type DeleteWebhookEndpointResponse struct {
	Headers map[string][]string
}

func (d *DeleteWebhookEndpointResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}
