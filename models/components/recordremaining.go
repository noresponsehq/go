package components

// RecordRemaining - Rows one search matched in one source that the lookup's records leave out, so no answer is cut short in silence.
type RecordRemaining struct {
	// Where a record came from.
	Source RecordSource `json:"source"`
	// The kind of searched value whose search matched the rows, as a record's matched names it.
	By string `json:"by"`
	// The search a person started that matched the rows. Searching the same field again with this source's table fetches them.
	Search *string `json:"search,omitzero"`
	// How many rows of the source that search matched beyond the records the lookup carries.
	Rows int64 `json:"rows"`
}

func (r *RecordRemaining) GetSource() RecordSource {
	if r == nil {
		return RecordSource{}
	}
	return r.Source
}

func (r *RecordRemaining) GetBy() string {
	if r == nil {
		return ""
	}
	return r.By
}

func (r *RecordRemaining) GetSearch() *string {
	if r == nil {
		return nil
	}
	return r.Search
}

func (r *RecordRemaining) GetRows() int64 {
	if r == nil {
		return 0
	}
	return r.Rows
}
