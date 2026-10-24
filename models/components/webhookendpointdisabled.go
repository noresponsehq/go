package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

// WebhookEndpointDisabled - One of the environment's endpoints stopped receiving events: every delivery attempt to it failed from failing_since until disabled_at, 72 hours or more. Its queued deliveries are dropped. Fix the receiver and verify it again with POST /v1/webhooks/{endpoint_id}/test (type endpoint.verification) to resume delivery.
type WebhookEndpointDisabled struct {
	// The endpoint that was disabled.
	EndpointID string `json:"endpoint_id"`
	// The first failed attempt of the unbroken failure streak.
	FailingSince string `json:"failing_since"`
	// When the endpoint was disabled; equal to the event's created_at.
	DisabledAt string `json:"disabled_at"`
}

func (w WebhookEndpointDisabled) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhookEndpointDisabled) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WebhookEndpointDisabled) GetEndpointID() string {
	if w == nil {
		return ""
	}
	return w.EndpointID
}

func (w *WebhookEndpointDisabled) GetFailingSince() string {
	if w == nil {
		return ""
	}
	return w.FailingSince
}

func (w *WebhookEndpointDisabled) GetDisabledAt() string {
	if w == nil {
		return ""
	}
	return w.DisabledAt
}
