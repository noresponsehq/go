package operations

type RemoveMemberRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The member's mem_ id. A member of another workspace, or a removed one, answers 404.
	MemberID string `pathParam:"style=simple,explode=false,name=member_id"`
}

func (r *RemoveMemberRequest) GetWorkspaceID() string {
	if r == nil {
		return ""
	}
	return r.WorkspaceID
}

func (r *RemoveMemberRequest) GetMemberID() string {
	if r == nil {
		return ""
	}
	return r.MemberID
}

type RemoveMemberResponse struct {
	Headers map[string][]string
}

func (r *RemoveMemberResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}
