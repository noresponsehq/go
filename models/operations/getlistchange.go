package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type GetListChangeRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	ChangeID    string                    `pathParam:"style=simple,explode=false,name=change_id"`
}

func (g *GetListChangeRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetListChangeRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetListChangeRequest) GetChangeID() string {
	if g == nil {
		return ""
	}
	return g.ChangeID
}

type GetListChangeResponse struct {
	Headers map[string][]string
	Result  components.ListChange
}

func (g *GetListChangeResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetListChangeResponse) GetResult() components.ListChange {
	if g == nil {
		return components.ListChange{}
	}
	return g.Result
}
