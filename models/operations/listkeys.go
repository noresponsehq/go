package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListKeysRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (l *ListKeysRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

type ListKeysResponse struct {
	Headers map[string][]string
	Result  components.KeyList
}

func (l *ListKeysResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListKeysResponse) GetResult() components.KeyList {
	if l == nil {
		return components.KeyList{}
	}
	return l.Result
}
