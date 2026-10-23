package components

type WebhookDeliveryAttempt struct {
	// The event id the receiver saw as NR-Webhook-Id and as the body's id. Replay it with POST /v1/webhooks/{endpoint_id}/deliveries/{event_id}/replay.
	EventID string `json:"event_id"`
	// The event's type, as its body's type: assessment.updated, account.deleted or account.risk_changed.
	EventType string `json:"event_type"`
	// Which attempt of this delivery this was: 1 for the first send, up to 9 with retries. A replay counts its own attempts from 1.
	Attempt int64 `json:"attempt"`
	// true for an attempt of a replay you asked for.
	Replay bool `json:"replay"`
	// delivered when the endpoint answered 2xx; failed otherwise, with failure naming why.
	Status WebhookDeliveryAttemptStatus `json:"status"`
	// The HTTP status the endpoint answered. Absent when no response was read.
	ResponseCode *int64 `json:"response_code,omitzero"`
	// Why the attempt failed: the URL is no longer valid, the host did not resolve, it resolved to an address that is not public, no answer within 5 seconds, the connection failed, a 3xx (redirects are not followed) or another non-2xx status. Absent for a delivered attempt.
	Failure *Failure `json:"failure,omitzero"`
	// Milliseconds from sending the request to the answer or the failure.
	LatencyMs int64 `json:"latency_ms"`
	// When the request was sent. Its NR-Webhook-Timestamp is this time in Unix seconds.
	AttemptedAt string `json:"attempted_at"`
}

func (w *WebhookDeliveryAttempt) GetEventID() string {
	if w == nil {
		return ""
	}
	return w.EventID
}

func (w *WebhookDeliveryAttempt) GetEventType() string {
	if w == nil {
		return ""
	}
	return w.EventType
}

func (w *WebhookDeliveryAttempt) GetAttempt() int64 {
	if w == nil {
		return 0
	}
	return w.Attempt
}

func (w *WebhookDeliveryAttempt) GetReplay() bool {
	if w == nil {
		return false
	}
	return w.Replay
}

func (w *WebhookDeliveryAttempt) GetStatus() WebhookDeliveryAttemptStatus {
	if w == nil {
		return WebhookDeliveryAttemptStatus("")
	}
	return w.Status
}

func (w *WebhookDeliveryAttempt) GetResponseCode() *int64 {
	if w == nil {
		return nil
	}
	return w.ResponseCode
}

func (w *WebhookDeliveryAttempt) GetFailure() *Failure {
	if w == nil {
		return nil
	}
	return w.Failure
}

func (w *WebhookDeliveryAttempt) GetLatencyMs() int64 {
	if w == nil {
		return 0
	}
	return w.LatencyMs
}

func (w *WebhookDeliveryAttempt) GetAttemptedAt() string {
	if w == nil {
		return ""
	}
	return w.AttemptedAt
}
