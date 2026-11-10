package operations

import (
	"github.com/noresponsehq/go/v10/internal/utils"
	"github.com/noresponsehq/go/v10/models/components"
)

type ListAttendeesRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	EventID     string                    `pathParam:"style=simple,explode=false,name=event_id"`
	// At most this many attendees.
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Read attendees after this one.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Only attendees whose highest finding has this band; none for attendees with no finding.
	Band *components.ListAttendeesBand `queryParam:"style=form,explode=true,name=band"`
	// Only attendees with this decision; undecided for none yet.
	Decision *components.ListAttendeesDecision `queryParam:"style=form,explode=true,name=decision"`
	// Only attendees with this role; unstated for attendees whose role was not given.
	Role *components.ListAttendeesRole `queryParam:"style=form,explode=true,name=role"`
}

func (l ListAttendeesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAttendeesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListAttendeesRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListAttendeesRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListAttendeesRequest) GetEventID() string {
	if l == nil {
		return ""
	}
	return l.EventID
}

func (l *ListAttendeesRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListAttendeesRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

func (l *ListAttendeesRequest) GetBand() *components.ListAttendeesBand {
	if l == nil {
		return nil
	}
	return l.Band
}

func (l *ListAttendeesRequest) GetDecision() *components.ListAttendeesDecision {
	if l == nil {
		return nil
	}
	return l.Decision
}

func (l *ListAttendeesRequest) GetRole() *components.ListAttendeesRole {
	if l == nil {
		return nil
	}
	return l.Role
}

type ListAttendeesResponse struct {
	Headers map[string][]string
	Result  components.AttendeeList
}

func (l *ListAttendeesResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAttendeesResponse) GetResult() components.AttendeeList {
	if l == nil {
		return components.AttendeeList{}
	}
	return l.Result
}
