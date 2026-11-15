package components

import (
	"github.com/noresponsehq/go/v12/internal/utils"
)

// RecordField - One value a source returned.
type RecordField struct {
	// The source's own name for the field. A nested value's name is dotted, and a list element's carries its index.
	Field string `json:"field"`
	// The value as the source wrote it. A count is written in decimal and a time in RFC 3339; a breach column's value is as the leak holds it.
	Original string `json:"original"`
	// What sort of value a field holds, which decides how it is cleaned, compared and searched again. It comes from the value itself where its shape is certain (an email, an IP, a URL, an international phone number), else from the answer committed for the source's column, else from the answer stored for it, else text. A value of any kind is returned. Each kind's description says what it holds and the form value takes.
	Kind *RecordKind `json:"kind,omitzero"`
	// Which common field a value is, where many sources share it: a join date and a birth date are both dates, with the roles joined and born. A platform account's field and a breach row's column with the same meaning carry the same kind and role. A linked_account field is named by its network. Absent for a field only its source returns. Each role's description names the kinds its values have.
	Role *RecordRole `json:"role,omitzero"`
	// The value cleaned to its kind's one form: an email in lower case, a phone number in international form when its country is known, a date in ISO form, a country as its two-letter code, a flag as true or false. Absent when the value cannot be cleaned as its kind, or when the source masked it or cut it short, with a run of asterisks or bullets or a trailing ellipsis, in a kind that links people or is searched again; original still holds it.
	Value *string `json:"value,omitzero"`
	// What the value shows of itself beyond its kind: the likely algorithm of a hash, the likely coin of a crypto address or the currency of an amount.
	Detail *string `json:"detail,omitzero"`
	// How distinctive a searchable value is: an email, phone, username, password, hash, person name or IP. Absent for other values, for a field with no value, and when its rating could not be worked out, such as when a count failed.
	Distinctive *RecordDistinctive `json:"distinctive,omitzero"`
}

func (r RecordField) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RecordField) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
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

func (r *RecordField) GetDistinctive() *RecordDistinctive {
	if r == nil {
		return nil
	}
	return r.Distinctive
}
