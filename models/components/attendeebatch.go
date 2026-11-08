package components

// AttendeeBatch - What the batch did, row by row.
type AttendeeBatch struct {
	Created int64               `json:"created"`
	Updated int64               `json:"updated"`
	Items   []AttendeeBatchItem `json:"items"`
}

func (a *AttendeeBatch) GetCreated() int64 {
	if a == nil {
		return 0
	}
	return a.Created
}

func (a *AttendeeBatch) GetUpdated() int64 {
	if a == nil {
		return 0
	}
	return a.Updated
}

func (a *AttendeeBatch) GetItems() []AttendeeBatchItem {
	if a == nil {
		return []AttendeeBatchItem{}
	}
	return a.Items
}
