package operations

type RevokeInviteRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The invite's inv_ id. An invite of another workspace, or one no longer pending, answers 404.
	InviteID string `pathParam:"style=simple,explode=false,name=invite_id"`
}

func (r *RevokeInviteRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RevokeInviteRequest) GetInviteID() string {
	if r == nil {
		return ""
	}
	return r.InviteID
}

type RevokeInviteResponse struct {
	Headers map[string][]string
}

func (r *RevokeInviteResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}
