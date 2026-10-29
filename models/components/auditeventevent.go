package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AuditEventEvent - The audit.event event.
type AuditEventEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always audit.event.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"audit.event" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// One privileged workspace action with its actor, target, outcome and time. Successful data changes commit their audit evidence atomically. Sensitive exports record access before returning bytes. Details exclude secrets and subject identifiers.
	Data AuditEvent `json:"data"`
}

func (a AuditEventEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditEventEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuditEventEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AuditEventEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AuditEventEvent) GetType() string {
	return "audit.event"
}

func (a *AuditEventEvent) GetTest() *bool {
	if a == nil {
		return nil
	}
	return a.Test
}

func (a *AuditEventEvent) GetData() AuditEvent {
	if a == nil {
		return AuditEvent{}
	}
	return a.Data
}
