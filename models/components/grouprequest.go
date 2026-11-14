package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// GroupRequest - Creates a group.
type GroupRequest struct {
	// The group's name.
	Name string `json:"name"`
	// When the group starts.
	StartsAt string `json:"starts_at"`
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
}

func (g GroupRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (g *GroupRequest) GetName() string {
	if g == nil {
		return ""
	}
	return g.Name
}

func (g *GroupRequest) GetStartsAt() string {
	if g == nil {
		return ""
	}
	return g.StartsAt
}

func (g *GroupRequest) GetEndsAt() *string {
	if g == nil {
		return nil
	}
	return g.EndsAt
}

func (g *GroupRequest) GetVenue() *string {
	if g == nil {
		return nil
	}
	return g.Venue
}

func (g *GroupRequest) GetHost() *string {
	if g == nil {
		return nil
	}
	return g.Host
}

func (g *GroupRequest) GetPrincipals() []Principal {
	if g == nil {
		return nil
	}
	return g.Principals
}

func (g *GroupRequest) GetTopics() []string {
	if g == nil {
		return nil
	}
	return g.Topics
}
