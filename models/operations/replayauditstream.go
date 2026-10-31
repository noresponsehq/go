package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ReplayAuditStreamRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                              `pathParam:"style=simple,explode=false,name=workspace_id"`
	Environment components.Environment              `pathParam:"style=simple,explode=false,name=environment"`
	EndpointID  string                              `pathParam:"style=simple,explode=false,name=endpoint_id"`
	Body        components.AuditStreamReplayRequest `request:"mediaType=application/json"`
}

func (r *ReplayAuditStreamRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *ReplayAuditStreamRequest) GetEnvironment() components.Environment {
	if r == nil {
		return components.Environment("")
	}
	return r.Environment
}

func (r *ReplayAuditStreamRequest) GetEndpointID() string {
	if r == nil {
		return ""
	}
	return r.EndpointID
}

func (r *ReplayAuditStreamRequest) GetBody() components.AuditStreamReplayRequest {
	if r == nil {
		return components.AuditStreamReplayRequest{}
	}
	return r.Body
}

type ReplayAuditStreamResponse struct {
	Headers map[string][]string
	Result  components.AuditStreamList
}

func (r *ReplayAuditStreamResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReplayAuditStreamResponse) GetResult() components.AuditStreamList {
	if r == nil {
		return components.AuditStreamList{}
	}
	return r.Result
}
