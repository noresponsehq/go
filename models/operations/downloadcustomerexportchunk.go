package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type DownloadCustomerExportChunkRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	ExportID    string `pathParam:"style=simple,explode=false,name=export_id"`
	ChunkIndex  int64  `pathParam:"style=simple,explode=false,name=chunk_index"`
}

func (d *DownloadCustomerExportChunkRequest) GetWorkspaceID() string {
	if d == nil {
		return ""
	}
	return d.WorkspaceID
}

func (d *DownloadCustomerExportChunkRequest) GetExportID() string {
	if d == nil {
		return ""
	}
	return d.ExportID
}

func (d *DownloadCustomerExportChunkRequest) GetChunkIndex() int64 {
	if d == nil {
		return 0
	}
	return d.ChunkIndex
}

type DownloadCustomerExportChunkResponse struct {
	Headers map[string][]string
	Result  components.CustomerExportChunk
}

func (d *DownloadCustomerExportChunkResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}

func (d *DownloadCustomerExportChunkResponse) GetResult() components.CustomerExportChunk {
	if d == nil {
		return components.CustomerExportChunk{}
	}
	return d.Result
}
