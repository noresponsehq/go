package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type UpdateReviewViewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment    `pathParam:"style=simple,explode=false,name=environment"`
	ViewID      string                       `pathParam:"style=simple,explode=false,name=view_id"`
	Body        components.ReviewViewRequest `request:"mediaType=application/json"`
}

func (u *UpdateReviewViewRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateReviewViewRequest) GetEnvironment() components.KeyEnvironment {
	if u == nil {
		return components.KeyEnvironment("")
	}
	return u.Environment
}

func (u *UpdateReviewViewRequest) GetViewID() string {
	if u == nil {
		return ""
	}
	return u.ViewID
}

func (u *UpdateReviewViewRequest) GetBody() components.ReviewViewRequest {
	if u == nil {
		return components.ReviewViewRequest{}
	}
	return u.Body
}

type UpdateReviewViewResponse struct {
	Headers map[string][]string
	Result  components.ReviewView
}

func (u *UpdateReviewViewResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateReviewViewResponse) GetResult() components.ReviewView {
	if u == nil {
		return components.ReviewView{}
	}
	return u.Result
}
