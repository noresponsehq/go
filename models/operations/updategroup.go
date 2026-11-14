package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type UpdateGroupRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
	Body        components.GroupUpdate    `request:"mediaType=application/json"`
}

func (u *UpdateGroupRequest) GetWorkspaceID() string {
	if u == nil {
		return ""
	}
	return u.WorkspaceID
}

func (u *UpdateGroupRequest) GetEnvironment() components.KeyEnvironment {
	if u == nil {
		return components.KeyEnvironment("")
	}
	return u.Environment
}

func (u *UpdateGroupRequest) GetGroupID() string {
	if u == nil {
		return ""
	}
	return u.GroupID
}

func (u *UpdateGroupRequest) GetBody() components.GroupUpdate {
	if u == nil {
		return components.GroupUpdate{}
	}
	return u.Body
}

type UpdateGroupResponse struct {
	Headers map[string][]string
	Result  components.Group
}

func (u *UpdateGroupResponse) GetHeaders() map[string][]string {
	if u == nil {
		return map[string][]string{}
	}
	return u.Headers
}

func (u *UpdateGroupResponse) GetResult() components.Group {
	if u == nil {
		return components.Group{}
	}
	return u.Result
}
