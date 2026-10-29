package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CreateReviewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment      `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.ReviewCreateRequest `request:"mediaType=application/json"`
}

func (c *CreateReviewRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateReviewRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CreateReviewRequest) GetBody() components.ReviewCreateRequest {
	if c == nil {
		return components.ReviewCreateRequest{}
	}
	return c.Body
}

type CreateReviewResponse struct {
	Headers map[string][]string
	Result  components.Review
}

func (c *CreateReviewResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateReviewResponse) GetResult() components.Review {
	if c == nil {
		return components.Review{}
	}
	return c.Result
}
