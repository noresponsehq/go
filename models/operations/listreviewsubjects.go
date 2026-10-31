package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListReviewSubjectsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                    `pathParam:"style=simple,explode=false,name=review_id"`
}

func (l *ListReviewSubjectsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListReviewSubjectsRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListReviewSubjectsRequest) GetReviewID() string {
	if l == nil {
		return ""
	}
	return l.ReviewID
}

type ListReviewSubjectsResponse struct {
	Headers map[string][]string
	Result  components.ReviewSubjectList
}

func (l *ListReviewSubjectsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListReviewSubjectsResponse) GetResult() components.ReviewSubjectList {
	if l == nil {
		return components.ReviewSubjectList{}
	}
	return l.Result
}
