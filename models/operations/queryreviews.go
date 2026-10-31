package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type QueryReviewsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment     `pathParam:"style=simple,explode=false,name=environment"`
	Body        components.ReviewQueryRequest `request:"mediaType=application/json"`
}

func (q *QueryReviewsRequest) GetWorkspaceID() string {
	if q == nil {
		return ""
	}
	return q.WorkspaceID
}

func (q *QueryReviewsRequest) GetEnvironment() components.KeyEnvironment {
	if q == nil {
		return components.KeyEnvironment("")
	}
	return q.Environment
}

func (q *QueryReviewsRequest) GetBody() components.ReviewQueryRequest {
	if q == nil {
		return components.ReviewQueryRequest{}
	}
	return q.Body
}

type QueryReviewsResponse struct {
	Headers map[string][]string
	Result  components.ReviewList
}

func (q *QueryReviewsResponse) GetHeaders() map[string][]string {
	if q == nil {
		return map[string][]string{}
	}
	return q.Headers
}

func (q *QueryReviewsResponse) GetResult() components.ReviewList {
	if q == nil {
		return components.ReviewList{}
	}
	return q.Result
}
