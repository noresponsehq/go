package operations

import (
	"github.com/noresponsehq/go/v11/models/components"
)

type DeleteGroupRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                    `pathParam:"style=simple,explode=false,name=group_id"`
}

func (d *DeleteGroupRequest) GetWorkspaceID() string {
	if d == nil {
		return ""
	}
	return d.WorkspaceID
}

func (d *DeleteGroupRequest) GetEnvironment() components.KeyEnvironment {
	if d == nil {
		return components.KeyEnvironment("")
	}
	return d.Environment
}

func (d *DeleteGroupRequest) GetGroupID() string {
	if d == nil {
		return ""
	}
	return d.GroupID
}

type DeleteGroupResponse struct {
	Headers map[string][]string
}

func (d *DeleteGroupResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}
