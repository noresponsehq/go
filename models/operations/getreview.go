package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetReviewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                    `pathParam:"style=simple,explode=false,name=review_id"`
}

func (g *GetReviewRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetReviewRequest) GetEnvironment() components.KeyEnvironment {
	if g == nil {
		return components.KeyEnvironment("")
	}
	return g.Environment
}

func (g *GetReviewRequest) GetReviewID() string {
	if g == nil {
		return ""
	}
	return g.ReviewID
}

type GetReviewResponse struct {
	Headers map[string][]string
	Result  components.Review
}

func (g *GetReviewResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetReviewResponse) GetResult() components.Review {
	if g == nil {
		return components.Review{}
	}
	return g.Result
}
