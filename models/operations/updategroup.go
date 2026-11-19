package operations

import (
	"github.com/noresponsehq/go/v16/models/components"
)

type UpdateGroupRequest struct {
	GroupID string                 `pathParam:"style=simple,explode=false,name=group_id"`
	Body    components.GroupUpdate `request:"mediaType=application/json"`
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
