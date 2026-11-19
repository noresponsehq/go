package components

import (
	"github.com/noresponsehq/go/v17/internal/utils"
)

// InvestigationCompletedEvent - The investigation.completed event.
type InvestigationCompletedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always investigation.completed.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"investigation.completed" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool                  `json:"test,omitzero"`
	Data InvestigationCompleted `json:"data"`
}

func (i InvestigationCompletedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvestigationCompletedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *InvestigationCompletedEvent) GetID() string {
	if i == nil {
		return ""
	}
	return i.ID
}

func (i *InvestigationCompletedEvent) GetCreatedAt() string {
	if i == nil {
		return ""
	}
	return i.CreatedAt
}

func (i *InvestigationCompletedEvent) GetType() string {
	return "investigation.completed"
}

func (i *InvestigationCompletedEvent) GetTest() *bool {
	if i == nil {
		return nil
	}
	return i.Test
}

func (i *InvestigationCompletedEvent) GetData() InvestigationCompleted {
	if i == nil {
		return InvestigationCompleted{}
	}
	return i.Data
}
