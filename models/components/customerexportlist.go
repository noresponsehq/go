package components

// CustomerExportList - List exports for one collection.
type CustomerExportList struct {
	Items      []CustomerExport `json:"items"`
	NextCursor *string          `json:"next_cursor,omitzero"`
}

func (c *CustomerExportList) GetItems() []CustomerExport {
	if c == nil {
		return []CustomerExport{}
	}
	return c.Items
}

func (c *CustomerExportList) GetNextCursor() *string {
	if c == nil {
		return nil
	}
	return c.NextCursor
}
