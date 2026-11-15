package components

// WorkspacePeopleList - List workspace people.
type WorkspacePeopleList struct {
	Items      []WorkspacePerson `json:"items"`
	HasMore    bool              `json:"has_more"`
	NextCursor *string           `json:"next_cursor"`
}

func (w *WorkspacePeopleList) GetItems() []WorkspacePerson {
	if w == nil {
		return []WorkspacePerson{}
	}
	return w.Items
}

func (w *WorkspacePeopleList) GetHasMore() bool {
	if w == nil {
		return false
	}
	return w.HasMore
}

func (w *WorkspacePeopleList) GetNextCursor() *string {
	if w == nil {
		return nil
	}
	return w.NextCursor
}
