package components

// Event - An event and its screening context.
type Event struct {
	ID          string                   `json:"id"`
	Environment AuditSettingsEnvironment `json:"environment"`
	// The event's name.
	Name string `json:"name"`
	// When the event starts.
	StartsAt string `json:"starts_at"`
	// When the event ends, at or after its start; the start when omitted.
	EndsAt string `json:"ends_at"`
	// Where the event is held.
	Venue *string `json:"venue,omitzero"`
	// The organisation hosting the event.
	Host *string `json:"host,omitzero"`
	// The people the event protects.
	Principals []Principal `json:"principals"`
	// Sensitive topics.
	Topics []string `json:"topics"`
	// How many attendees the event holds.
	Attendees int64 `json:"attendees"`
	// Goes up by one with every change.
	Version int64 `json:"version"`
	// When raw collected items and unmatched candidate profiles are purged: 30 days after the event ends.
	PurgeAfter string `json:"purge_after"`
	// The member or key label that created the event.
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (e *Event) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *Event) GetEnvironment() AuditSettingsEnvironment {
	if e == nil {
		return AuditSettingsEnvironment("")
	}
	return e.Environment
}

func (e *Event) GetName() string {
	if e == nil {
		return ""
	}
	return e.Name
}

func (e *Event) GetStartsAt() string {
	if e == nil {
		return ""
	}
	return e.StartsAt
}

func (e *Event) GetEndsAt() string {
	if e == nil {
		return ""
	}
	return e.EndsAt
}

func (e *Event) GetVenue() *string {
	if e == nil {
		return nil
	}
	return e.Venue
}

func (e *Event) GetHost() *string {
	if e == nil {
		return nil
	}
	return e.Host
}

func (e *Event) GetPrincipals() []Principal {
	if e == nil {
		return []Principal{}
	}
	return e.Principals
}

func (e *Event) GetTopics() []string {
	if e == nil {
		return []string{}
	}
	return e.Topics
}

func (e *Event) GetAttendees() int64 {
	if e == nil {
		return 0
	}
	return e.Attendees
}

func (e *Event) GetVersion() int64 {
	if e == nil {
		return 0
	}
	return e.Version
}

func (e *Event) GetPurgeAfter() string {
	if e == nil {
		return ""
	}
	return e.PurgeAfter
}

func (e *Event) GetCreatedBy() string {
	if e == nil {
		return ""
	}
	return e.CreatedBy
}

func (e *Event) GetCreatedAt() string {
	if e == nil {
		return ""
	}
	return e.CreatedAt
}

func (e *Event) GetUpdatedAt() string {
	if e == nil {
		return ""
	}
	return e.UpdatedAt
}
