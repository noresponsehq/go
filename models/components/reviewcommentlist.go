package components

// ReviewCommentList - Read retained review comments.
type ReviewCommentList struct {
	Items      []ReviewComment `json:"items"`
	NextCursor *string         `json:"next_cursor,omitzero"`
}

func (r *ReviewCommentList) GetItems() []ReviewComment {
	if r == nil {
		return []ReviewComment{}
	}
	return r.Items
}

func (r *ReviewCommentList) GetNextCursor() *string {
	if r == nil {
		return nil
	}
	return r.NextCursor
}
