package operations

type DeleteIntegrationRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the integration_id shape, or another workspace's id, answers 404.
	IntegrationID string `pathParam:"style=simple,explode=false,name=integration_id"`
}

func (d *DeleteIntegrationRequest) GetIntegrationID() string {
	if d == nil {
		return ""
	}
	return d.IntegrationID
}

type DeleteIntegrationResponse struct {
	Headers map[string][]string
}

func (d *DeleteIntegrationResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}
