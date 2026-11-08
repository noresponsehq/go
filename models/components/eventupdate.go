package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// EventUpdate - Changes the fields it names; at least one besides expected_version.
type EventUpdate struct {
	// The event's name.
	Name *string `json:"name,omitzero"`
	// When the event starts.
	StartsAt *string `json:"starts_at,omitzero"`
	// When the event ends, at or after its start; the start when omitted.
	EndsAt *string `json:"ends_at,omitzero"`
	// Where the event is held.
	Venue *string `json:"venue,omitzero"`
	// The organisation hosting the event.
	Host *string `json:"host,omitzero"`
	// The people the event protects, at most 20.
	Principals []Principal `json:"principals,omitzero"`
	// Sensitive topics the screening reads posts against, at most 50.
	Topics []string `json:"topics,omitzero"`
	// When present, the change applies only if the event is still at this version (409 event_conflict otherwise).
	ExpectedVersion *int64 `json:"expected_version,omitzero"`
}

func (e EventUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *EventUpdate) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *EventUpdate) GetStartsAt() *string {
	if e == nil {
		return nil
	}
	return e.StartsAt
}

func (e *EventUpdate) GetEndsAt() *string {
	if e == nil {
		return nil
	}
	return e.EndsAt
}

func (e *EventUpdate) GetVenue() *string {
	if e == nil {
		return nil
	}
	return e.Venue
}

func (e *EventUpdate) GetHost() *string {
	if e == nil {
		return nil
	}
	return e.Host
}

func (e *EventUpdate) GetPrincipals() []Principal {
	if e == nil {
		return nil
	}
	return e.Principals
}

func (e *EventUpdate) GetTopics() []string {
	if e == nil {
		return nil
	}
	return e.Topics
}

func (e *EventUpdate) GetExpectedVersion() *int64 {
	if e == nil {
		return nil
	}
	return e.ExpectedVersion
}
