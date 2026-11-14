package components

type PersonBatchItem struct {
	// The row's position in the request.
	Index int64  `json:"index"`
	ID    string `json:"id"`
	// unchanged: the row matched a person with the same details.
	Outcome PersonBatchItemOutcome `json:"outcome"`
}

func (p *PersonBatchItem) GetIndex() int64 {
	if p == nil {
		return 0
	}
	return p.Index
}

func (p *PersonBatchItem) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PersonBatchItem) GetOutcome() PersonBatchItemOutcome {
	if p == nil {
		return PersonBatchItemOutcome("")
	}
	return p.Outcome
}
