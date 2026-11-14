package components

// PersonList - People in upload order.
type PersonList struct {
	Items []Person `json:"items"`
	// Pass as after to read the next page.
	NextAfter *string `json:"next_after,omitzero"`
}

func (p *PersonList) GetItems() []Person {
	if p == nil {
		return []Person{}
	}
	return p.Items
}

func (p *PersonList) GetNextAfter() *string {
	if p == nil {
		return nil
	}
	return p.NextAfter
}
