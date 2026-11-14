package components

// RecordField - One value a source returned.
type RecordField struct {
	// The source's own name for the field. A nested value's name is dotted, and a list element's carries its index.
	Field string `json:"field"`
	// The value as the source wrote it. A count is written in decimal and a time in RFC 3339.
	Original string `json:"original"`
	// What sort of value the field holds. Absent until the field's kind is known.
	Kind *string `json:"kind,omitzero"`
	// Which common field it is, where many sources share it, such as handle, display_name, account_id, profile_link, picture, bio, website, linked_account (named by its network), account_status, joined, last_updated, organisation, location, followers, following or posts. Absent for a field only this source returns.
	Role *string `json:"role,omitzero"`
	// The cleaned value. Absent when it cannot be cleaned, or until the field's kind is known.
	Value *string `json:"value,omitzero"`
}

func (r *RecordField) GetField() string {
	if r == nil {
		return ""
	}
	return r.Field
}

func (r *RecordField) GetOriginal() string {
	if r == nil {
		return ""
	}
	return r.Original
}

func (r *RecordField) GetKind() *string {
	if r == nil {
		return nil
	}
	return r.Kind
}

func (r *RecordField) GetRole() *string {
	if r == nil {
		return nil
	}
	return r.Role
}

func (r *RecordField) GetValue() *string {
	if r == nil {
		return nil
	}
	return r.Value
}
