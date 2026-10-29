package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type GetCustomerExportRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	ExportID    string `pathParam:"style=simple,explode=false,name=export_id"`
}

func (g *GetCustomerExportRequest) GetWorkspaceID() string {
	if g == nil {
		return ""
	}
	return g.WorkspaceID
}

func (g *GetCustomerExportRequest) GetExportID() string {
	if g == nil {
		return ""
	}
	return g.ExportID
}

type GetCustomerExportResponse struct {
	Headers map[string][]string
	Result  components.CustomerExport
}

func (g *GetCustomerExportResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetCustomerExportResponse) GetResult() components.CustomerExport {
	if g == nil {
		return components.CustomerExport{}
	}
	return g.Result
}
