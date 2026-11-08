package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type UpdateEventRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	EventID     string                    `pathParam:"style=simple,explode=false,name=event_id"`
	Body        components.EventUpdate    `request:"mediaType=application/json"`
}

func (u *UpdateEventRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateEventRequest) GetEnvironment() components.KeyEnvironment {
	if u == nil {
		return components.KeyEnvironment("")
	}
	return u.Environment
}

func (u *UpdateEventRequest) GetEventID() string {
	if u == nil {
		return ""
	}
	return u.EventID
}

func (u *UpdateEventRequest) GetBody() components.EventUpdate {
	if u == nil {
		return components.EventUpdate{}
	}
	return u.Body
}

type UpdateEventResponse struct {
	Headers map[string][]string
	Result  components.Event
}

func (u *UpdateEventResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateEventResponse) GetResult() components.Event {
	if u == nil {
		return components.Event{}
	}
	return u.Result
}
