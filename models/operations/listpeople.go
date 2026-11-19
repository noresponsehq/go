package operations

import (
	"github.com/noresponsehq/go/v17/internal/utils"
	"github.com/noresponsehq/go/v17/models/components"
)

type ListPeopleRequest struct {
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// At most this many people.
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Read people after this one.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Only people whose highest finding has this band; none for people with no finding.
	Band *components.ListPeopleBand `queryParam:"style=form,explode=true,name=band"`
	// Only people with this decision; undecided for none yet.
	Decision *components.ListPeopleDecision `queryParam:"style=form,explode=true,name=decision"`
	// Only people with this role; unstated for people whose role was not given.
	Role *components.ListPeopleRole `queryParam:"style=form,explode=true,name=role"`
}

func (l ListPeopleRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPeopleRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListPeopleRequest) GetGroupID() string {
	if l == nil {
		return ""
	}
	return l.GroupID
}

func (l *ListPeopleRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListPeopleRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

func (l *ListPeopleRequest) GetBand() *components.ListPeopleBand {
	if l == nil {
		return nil
	}
	return l.Band
}

func (l *ListPeopleRequest) GetDecision() *components.ListPeopleDecision {
	if l == nil {
		return nil
	}
	return l.Decision
}

func (l *ListPeopleRequest) GetRole() *components.ListPeopleRole {
	if l == nil {
		return nil
	}
	return l.Role
}

type ListPeopleResponse struct {
	Headers map[string][]string
	Result  components.PersonList
}

func (l *ListPeopleResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListPeopleResponse) GetResult() components.PersonList {
	if l == nil {
		return components.PersonList{}
	}
	return l.Result
}
