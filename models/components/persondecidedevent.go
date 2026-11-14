package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// PersonDecidedEvent - The person.decided event.
type PersonDecidedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always person.decided.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"person.decided" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// A reviewer recorded a decision on a person. The decision's note stays in the API. Sent once per recorded review; a person's latest review is its decision, so order by decided_at.
	Data PersonDecided `json:"data"`
}

func (p PersonDecidedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonDecidedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonDecidedEvent) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PersonDecidedEvent) GetCreatedAt() string {
	if p == nil {
		return ""
	}
	return p.CreatedAt
}

func (p *PersonDecidedEvent) GetType() string {
	return "person.decided"
}

func (p *PersonDecidedEvent) GetTest() *bool {
	if p == nil {
		return nil
	}
	return p.Test
}

func (p *PersonDecidedEvent) GetData() PersonDecided {
	if p == nil {
		return PersonDecided{}
	}
	return p.Data
}
