package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ResolveReviewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment       `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                          `pathParam:"style=simple,explode=false,name=review_id"`
	Body        components.ReviewResolveRequest `request:"mediaType=application/json"`
}

func (r *ResolveReviewRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *ResolveReviewRequest) GetEnvironment() components.KeyEnvironment {
	if r == nil {
		return components.KeyEnvironment("")
	}
	return r.Environment
}

func (r *ResolveReviewRequest) GetReviewID() string {
	if r == nil {
		return ""
	}
	return r.ReviewID
}

func (r *ResolveReviewRequest) GetBody() components.ReviewResolveRequest {
	if r == nil {
		return components.ReviewResolveRequest{}
	}
	return r.Body
}

type ResolveReviewResponse struct {
	Headers map[string][]string
	Result  components.Review
}

func (r *ResolveReviewResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ResolveReviewResponse) GetResult() components.Review {
	if r == nil {
		return components.Review{}
	}
	return r.Result
}
