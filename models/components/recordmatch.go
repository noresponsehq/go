package components

// RecordMatch - One searched value that found the record.
type RecordMatch struct {
	// The kind of searched value: email, phone, username, ip, name, or handle for a profile read of a handle.
	By string `json:"by"`
	// The field the searched value hit, when the record holds it.
	Field *string `json:"field,omitzero"`
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
