package components

// InvestigationsList - List investigation history.
type InvestigationsList struct {
	Items      []Investigation `json:"items"`
	HasMore    bool            `json:"has_more"`
	NextCursor *string         `json:"next_cursor"`
}

func (i *InvestigationsList) GetItems() []Investigation {
	if i == nil {
		return []Investigation{}
	}
	return i.Items
}

func (i *InvestigationsList) GetHasMore() bool {
	if i == nil {
		return false
	}
	return i.HasMore
}

func (i *InvestigationsList) GetNextCursor() *string {
	if i == nil {
		return nil
	}
	return i.NextCursor
}
