package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ExportReviewRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	ReviewID    string                    `pathParam:"style=simple,explode=false,name=review_id"`
}

func (e *ExportReviewRequest) GetWorkspaceID() string {
	if e == nil {
		return ""
	}
	return e.WorkspaceID
}

func (e *ExportReviewRequest) GetEnvironment() components.KeyEnvironment {
	if e == nil {
		return components.KeyEnvironment("")
	}
	return e.Environment
}

func (e *ExportReviewRequest) GetReviewID() string {
	if e == nil {
		return ""
	}
	return e.ReviewID
}

type ExportReviewResponse struct {
	Headers map[string][]string
	Result  components.Export
}

func (e *ExportReviewResponse) GetHeaders() map[string][]string {
	if e == nil {
		return map[string][]string{}
	}
	return e.Headers
}

func (e *ExportReviewResponse) GetResult() components.Export {
	if e == nil {
		return components.Export{}
	}
	return e.Result
}
