package components

// OtherGroups - Where else in the environment the person appears, read live from those groups: deleting one removes it here.
type OtherGroups struct {
	// How many other groups include this person, of which items lists the 20 most recent.
	Total int64 `json:"total"`
	// This person's memberships in the 20 most recent other groups, newest group first.
	Items []OtherGroup `json:"items"`
}

func (o *OtherGroups) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}

func (o *OtherGroups) GetItems() []OtherGroup {
	if o == nil {
		return []OtherGroup{}
	}
	return o.Items
}
