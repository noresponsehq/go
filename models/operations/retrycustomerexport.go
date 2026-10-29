package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type RetryCustomerExportRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                                 `pathParam:"style=simple,explode=false,name=workspace_id"`
	ExportID    string                                 `pathParam:"style=simple,explode=false,name=export_id"`
	Body        components.CustomerExportActionRequest `request:"mediaType=application/json"`
}

func (r *RetryCustomerExportRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RetryCustomerExportRequest) GetExportID() string {
	if r == nil {
		return ""
	}
	return r.ExportID
}

func (r *RetryCustomerExportRequest) GetBody() components.CustomerExportActionRequest {
	if r == nil {
		return components.CustomerExportActionRequest{}
	}
	return r.Body
}

type RetryCustomerExportResponse struct {
	Headers map[string][]string
	Result  components.CustomerExport
}

func (r *RetryCustomerExportResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RetryCustomerExportResponse) GetResult() components.CustomerExport {
	if r == nil {
		return components.CustomerExport{}
	}
	return r.Result
}
