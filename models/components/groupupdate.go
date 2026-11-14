package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// GroupUpdate - Changes the fields it names; at least one besides expected_version.
type GroupUpdate struct {
	// The group's name.
	Name *string `json:"name,omitzero"`
	// When the group starts.
	StartsAt *string `json:"starts_at,omitzero"`
	// When the group ends, at or after its start; the start when omitted.
	EndsAt *string `json:"ends_at,omitzero"`
	// Where the group is held.
	Venue *string `json:"venue,omitzero"`
	// The organisation hosting the group.
	Host *string `json:"host,omitzero"`
	// The people the group protects, at most 20.
	Principals []Principal `json:"principals,omitzero"`
	// Sensitive topics the screening reads posts against, at most 50.
	Topics []string `json:"topics,omitzero"`
	// When present, the change applies only if the group is still at this version (409 group_conflict otherwise).
	ExpectedVersion *int64 `json:"expected_version,omitzero"`
}

func (g GroupUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GroupUpdate) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GroupUpdate) GetStartsAt() *string {
	if g == nil {
		return nil
	}
	return g.StartsAt
}

func (g *GroupUpdate) GetEndsAt() *string {
	if g == nil {
		return nil
	}
	return g.EndsAt
}

func (g *GroupUpdate) GetVenue() *string {
	if g == nil {
		return nil
	}
	return g.Venue
}

func (g *GroupUpdate) GetHost() *string {
	if g == nil {
		return nil
	}
	return g.Host
}

func (g *GroupUpdate) GetPrincipals() []Principal {
	if g == nil {
		return nil
	}
	return g.Principals
}

func (g *GroupUpdate) GetTopics() []string {
	if g == nil {
		return nil
	}
	return g.Topics
}

func (g *GroupUpdate) GetExpectedVersion() *int64 {
	if g == nil {
		return nil
	}
	return g.ExpectedVersion
}
