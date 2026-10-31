package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type AddReviewSubjectRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment       `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                          `pathParam:"style=simple,explode=false,name=review_id"`
	Body        components.ReviewSubjectRequest `request:"mediaType=application/json"`
}

func (a *AddReviewSubjectRequest) GetWorkspaceID() string {
	if a == nil {
		return ""
	}
	return a.WorkspaceID
}

func (a *AddReviewSubjectRequest) GetEnvironment() components.KeyEnvironment {
	if a == nil {
		return components.KeyEnvironment("")
	}
	return a.Environment
}

func (a *AddReviewSubjectRequest) GetReviewID() string {
	if a == nil {
		return ""
	}
	return a.ReviewID
}

func (a *AddReviewSubjectRequest) GetBody() components.ReviewSubjectRequest {
	if a == nil {
		return components.ReviewSubjectRequest{}
	}
	return a.Body
}

type AddReviewSubjectResponse struct {
	Headers map[string][]string
	Result  components.ReviewSubjectChange
}

func (a *AddReviewSubjectResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AddReviewSubjectResponse) GetResult() components.ReviewSubjectChange {
	if a == nil {
		return components.ReviewSubjectChange{}
	}
	return a.Result
}
