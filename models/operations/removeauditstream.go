package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type RemoveAuditStreamRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                 `pathParam:"style=simple,explode=false,name=workspace_id"`
	Environment components.Environment `pathParam:"style=simple,explode=false,name=environment"`
	EndpointID  string                 `pathParam:"style=simple,explode=false,name=endpoint_id"`
}

func (r *RemoveAuditStreamRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RemoveAuditStreamRequest) GetEnvironment() components.Environment {
	if r == nil {
		return components.Environment("")
	}
	return r.Environment
}

func (r *RemoveAuditStreamRequest) GetEndpointID() string {
	if r == nil {
		return ""
	}
	return r.EndpointID
}

type RemoveAuditStreamResponse struct {
	Headers map[string][]string
	Result  components.AuditStreamList
}

func (r *RemoveAuditStreamResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RemoveAuditStreamResponse) GetResult() components.AuditStreamList {
	if r == nil {
		return components.AuditStreamList{}
	}
	return r.Result
}
