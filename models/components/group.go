package components

// Group - A group and its screening context.
type Group struct {
	ID          string                   `json:"id"`
	Environment AuditSettingsEnvironment `json:"environment"`
	// The group's name.
	Name string `json:"name"`
	// When the group starts.
	StartsAt string `json:"starts_at"`
	// When the group ends, at or after its start; the start when omitted.
	EndsAt string `json:"ends_at"`
	// Where the group is held.
	Venue *string `json:"venue,omitzero"`
	// The organisation hosting the group.
	Host *string `json:"host,omitzero"`
	// The people the group protects.
	Principals []Principal `json:"principals"`
	// Sensitive topics.
	Topics []string `json:"topics"`
	// How many people the group holds.
	People int64 `json:"people"`
	// Goes up by one with every change.
	Version int64 `json:"version"`
	// When raw collected items and unmatched candidate profiles are purged: 30 days after the group ends.
	PurgeAfter string `json:"purge_after"`
	// The member or key label that created the group.
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (g *Group) GetID() string {
	if g == nil {
		return ""
	}
	return g.ID
}

func (g *Group) GetEnvironment() AuditSettingsEnvironment {
	if g == nil {
		return AuditSettingsEnvironment("")
	}
	return g.Environment
}

func (g *Group) GetName() string {
	if g == nil {
		return ""
	}
	return g.Name
}

func (g *Group) GetStartsAt() string {
	if g == nil {
		return ""
	}
	return g.StartsAt
}

func (g *Group) GetEndsAt() string {
	if g == nil {
		return ""
	}
	return g.EndsAt
}

func (g *Group) GetVenue() *string {
	if g == nil {
		return nil
	}
	return g.Venue
}

func (g *Group) GetHost() *string {
	if g == nil {
		return nil
	}
	return g.Host
}

func (g *Group) GetPrincipals() []Principal {
	if g == nil {
		return []Principal{}
	}
	return g.Principals
}

func (g *Group) GetTopics() []string {
	if g == nil {
		return []string{}
	}
	return g.Topics
}

func (g *Group) GetPeople() int64 {
	if g == nil {
		return 0
	}
	return g.People
}

func (g *Group) GetVersion() int64 {
	if g == nil {
		return 0
	}
	return g.Version
}

func (g *Group) GetPurgeAfter() string {
	if g == nil {
		return ""
	}
	return g.PurgeAfter
}

func (g *Group) GetCreatedBy() string {
	if g == nil {
		return ""
	}
	return g.CreatedBy
}

func (g *Group) GetCreatedAt() string {
	if g == nil {
		return ""
	}
	return g.CreatedAt
}

func (g *Group) GetUpdatedAt() string {
	if g == nil {
		return ""
	}
	return g.UpdatedAt
}
