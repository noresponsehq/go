package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// EventRequest - Creates an event.
type EventRequest struct {
	// The event's name.
	Name string `json:"name"`
	// When the event starts.
	StartsAt string `json:"starts_at"`
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
}

func (e EventRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *EventRequest) GetName() string {
	if e == nil {
		return ""
	}
	return e.Name
}

func (e *EventRequest) GetStartsAt() string {
	if e == nil {
		return ""
	}
	return e.StartsAt
}

func (e *EventRequest) GetEndsAt() *string {
	if e == nil {
		return nil
	}
	return e.EndsAt
}

func (e *EventRequest) GetVenue() *string {
	if e == nil {
		return nil
	}
	return e.Venue
}

func (e *EventRequest) GetHost() *string {
	if e == nil {
		return nil
	}
	return e.Host
}

func (e *EventRequest) GetPrincipals() []Principal {
	if e == nil {
		return nil
	}
	return e.Principals
}

func (e *EventRequest) GetTopics() []string {
	if e == nil {
		return nil
	}
	return e.Topics
}
