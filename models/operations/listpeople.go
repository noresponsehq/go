package operations

import (
	"github.com/noresponsehq/go/v11/internal/utils"
	"github.com/noresponsehq/go/v11/models/components"
)

type ListPeopleRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
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

func (l *ListPeopleRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListPeopleRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
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
