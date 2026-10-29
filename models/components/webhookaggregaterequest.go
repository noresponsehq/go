package components

// WebhookAggregateRequest - Count every retained matching attempt in [from,to). Capture is best effort; missing attempts cannot be inferred. UTC buckets may be partial at the edges. The window is at most 30 days.
type WebhookAggregateRequest struct {
	From   string                         `json:"from"`
	To     string                         `json:"to"`
	Bucket *RequestAggregateRequestBucket `json:"bucket,omitzero"`
	Status *int64                         `json:"status,omitzero"`
	// Opaque id of the endpoint registration.
	EndpointID *string `json:"endpoint_id,omitzero"`
}

func (w *WebhookAggregateRequest) GetFrom() string {
	if w == nil {
		return ""
	}
	return w.From
}

func (w *WebhookAggregateRequest) GetTo() string {
	if w == nil {
		return ""
	}
	return w.To
}

func (w *WebhookAggregateRequest) GetBucket() *RequestAggregateRequestBucket {
	if w == nil {
		return nil
	}
	return w.Bucket
}

func (w *WebhookAggregateRequest) GetStatus() *int64 {
	if w == nil {
		return nil
	}
	return w.Status
}

func (w *WebhookAggregateRequest) GetEndpointID() *string {
	if w == nil {
		return nil
	}
	return w.EndpointID
}
