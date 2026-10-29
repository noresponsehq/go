package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type UpdateReviewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment      `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                         `pathParam:"style=simple,explode=false,name=review_id"`
	Body        components.ReviewUpdateRequest `request:"mediaType=application/json"`
}

func (u *UpdateReviewRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateReviewRequest) GetEnvironment() components.KeyEnvironment {
	if u == nil {
		return components.KeyEnvironment("")
	}
	return u.Environment
}

func (u *UpdateReviewRequest) GetReviewID() string {
	if u == nil {
		return ""
	}
	return u.ReviewID
}

func (u *UpdateReviewRequest) GetBody() components.ReviewUpdateRequest {
	if u == nil {
		return components.ReviewUpdateRequest{}
	}
	return u.Body
}

type UpdateReviewResponse struct {
	Headers map[string][]string
	Result  components.Review
}

func (u *UpdateReviewResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateReviewResponse) GetResult() components.Review {
	if u == nil {
		return components.Review{}
	}
	return u.Result
}
