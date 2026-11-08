package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AttendeeDecidedEvent - The attendee.decided event.
type AttendeeDecidedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always attendee.decided.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"attendee.decided" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// A reviewer recorded a decision on an attendee. The decision's note stays in the API. Sent once per recorded review; an attendee's latest review is its decision, so order by decided_at.
	Data AttendeeDecided `json:"data"`
}

func (a AttendeeDecidedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttendeeDecidedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttendeeDecidedEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AttendeeDecidedEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AttendeeDecidedEvent) GetType() string {
	return "attendee.decided"
}

func (a *AttendeeDecidedEvent) GetTest() *bool {
	if a == nil {
		return nil
	}
	return a.Test
}

func (a *AttendeeDecidedEvent) GetData() AttendeeDecided {
	if a == nil {
		return AttendeeDecided{}
	}
	return a.Data
}
