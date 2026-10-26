package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// WebhookEndpointDisabledEvent - The webhook_endpoint.disabled event.
type WebhookEndpointDisabledEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always webhook_endpoint.disabled.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"webhook_endpoint.disabled" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// One of the environment's endpoints stopped receiving events: every delivery attempt to it failed from failing_since until disabled_at, 72 hours or more. Its queued deliveries are dropped. Fix the receiver and verify it again with POST /v1/webhooks/{endpoint_id}/test (type endpoint.verification) to resume delivery.
	Data WebhookEndpointDisabled `json:"data"`
}

func (w WebhookEndpointDisabledEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhookEndpointDisabledEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WebhookEndpointDisabledEvent) GetID() string {
	if w == nil {
		return ""
	}
	return w.ID
}

func (w *WebhookEndpointDisabledEvent) GetCreatedAt() string {
	if w == nil {
		return ""
	}
	return w.CreatedAt
}

func (w *WebhookEndpointDisabledEvent) GetType() string {
	return "webhook_endpoint.disabled"
}

func (w *WebhookEndpointDisabledEvent) GetTest() *bool {
	if w == nil {
		return nil
	}
	return w.Test
}

func (w *WebhookEndpointDisabledEvent) GetData() WebhookEndpointDisabled {
	if w == nil {
		return WebhookEndpointDisabled{}
	}
	return w.Data
}
