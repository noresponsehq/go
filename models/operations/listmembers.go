package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListMembersRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (l *ListMembersRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

type ListMembersResponse struct {
	Headers map[string][]string
	Result  components.MemberList
}

func (l *ListMembersResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListMembersResponse) GetResult() components.MemberList {
	if l == nil {
		return components.MemberList{}
	}
	return l.Result
}
