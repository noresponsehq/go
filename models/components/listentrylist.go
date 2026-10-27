package components

// ListEntryList - GET /v1/workspaces/{workspace_id}/lists/{environment}: every entry of the environment's list, newest first. An environment holds at most 10,000 entries.
type ListEntryList struct {
	Items []ListEntry `json:"items"`
}

func (l *ListEntryList) GetItems() []ListEntry {
	if l == nil {
		return []ListEntry{}
	}
	return l.Items
}
