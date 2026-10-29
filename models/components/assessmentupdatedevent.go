package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// AssessmentUpdatedEvent - The assessment.updated event.
type AssessmentUpdatedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always assessment.updated.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"assessment.updated" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// The assessment, or the minimal notice that it expired.
	Data AssessmentUpdatedData `json:"data"`
}

func (a AssessmentUpdatedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssessmentUpdatedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AssessmentUpdatedEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AssessmentUpdatedEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AssessmentUpdatedEvent) GetType() string {
	return "assessment.updated"
}

func (a *AssessmentUpdatedEvent) GetTest() *bool {
	if a == nil {
		return nil
	}
	return a.Test
}

func (a *AssessmentUpdatedEvent) GetData() AssessmentUpdatedData {
	if a == nil {
		return AssessmentUpdatedData{}
	}
	return a.Data
}

func (a *AssessmentUpdatedEvent) GetDataComplete() *Assessment {
	return a.GetData().Assessment
}

func (a *AssessmentUpdatedEvent) GetDataPartial() *Assessment {
	return a.GetData().Assessment
}

func (a *AssessmentUpdatedEvent) GetDataPending() *Assessment {
	return a.GetData().Assessment
}

func (a *AssessmentUpdatedEvent) GetDataExpired() *AssessmentExpired {
	return a.GetData().AssessmentExpired
}
