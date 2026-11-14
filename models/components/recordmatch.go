package components

// RecordMatch - One searched value that found the record.
type RecordMatch struct {
	// The kind of searched value: email, phone, username, ip, name, password, hash, or handle for a profile read of a handle.
	By string `json:"by"`
	// The field the searched value hit, when the record holds it.
	Field *string `json:"field,omitzero"`
	// The search a person started on a value of another record that found this one, among the person report's searches, which names the value and the record and field it came from. A record found only by such searches is a lead, whatever kind of value found it. Absent for a match of the lookups.
	Search *string `json:"search,omitzero"`
}

func (r *RecordMatch) GetBy() string {
	if r == nil {
		return ""
	}
	return r.By
}

func (r *RecordMatch) GetField() *string {
	if r == nil {
		return nil
	}
	return r.Field
}

func (r *RecordMatch) GetSearch() *string {
	if r == nil {
		return nil
	}
	return r.Search
}
