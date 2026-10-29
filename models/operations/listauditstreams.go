package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type ListAuditStreamsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                 `pathParam:"style=simple,explode=false,name=workspace_id"`
	Environment components.Environment `pathParam:"style=simple,explode=false,name=environment"`
}

func (l *ListAuditStreamsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListAuditStreamsRequest) GetEnvironment() components.Environment {
	if l == nil {
		return components.Environment("")
	}
	return l.Environment
}

type ListAuditStreamsResponse struct {
	Headers map[string][]string
	Result  components.AuditStreamList
}

func (l *ListAuditStreamsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAuditStreamsResponse) GetResult() components.AuditStreamList {
	if l == nil {
		return components.AuditStreamList{}
	}
	return l.Result
}
