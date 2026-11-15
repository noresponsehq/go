package operations

import (
	"github.com/noresponsehq/go/v13/models/components"
)

type RunSearchRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment         `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                            `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID    string                            `pathParam:"style=simple,explode=false,name=person_id"`
	SearchID    string                            `pathParam:"style=simple,explode=false,name=search_id"`
	Body        components.ScreeningCancelRequest `request:"mediaType=application/json"`
}

func (r *RunSearchRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RunSearchRequest) GetEnvironment() components.KeyEnvironment {
	if r == nil {
		return components.KeyEnvironment("")
	}
	return r.Environment
}

func (r *RunSearchRequest) GetGroupID() string {
	if r == nil {
		return ""
	}
	return r.GroupID
}

func (r *RunSearchRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

func (r *RunSearchRequest) GetSearchID() string {
	if r == nil {
		return ""
	}
	return r.SearchID
}

func (r *RunSearchRequest) GetBody() components.ScreeningCancelRequest {
	if r == nil {
		return components.ScreeningCancelRequest{}
	}
	return r.Body
}

type RunSearchResponse struct {
	Headers map[string][]string
	Result  components.Search
}

func (r *RunSearchResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RunSearchResponse) GetResult() components.Search {
	if r == nil {
		return components.Search{}
	}
	return r.Result
}
