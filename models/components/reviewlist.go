package components

// ReviewList - List retained account reviews.
type ReviewList struct {
	Items      []Review `json:"items"`
	NextCursor *string  `json:"next_cursor,omitzero"`
}

func (r *ReviewList) GetItems() []Review {
	if r == nil {
		return []Review{}
	}
	return r.Items
}

func (r *ReviewList) GetNextCursor() *string {
	if r == nil {
		return nil
	}
	return r.NextCursor
}
