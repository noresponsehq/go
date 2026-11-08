package components

type AttendeeBatchItem struct {
	// The row's position in the request.
	Index int64  `json:"index"`
	ID    string `json:"id"`
	// unchanged: the row matched an attendee with the same details.
	Outcome AttendeeBatchItemOutcome `json:"outcome"`
}

func (a *AttendeeBatchItem) GetIndex() int64 {
	if a == nil {
		return 0
	}
	return a.Index
}

func (a *AttendeeBatchItem) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AttendeeBatchItem) GetOutcome() AttendeeBatchItemOutcome {
	if a == nil {
		return AttendeeBatchItemOutcome("")
	}
	return a.Outcome
}
