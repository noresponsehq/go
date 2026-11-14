package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// PersonFlaggedEvent - The person.flagged event.
type PersonFlaggedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always person.flagged.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"person.flagged" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// A screening recorded a finding on a person. It says which person and how severe, not what they posted, who they are or why: GET the person report for the findings. Sent once per screening that flags the person, so a re-screen of someone on watch announces them again.
	Data PersonFlagged `json:"data"`
}

func (p PersonFlaggedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonFlaggedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonFlaggedEvent) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PersonFlaggedEvent) GetCreatedAt() string {
	if p == nil {
		return ""
	}
	return p.CreatedAt
}

func (p *PersonFlaggedEvent) GetType() string {
	return "person.flagged"
}

func (p *PersonFlaggedEvent) GetTest() *bool {
	if p == nil {
		return nil
	}
	return p.Test
}

func (p *PersonFlaggedEvent) GetData() PersonFlagged {
	if p == nil {
		return PersonFlagged{}
	}
	return p.Data
}
