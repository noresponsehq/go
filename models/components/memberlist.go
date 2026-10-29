package components

// MemberList - GET /v1/workspaces/{workspace_id}/members: the workspace's active members, oldest first, at most 100.
type MemberList struct {
	Items []Member `json:"items"`
}

func (m *MemberList) GetItems() []Member {
	if m == nil {
		return []Member{}
	}
	return m.Items
}
