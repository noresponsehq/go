package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ReviewAttendeeRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment        `pathParam:"style=simple,explode=false,name=environment"`
	EventID     string                           `pathParam:"style=simple,explode=false,name=event_id"`
	AttendeeID  string                           `pathParam:"style=simple,explode=false,name=attendee_id"`
	Body        components.AttendeeReviewRequest `request:"mediaType=application/json"`
}

func (r *ReviewAttendeeRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *ReviewAttendeeRequest) GetEnvironment() components.KeyEnvironment {
	if r == nil {
		return components.KeyEnvironment("")
	}
	return r.Environment
}

func (r *ReviewAttendeeRequest) GetEventID() string {
	if r == nil {
		return ""
	}
	return r.EventID
}

func (r *ReviewAttendeeRequest) GetAttendeeID() string {
	if r == nil {
		return ""
	}
	return r.AttendeeID
}

func (r *ReviewAttendeeRequest) GetBody() components.AttendeeReviewRequest {
	if r == nil {
		return components.AttendeeReviewRequest{}
	}
	return r.Body
}

type ReviewAttendeeResponse struct {
	Headers map[string][]string
	Result  components.Attendee
}

func (r *ReviewAttendeeResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReviewAttendeeResponse) GetResult() components.Attendee {
	if r == nil {
		return components.Attendee{}
	}
	return r.Result
}
