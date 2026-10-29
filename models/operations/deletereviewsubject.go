package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type DeleteReviewSubjectRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment          `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                             `pathParam:"style=simple,explode=false,name=review_id"`
	SubjectID   string                             `pathParam:"style=simple,explode=false,name=subject_id"`
	Body        components.ReviewViewDeleteRequest `request:"mediaType=application/json"`
}

func (d *DeleteReviewSubjectRequest) GetWorkspaceID() string {
	if d == nil {
		return ""
	}
	return d.WorkspaceID
}

func (d *DeleteReviewSubjectRequest) GetEnvironment() components.KeyEnvironment {
	if d == nil {
		return components.KeyEnvironment("")
	}
	return d.Environment
}

func (d *DeleteReviewSubjectRequest) GetReviewID() string {
	if d == nil {
		return ""
	}
	return d.ReviewID
}

func (d *DeleteReviewSubjectRequest) GetSubjectID() string {
	if d == nil {
		return ""
	}
	return d.SubjectID
}

func (d *DeleteReviewSubjectRequest) GetBody() components.ReviewViewDeleteRequest {
	if d == nil {
		return components.ReviewViewDeleteRequest{}
	}
	return d.Body
}

type DeleteReviewSubjectResponse struct {
	Headers map[string][]string
	Result  components.Review
}

func (d *DeleteReviewSubjectResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}

func (d *DeleteReviewSubjectResponse) GetResult() components.Review {
	if d == nil {
		return components.Review{}
	}
	return d.Result
}
