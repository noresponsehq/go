package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AttendeeFlaggedEvent - The attendee.flagged event.
type AttendeeFlaggedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always attendee.flagged.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"attendee.flagged" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// A screening recorded a finding on an attendee. It says which attendee and how severe, not what they posted, who they are or why: GET the attendee report for the findings. Sent once per screening that flags the attendee, so a re-screen of someone on watch announces them again.
	Data AttendeeFlagged `json:"data"`
}

func (a AttendeeFlaggedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttendeeFlaggedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttendeeFlaggedEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AttendeeFlaggedEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AttendeeFlaggedEvent) GetType() string {
	return "attendee.flagged"
}

func (a *AttendeeFlaggedEvent) GetTest() *bool {
	if a == nil {
		return nil
	}
	return a.Test
}

func (a *AttendeeFlaggedEvent) GetData() AttendeeFlagged {
	if a == nil {
		return AttendeeFlagged{}
	}
	return a.Data
}
