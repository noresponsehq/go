package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type GetScreeningRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
	ScreeningID string                    `pathParam:"style=simple,explode=false,name=screening_id"`
}

func (g *GetScreeningRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetScreeningRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetScreeningRequest) GetGroupID() string {
	if g == nil {
		return ""
	}
	return g.GroupID
}

func (g *GetScreeningRequest) GetScreeningID() string {
	if g == nil {
		return ""
	}
	return g.ScreeningID
}

type GetScreeningResponse struct {
	Headers map[string][]string
	Result  components.Screening
}

func (g *GetScreeningResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetScreeningResponse) GetResult() components.Screening {
	if g == nil {
		return components.Screening{}
	}
	return g.Result
}
