package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// ScreeningCompletedEvent - The screening.completed event.
type ScreeningCompletedEvent struct {
	// The event id, equal to the NR-Webhook-Id header and across redeliveries and every endpoint the event fans out to.
	ID string `json:"id"`
	// When the event was created. Equal across redeliveries; the signing timestamp header is per delivery.
	CreatedAt string `json:"created_at"`
	// Always screening.completed.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"screening.completed" json:"type"`
	// Present, and true, only on an event sent by POST /v1/webhooks/{endpoint_id}/test: its data is a fixed example assessment, not one of the workspace's, so a receiver acknowledges it and applies nothing. Real events have no test member.
	Test *bool `json:"test,omitzero"`
	// A screening of a group finished. It counts only the people this screening set out to reach, not who they are: read the roster with GET /v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people, filtered by band, to act on them. people_screened + partial_coverage = people_total, people_flagged is the sum of by_band, no_footprint counts screened people with no flag whose accounts were not found, and no_profile_to_check those the screening had nothing to look up for on any requested platform. A cancelled screening sends no event.
	Data ScreeningCompleted `json:"data"`
}

func (s ScreeningCompletedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningCompletedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningCompletedEvent) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *ScreeningCompletedEvent) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}

func (s *ScreeningCompletedEvent) GetType() string {
	return "screening.completed"
}

func (s *ScreeningCompletedEvent) GetTest() *bool {
	if s == nil {
		return nil
	}
	return s.Test
}

func (s *ScreeningCompletedEvent) GetData() ScreeningCompleted {
	if s == nil {
		return ScreeningCompleted{}
	}
	return s.Data
}
