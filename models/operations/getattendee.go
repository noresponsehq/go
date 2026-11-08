package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetAttendeeRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	EventID     string                    `pathParam:"style=simple,explode=false,name=event_id"`
	AttendeeID  string                    `pathParam:"style=simple,explode=false,name=attendee_id"`
}

func (g *GetAttendeeRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetAttendeeRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetAttendeeRequest) GetEventID() string {
	if g == nil {
		return ""
	}
	return g.EventID
}

func (g *GetAttendeeRequest) GetAttendeeID() string {
	if g == nil {
		return ""
	}
	return g.AttendeeID
}

type GetAttendeeResponse struct {
	Headers map[string][]string
	Result  components.AttendeeReport
}

func (g *GetAttendeeResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetAttendeeResponse) GetResult() components.AttendeeReport {
	if g == nil {
		return components.AttendeeReport{}
	}
	return g.Result
}
