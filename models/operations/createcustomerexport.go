package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type CreateCustomerExportRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID    string                           `pathParam:"style=simple,explode=false,name=workspace_id"`
	IdempotencyKey *string                          `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.CustomerExportRequest `request:"mediaType=application/json"`
}

func (c *CreateCustomerExportRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateCustomerExportRequest) GetIdempotencyKey() *string {
	if c == nil {
		return nil
	}
	return c.IdempotencyKey
}

func (c *CreateCustomerExportRequest) GetBody() components.CustomerExportRequest {
	if c == nil {
		return components.CustomerExportRequest{}
	}
	return c.Body
}

type CreateCustomerExportResponse struct {
	Headers map[string][]string
	Result  components.CustomerExport
}

func (c *CreateCustomerExportResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateCustomerExportResponse) GetResult() components.CustomerExport {
	if c == nil {
		return components.CustomerExport{}
	}
	return c.Result
}
