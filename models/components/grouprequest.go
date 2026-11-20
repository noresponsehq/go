package components

import (
	"github.com/noresponsehq/go/v20/internal/utils"
	"github.com/noresponsehq/go/v20/optionalnullable"
)

// GroupRequest - Creates an ongoing group or an event with optional dates.
type GroupRequest struct {
	// The group's name.
	Name string `json:"name"`
	// What the group's people share: an event, an organisation, a place or another attribute.
	Type *GroupType `json:"type,omitzero"`
	// The attribute its people share, such as the employer's or the place's name.
	Attribute *string `json:"attribute,omitzero"`
	// Optional event start; null for an ongoing group.
	StartsAt optionalnullable.OptionalNullable[string] `json:"starts_at,omitzero"`
	// Optional event end; requires starts_at and must be at or after it. Null when no end is set.
	EndsAt optionalnullable.OptionalNullable[string] `json:"ends_at,omitzero"`
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

func (g *GroupRequest) GetType() *GroupType {
	if g == nil {
		return nil
	}
	return g.Type
}

func (g *GroupRequest) GetAttribute() *string {
	if g == nil {
		return nil
	}
	return g.Attribute
}

func (g *GroupRequest) GetStartsAt() optionalnullable.OptionalNullable[string] {
	if g == nil {
		return nil
	}
	return g.StartsAt
}

func (g *GroupRequest) GetEndsAt() optionalnullable.OptionalNullable[string] {
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
