package components

// RecordField - One value a source returned.
type RecordField struct {
	// The source's own name for the field. A nested value's name is dotted, and a list element's carries its index.
	Field string `json:"field"`
	// The value as the source wrote it. A count is written in decimal and a time in RFC 3339; a breach column's value is as the leak holds it.
	Original string `json:"original"`
	// What sort of value a field holds, which decides how it is cleaned, compared and searched again. It comes from the value itself where its shape is certain (an email, an IP, a URL, an international phone number), else from the answer stored for the source's column, else text. A value of any kind is returned.
	Kind *RecordKind `json:"kind,omitzero"`
	// Which common field a value is, where many sources share it: a join date and a birth date are both dates, with the roles joined and born. A platform account's field and a breach row's column with the same meaning carry the same kind and role. A linked_account field is named by its network. Absent for a field only its source returns.
	Role *RecordRole `json:"role,omitzero"`
	// The value cleaned to its kind's one form: an email in lower case, a phone number in international form when its country is known, a date in ISO form, a country as its two-letter code, a flag as true or false. Absent when the value cannot be cleaned as its kind; original still holds it.
	Value *string `json:"value,omitzero"`
	// What the value shows of itself beyond its kind: the likely algorithm of a hash, the likely coin of a crypto address or the currency of an amount.
	Detail *string `json:"detail,omitzero"`
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

func (r *RecordField) GetKind() *RecordKind {
	if r == nil {
		return nil
	}
	return r.Kind
}

func (r *RecordField) GetRole() *RecordRole {
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

func (r *RecordField) GetDetail() *string {
	if r == nil {
		return nil
	}
	return r.Detail
}
