package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ListInvitesRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (l *ListInvitesRequest) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

type ListInvitesResponse struct {
	Headers map[string][]string
	Result  components.InviteList
}

func (l *ListInvitesResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListInvitesResponse) GetResult() components.InviteList {
	if l == nil {
		return components.InviteList{}
	}
	return l.Result
}
