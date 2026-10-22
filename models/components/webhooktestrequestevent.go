package components

// WebhookTestRequestEvent - endpoint.verification sends a fresh challenge now, in this call, and the endpoint is verified if it echoes it. assessment.updated queues an event whose data is a fixed example assessment, with test true, delivered and retried like any other; the endpoint must be verified.
type WebhookTestRequestEvent string

const (
	WebhookTestRequestEventEndpointVerification WebhookTestRequestEvent = "endpoint.verification"
	WebhookTestRequestEventAssessmentUpdated    WebhookTestRequestEvent = "assessment.updated"
)

func (e WebhookTestRequestEvent) ToPointer() *WebhookTestRequestEvent {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *WebhookTestRequestEvent) IsExact() bool {
	if e != nil {
		switch *e {
		case "endpoint.verification", "assessment.updated":
			return true
		}
	}
	return false
}
