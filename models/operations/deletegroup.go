package operations

type DeleteGroupRequest struct {
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
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
