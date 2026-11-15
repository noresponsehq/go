package operations

import (
	"github.com/noresponsehq/go/v13/models/components"
)

type StartSearchRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID    string                    `pathParam:"style=simple,explode=false,name=person_id"`
	Body        components.SearchRequest  `request:"mediaType=application/json"`
}

func (s *StartSearchRequest) GetWorkspaceID() string {
	if s == nil {
		return ""
	}
	return s.WorkspaceID
}

func (s *StartSearchRequest) GetEnvironment() components.KeyEnvironment {
	if s == nil {
		return components.KeyEnvironment("")
	}
	return s.Environment
}

func (s *StartSearchRequest) GetGroupID() string {
	if s == nil {
		return ""
	}
	return s.GroupID
}

func (s *StartSearchRequest) GetPersonID() string {
	if s == nil {
		return ""
	}
	return s.PersonID
}

func (s *StartSearchRequest) GetBody() components.SearchRequest {
	if s == nil {
		return components.SearchRequest{}
	}
	return s.Body
}

type StartSearchResponse struct {
	Headers map[string][]string
	Result  components.Search
}

func (s *StartSearchResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *StartSearchResponse) GetResult() components.Search {
	if s == nil {
		return components.Search{}
	}
	return s.Result
}
