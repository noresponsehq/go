package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type UpdateMemberRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The member's mem_ id. A member of another workspace, or a removed one, answers 404.
	MemberID string                         `pathParam:"style=simple,explode=false,name=member_id"`
	Body     components.MemberUpdateRequest `request:"mediaType=application/json"`
}

func (u *UpdateMemberRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateMemberRequest) GetMemberID() string {
	if u == nil {
		return ""
	}
	return u.MemberID
}

func (u *UpdateMemberRequest) GetBody() components.MemberUpdateRequest {
	if u == nil {
		return components.MemberUpdateRequest{}
	}
	return u.Body
}

type UpdateMemberResponse struct {
	Headers map[string][]string
	Result  components.Member
}

func (u *UpdateMemberResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateMemberResponse) GetResult() components.Member {
	if u == nil {
		return components.Member{}
	}
	return u.Result
}
