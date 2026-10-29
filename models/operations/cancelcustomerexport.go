package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type CancelCustomerExportRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                                 `pathParam:"style=simple,explode=false,name=workspace_id"`
	ExportID    string                                 `pathParam:"style=simple,explode=false,name=export_id"`
	Body        components.CustomerExportActionRequest `request:"mediaType=application/json"`
}

func (c *CancelCustomerExportRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CancelCustomerExportRequest) GetExportID() string {
	if c == nil {
		return ""
	}
	return c.ExportID
}

func (c *CancelCustomerExportRequest) GetBody() components.CustomerExportActionRequest {
	if c == nil {
		return components.CustomerExportActionRequest{}
	}
	return c.Body
}

type CancelCustomerExportResponse struct {
	Headers map[string][]string
	Result  components.CustomerExport
}

func (c *CancelCustomerExportResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CancelCustomerExportResponse) GetResult() components.CustomerExport {
	if c == nil {
		return components.CustomerExport{}
	}
	return c.Result
}
