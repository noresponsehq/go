package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type GetPersonRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID    string                    `pathParam:"style=simple,explode=false,name=person_id"`
}

func (g *GetPersonRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetPersonRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetPersonRequest) GetGroupID() string {
	if g == nil {
		return ""
	}
	return g.GroupID
}

func (g *GetPersonRequest) GetPersonID() string {
	if g == nil {
		return ""
	}
	return g.PersonID
}

type GetPersonResponse struct {
	Headers map[string][]string
	Result  components.PersonReport
}

func (g *GetPersonResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetPersonResponse) GetResult() components.PersonReport {
	if g == nil {
		return components.PersonReport{}
	}
	return g.Result
}
