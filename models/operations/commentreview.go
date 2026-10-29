package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CommentReviewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment       `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                          `pathParam:"style=simple,explode=false,name=review_id"`
	Body        components.ReviewCommentRequest `request:"mediaType=application/json"`
}

func (c *CommentReviewRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CommentReviewRequest) GetEnvironment() components.KeyEnvironment {
	if c == nil {
		return components.KeyEnvironment("")
	}
	return c.Environment
}

func (c *CommentReviewRequest) GetReviewID() string {
	if c == nil {
		return ""
	}
	return c.ReviewID
}

func (c *CommentReviewRequest) GetBody() components.ReviewCommentRequest {
	if c == nil {
		return components.ReviewCommentRequest{}
	}
	return c.Body
}

type CommentReviewResponse struct {
	Headers map[string][]string
	Result  components.ReviewComment
}

func (c *CommentReviewResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CommentReviewResponse) GetResult() components.ReviewComment {
	if c == nil {
		return components.ReviewComment{}
	}
	return c.Result
}
