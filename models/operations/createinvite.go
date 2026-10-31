package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type CreateInviteRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string                   `pathParam:"style=simple,explode=false,name=workspace_id"`
	Body        components.InviteRequest `request:"mediaType=application/json"`
}

func (c *CreateInviteRequest) GetWorkspaceID() string {
	if c == nil {
		return ""
	}
	return c.WorkspaceID
}

func (c *CreateInviteRequest) GetBody() components.InviteRequest {
	if c == nil {
		return components.InviteRequest{}
	}
	return c.Body
}

type CreateInviteResponse struct {
	Headers map[string][]string
	Result  components.Invite
}

func (c *CreateInviteResponse) GetHeaders() map[string][]string {
	if c == nil {
		return map[string][]string{}
	}
	return c.Headers
}

func (c *CreateInviteResponse) GetResult() components.Invite {
	if c == nil {
		return components.Invite{}
	}
	return c.Result
}
