package components

// ReviewViewList - List your private and shared review filters.
type ReviewViewList struct {
	Items []ReviewView `json:"items"`
}

func (r *ReviewViewList) GetItems() []ReviewView {
	if r == nil {
		return []ReviewView{}
	}
	return r.Items
}
