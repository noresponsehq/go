package components

// WebhookDeliveryList - The 200 answer of GET /v1/webhooks/{endpoint_id}/deliveries: one page of the endpoint's delivery attempts, newest first. Each item is one POST to the endpoint. Attempts are kept for 30 days, or the workspace's retention_days when that is shorter, and are deleted with the assessment they carried. next_cursor is present only when another page exists.
type WebhookDeliveryList struct {
	// Opaque id of the endpoint registration.
	EndpointID string                   `json:"endpoint_id"`
	Items      []WebhookDeliveryAttempt `json:"items"`
	// Opaque cursor for the next page, MAC'd by the server and bound to the workspace and endpoint it was issued for. Pass it back unchanged as cursor; a tampered, foreign or expired cursor answers 400 invalid_cursor.
	NextCursor *string `json:"next_cursor,omitzero"`
}

func (w *WebhookDeliveryList) GetEndpointID() string {
	if w == nil {
		return ""
	}
	return w.EndpointID
}

func (w *WebhookDeliveryList) GetItems() []WebhookDeliveryAttempt {
	if w == nil {
		return []WebhookDeliveryAttempt{}
	}
	return w.Items
}

func (w *WebhookDeliveryList) GetNextCursor() *string {
	if w == nil {
		return nil
	}
	return w.NextCursor
}
