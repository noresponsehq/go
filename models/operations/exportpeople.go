package operations

import (
	"github.com/noresponsehq/go/v20/internal/utils"
	"github.com/noresponsehq/go/v20/models/components"
)

type ExportPeopleRequest struct {
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
	// At most this many people.
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Read people after this one.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Only people with this decision; undecided for none yet.
	Decision *components.ListPeopleDecision `queryParam:"style=form,explode=true,name=decision"`
	// Only people with this role; unstated for people whose role was not given.
	Role *components.ListPeopleRole `queryParam:"style=form,explode=true,name=role"`
}

func (e ExportPeopleRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExportPeopleRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *ExportPeopleRequest) GetGroupID() string {
	if e == nil {
		return ""
	}
	return e.GroupID
}

func (e *ExportPeopleRequest) GetLimit() *int64 {
	if e == nil {
		return nil
	}
	return e.Limit
}

func (e *ExportPeopleRequest) GetAfter() *string {
	if e == nil {
		return nil
	}
	return e.After
}

func (e *ExportPeopleRequest) GetDecision() *components.ListPeopleDecision {
	if e == nil {
		return nil
	}
	return e.Decision
}

func (e *ExportPeopleRequest) GetRole() *components.ListPeopleRole {
	if e == nil {
		return nil
	}
	return e.Role
}

type ExportPeopleResponse struct {
	Headers map[string][]string
	Result  components.PersonList
}

func (e *ExportPeopleResponse) GetHeaders() map[string][]string {
	if e == nil {
		return map[string][]string{}
	}
	return e.Headers
}

func (e *ExportPeopleResponse) GetResult() components.PersonList {
	if e == nil {
		return components.PersonList{}
	}
	return e.Result
}
