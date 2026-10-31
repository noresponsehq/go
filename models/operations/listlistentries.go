package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListListEntriesRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	Limit       *int64                    `queryParam:"style=form,explode=true,name=limit"`
	Cursor      *string                   `queryParam:"style=form,explode=true,name=cursor"`
}

func (l *ListListEntriesRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListListEntriesRequest) GetEnvironment() components.KeyEnvironment {
	if l == nil {
		return components.KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListListEntriesRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListListEntriesRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

type ListListEntriesResponse struct {
	Headers map[string][]string
	Result  components.ListEntryList

	Next func() (*ListListEntriesResponse, error)
}

func (l *ListListEntriesResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListListEntriesResponse) GetResult() components.ListEntryList {
	if l == nil {
		return components.ListEntryList{}
	}
	return l.Result
}
