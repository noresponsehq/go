package components

// WebhookEndpointList - GET /v1/webhooks answer: every endpoint of the calling tenant, in every state, ordered by endpoint_id. A tenant holds at most 5, so there is no page.
type WebhookEndpointList struct {
	Items []WebhookEndpoint `json:"items"`
}

func (w *WebhookEndpointList) GetItems() []WebhookEndpoint {
	if w == nil {
		return []WebhookEndpoint{}
	}
	return w.Items
}
