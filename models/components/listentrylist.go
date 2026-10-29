package components

// ListEntryList - Newest first. Without a query, returns every active entry for compatibility. With limit or cursor, returns up to 100 entries and an authenticated workspace/environment-bound cursor. Expired entries are omitted.
type ListEntryList struct {
	Items      []ListEntry `json:"items"`
	NextCursor *string     `json:"next_cursor,omitzero"`
}

func (l *ListEntryList) GetItems() []ListEntry {
	if l == nil {
		return []ListEntry{}
	}
	return l.Items
}

func (l *ListEntryList) GetNextCursor() *string {
	if l == nil {
		return nil
	}
	return l.NextCursor
}
