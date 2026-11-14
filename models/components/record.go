package components

// Record - One answer from one source: where it came from, what found it and every field it holds.
type Record struct {
	// Made from the source and the fields' names and original values, so the same answer found twice is one record with the same id.
	ID string `json:"id"`
	// Where a record came from.
	Source RecordSource `json:"source"`
	// Each kind of searched value that found the record.
	Matched []RecordMatch `json:"matched"`
	// Every field the source returned, in the order the source's connector lists them.
	Fields []RecordField `json:"fields"`
	// When the source last gave this answer.
	ObservedAt string `json:"observed_at"`
}

func (r *Record) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *Record) GetSource() RecordSource {
	if r == nil {
		return RecordSource{}
	}
	return r.Source
}

func (r *Record) GetMatched() []RecordMatch {
	if r == nil {
		return []RecordMatch{}
	}
	return r.Matched
}

func (r *Record) GetFields() []RecordField {
	if r == nil {
		return []RecordField{}
	}
	return r.Fields
}

func (r *Record) GetObservedAt() string {
	if r == nil {
		return ""
	}
	return r.ObservedAt
}
