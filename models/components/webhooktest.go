package components

// WebhookTest - The 202 answer of POST /v1/webhooks/{endpoint_id}/test: the event that was sent or queued.
type WebhookTest struct {
	// The event id the receiver sees as NR-Webhook-Id and as the body's id.
	DeliveryID string `json:"delivery_id"`
	// Opaque id of the endpoint registration.
	EndpointID string `json:"endpoint_id"`
	// endpoint.verification sends a fresh challenge now, in this call, and the endpoint is verified if it echoes it. assessment.updated queues an event whose data is a fixed example assessment, with test true, delivered and retried like any other; the endpoint must be verified.
	Event WebhookTestRequestEvent `json:"event"`
	// unverified until the endpoint echoes an endpoint.verification challenge (send one with POST /v1/webhooks/{endpoint_id}/test); verified receives assessment.updated events; disabled after 72 hours of continuous delivery failure, and verified again by the same challenge.
	State WebhookEndpointState `json:"state"`
	// endpoint.verification only: true when the endpoint echoed the challenge within 5 seconds. A false answer leaves the endpoint's state as it was.
	Verified *bool `json:"verified,omitzero"`
}

func (w *WebhookTest) GetDeliveryID() string {
	if w == nil {
		return ""
	}
	return w.DeliveryID
}

func (w *WebhookTest) GetEndpointID() string {
	if w == nil {
		return ""
	}
	return w.EndpointID
}

func (w *WebhookTest) GetEvent() WebhookTestRequestEvent {
	if w == nil {
		return WebhookTestRequestEvent("")
	}
	return w.Event
}

func (w *WebhookTest) GetState() WebhookEndpointState {
	if w == nil {
		return WebhookEndpointState("")
	}
	return w.State
}

func (w *WebhookTest) GetVerified() *bool {
	if w == nil {
		return nil
	}
	return w.Verified
}
