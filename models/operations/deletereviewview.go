package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type DeleteReviewViewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment          `pathParam:"style=simple,explode=false,name=environment"`
	ViewID      string                             `pathParam:"style=simple,explode=false,name=view_id"`
	Body        components.ReviewViewDeleteRequest `request:"mediaType=application/json"`
}

func (d *DeleteReviewViewRequest) GetWorkspaceID() string {
	if d == nil {
		return ""
	}
	return d.WorkspaceID
}

func (d *DeleteReviewViewRequest) GetEnvironment() components.KeyEnvironment {
	if d == nil {
		return components.KeyEnvironment("")
	}
	return d.Environment
}

func (d *DeleteReviewViewRequest) GetViewID() string {
	if d == nil {
		return ""
	}
	return d.ViewID
}

func (d *DeleteReviewViewRequest) GetBody() components.ReviewViewDeleteRequest {
	if d == nil {
		return components.ReviewViewDeleteRequest{}
	}
	return d.Body
}

type DeleteReviewViewResponse struct {
	Headers map[string][]string
}

func (d *DeleteReviewViewResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}
