package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type SummarizeWorkspaceOperationsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                                       `pathParam:"style=simple,explode=false,name=workspace_id"`
	Body        components.WorkspaceOperationsSummaryRequest `request:"mediaType=application/json"`
}

func (s *SummarizeWorkspaceOperationsRequest) GetWorkspaceID() string {
	if s == nil {
		return ""
	}
	return s.WorkspaceID
}

func (s *SummarizeWorkspaceOperationsRequest) GetBody() components.WorkspaceOperationsSummaryRequest {
	if s == nil {
		return components.WorkspaceOperationsSummaryRequest{}
	}
	return s.Body
}

type SummarizeWorkspaceOperationsResponse struct {
	Headers map[string][]string
	Result  components.WorkspaceOperationsSummary
}

func (s *SummarizeWorkspaceOperationsResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *SummarizeWorkspaceOperationsResponse) GetResult() components.WorkspaceOperationsSummary {
	if s == nil {
		return components.WorkspaceOperationsSummary{}
	}
	return s.Result
}
