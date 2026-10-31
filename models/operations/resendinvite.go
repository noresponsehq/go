package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type ResendInviteRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The invite's inv_ id. An invite of another workspace, or one no longer pending, answers 404.
	InviteID string                         `pathParam:"style=simple,explode=false,name=invite_id"`
	Body     components.InviteResendRequest `request:"mediaType=application/json"`
}

func (r *ResendInviteRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *ResendInviteRequest) GetInviteID() string {
	if r == nil {
		return ""
	}
	return r.InviteID
}

func (r *ResendInviteRequest) GetBody() components.InviteResendRequest {
	if r == nil {
		return components.InviteResendRequest{}
	}
	return r.Body
}

type ResendInviteResponse struct {
	Headers map[string][]string
	Result  components.Invite
}

func (r *ResendInviteResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ResendInviteResponse) GetResult() components.Invite {
	if r == nil {
		return components.Invite{}
	}
	return r.Result
}
