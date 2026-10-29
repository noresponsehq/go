package components

// InviteList - GET /v1/workspaces/{workspace_id}/invites: the workspace's pending invites, oldest first, at most 50, with no token.
type InviteList struct {
	Items []Invite `json:"items"`
}

func (i *InviteList) GetItems() []Invite {
	if i == nil {
		return []Invite{}
	}
	return i.Items
}
