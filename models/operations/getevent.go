package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetEventRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	EventID     string                    `pathParam:"style=simple,explode=false,name=event_id"`
}

func (g *GetEventRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetEventRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetEventRequest) GetEventID() string {
	if g == nil {
		return ""
	}
	return g.EventID
}

type GetEventResponse struct {
	Headers map[string][]string
	Result  components.Event
}

func (g *GetEventResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetEventResponse) GetResult() components.Event {
	if g == nil {
		return components.Event{}
	}
	return g.Result
}
