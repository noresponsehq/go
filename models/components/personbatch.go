package components

// PersonBatch - What the batch did, row by row.
type PersonBatch struct {
	Created int64             `json:"created"`
	Updated int64             `json:"updated"`
	Items   []PersonBatchItem `json:"items"`
}

func (p *PersonBatch) GetCreated() int64 {
	if p == nil {
		return 0
	}
	return p.Created
}

func (p *PersonBatch) GetUpdated() int64 {
	if p == nil {
		return 0
	}
	return p.Updated
}

func (p *PersonBatch) GetItems() []PersonBatchItem {
	if p == nil {
		return []PersonBatchItem{}
	}
	return p.Items
}
