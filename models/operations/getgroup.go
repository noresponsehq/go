package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type GetGroupRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
}

func (g *GetGroupRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetGroupRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetGroupRequest) GetGroupID() string {
	if g == nil {
		return ""
	}
	return g.GroupID
}

type GetGroupResponse struct {
	Headers map[string][]string
	Result  components.Group
}

func (g *GetGroupResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetGroupResponse) GetResult() components.Group {
	if g == nil {
		return components.Group{}
	}
	return g.Result
}
