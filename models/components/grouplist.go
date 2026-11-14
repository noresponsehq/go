package components

// GroupList - Groups, newest first.
type GroupList struct {
	Items []Group `json:"items"`
	// Pass as before to read the next, older page.
	NextBefore *string `json:"next_before,omitzero"`
}

func (g *GroupList) GetItems() []Group {
	if g == nil {
		return []Group{}
	}
	return g.Items
}

func (g *GroupList) GetNextBefore() *string {
	if g == nil {
		return nil
	}
	return g.NextBefore
}
