package operations

import (
	"github.com/noresponsehq/go/v15/models/components"
)

type GetWorkspacePersonRequest struct {
	PersonID string `pathParam:"style=simple,explode=false,name=person_id"`
}

func (g *GetWorkspacePersonRequest) GetPersonID() string {
	if g == nil {
		return ""
	}
	return g.PersonID
}

type GetWorkspacePersonResponse struct {
	Headers map[string][]string
	Result  components.WorkspacePersonReport
}

func (g *GetWorkspacePersonResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetWorkspacePersonResponse) GetResult() components.WorkspacePersonReport {
	if g == nil {
		return components.WorkspacePersonReport{}
	}
	return g.Result
}
