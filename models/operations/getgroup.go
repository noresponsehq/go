package operations

import (
	"github.com/noresponsehq/go/v20/models/components"
)

type GetGroupRequest struct {
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (g *GetGroupRequest) GetGroupID() string {
	if g == nil {
		return ""
	}
	return g.GroupID
}

type GetGroupResponse struct {
	Headers map[string][]string
	Result  components.Group
}

func (g *GetGroupResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetGroupResponse) GetResult() components.Group {
	if g == nil {
		return components.Group{}
	}
	return g.Result
}
