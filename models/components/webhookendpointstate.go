package components

// WebhookEndpointState - unverified until the endpoint echoes an endpoint.verification challenge (send one with POST /v1/webhooks/{endpoint_id}/test); verified receives assessment.updated events; disabled after 72 hours of continuous delivery failure, and verified again by the same challenge.
type WebhookEndpointState string

const (
	WebhookEndpointStateUnverified WebhookEndpointState = "unverified"
	WebhookEndpointStateVerified   WebhookEndpointState = "verified"
	WebhookEndpointStateDisabled   WebhookEndpointState = "disabled"
)

func (e WebhookEndpointState) ToPointer() *WebhookEndpointState {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *WebhookEndpointState) IsExact() bool {
	if e != nil {
		switch *e {
		case "unverified", "verified", "disabled":
			return true
		}
	}
	return false
}
