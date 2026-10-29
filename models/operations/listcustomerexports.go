package operations

import (
	"github.com/noresponsehq/go/v9/internal/utils"
	"github.com/noresponsehq/go/v9/models/components"
)

type ListCustomerExportsRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                               `pathParam:"style=simple,explode=false,name=workspace_id"`
	Kind        components.CustomerExportRequestKind `queryParam:"style=form,explode=true,name=kind"`
	Environment components.Environment               `queryParam:"style=form,explode=true,name=environment"`
	Limit       *int64                               `default:"25" queryParam:"style=form,explode=true,name=limit"`
	Cursor      *string                              `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListCustomerExportsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCustomerExportsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListCustomerExportsRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListCustomerExportsRequest) GetKind() components.CustomerExportRequestKind {
	if l == nil {
		return components.CustomerExportRequestKind("")
	}
	return l.Kind
}

func (l *ListCustomerExportsRequest) GetEnvironment() components.Environment {
	if l == nil {
		return components.Environment("")
	}
	return l.Environment
}

func (l *ListCustomerExportsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListCustomerExportsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

type ListCustomerExportsResponse struct {
	Headers map[string][]string
	Result  components.CustomerExportList

	Next func() (*ListCustomerExportsResponse, error)
}

func (l *ListCustomerExportsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListCustomerExportsResponse) GetResult() components.CustomerExportList {
	if l == nil {
		return components.CustomerExportList{}
	}
	return l.Result
}
