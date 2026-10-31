package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListReviewViewsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
}

func (l *ListReviewViewsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListReviewViewsRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

type ListReviewViewsResponse struct {
	Headers map[string][]string
	Result  components.ReviewViewList
}

func (l *ListReviewViewsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListReviewViewsResponse) GetResult() components.ReviewViewList {
	if l == nil {
		return components.ReviewViewList{}
	}
	return l.Result
}
