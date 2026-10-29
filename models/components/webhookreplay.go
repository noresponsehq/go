package components

// WebhookReplay - The 202 answer of POST /v1/webhooks/{endpoint_id}/deliveries/{event_id}/replay: the event queued to be sent again.
type WebhookReplay struct {
	// Opaque id of the endpoint registration.
	EndpointID string `json:"endpoint_id"`
	// The replayed event's id. The receiver sees the same id and body as the first time, with a new NR-Webhook-Timestamp and signature.
	EventID string `json:"event_id"`
	// The event's type, as its body's type: assessment.updated, endpoint.verification, account.deleted, workspace.deleted, webhook_endpoint.disabled, account.risk_changed, audit.event, or a type added later.
	EventType string `json:"event_type"`
}

func (w *WebhookReplay) GetEndpointID() string {
	if w == nil {
		return ""
	}
	return w.EndpointID
}

func (w *WebhookReplay) GetEventID() string {
	if w == nil {
		return ""
	}
	return w.EventID
}

func (w *WebhookReplay) GetEventType() string {
	if w == nil {
		return ""
	}
	return w.EventType
}
