package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type ReviewPersonRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment      `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                         `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID    string                         `pathParam:"style=simple,explode=false,name=person_id"`
	Body        components.PersonReviewRequest `request:"mediaType=application/json"`
}

func (r *ReviewPersonRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *ReviewPersonRequest) GetEnvironment() components.KeyEnvironment {
	if r == nil {
		return components.KeyEnvironment("")
	}
	return r.Environment
}

func (r *ReviewPersonRequest) GetGroupID() string {
	if r == nil {
		return ""
	}
	return r.GroupID
}

func (r *ReviewPersonRequest) GetPersonID() string {
	if r == nil {
		return ""
	}
	return r.PersonID
}

func (r *ReviewPersonRequest) GetBody() components.PersonReviewRequest {
	if r == nil {
		return components.PersonReviewRequest{}
	}
	return r.Body
}

type ReviewPersonResponse struct {
	Headers map[string][]string
	Result  components.Person
}

func (r *ReviewPersonResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReviewPersonResponse) GetResult() components.Person {
	if r == nil {
		return components.Person{}
	}
	return r.Result
}
