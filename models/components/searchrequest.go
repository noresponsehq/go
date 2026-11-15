package components

// SearchRequest - The value to search again, named by the record and field that hold it rather than sent again.
type SearchRequest struct {
	// One of the person's records: a lookup's, a profile's or one an earlier search found.
	RecordID string `json:"record_id"`
	// The name of the record's field whose value to search. Its kind must be email, phone, username, person_name, password, hash, ip or domain, and it must have a value: a field the source masked or that could not be cleaned is not searched.
	Field string `json:"field"`
	// Search only this breach table, as a remaining entry names it, to fetch the rows an earlier search left out. Not for a phone.
	Table *string `json:"table,omitzero"`
}

func (s *SearchRequest) GetRecordID() string {
	if s == nil {
		return ""
	}
	return s.RecordID
}

func (s *SearchRequest) GetField() string {
	if s == nil {
		return ""
	}
	return s.Field
}

func (s *SearchRequest) GetTable() *string {
	if s == nil {
		return nil
	}
	return s.Table
}
