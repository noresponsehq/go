package components

// RecordEvidence - One item of evidence behind a record's likelihood.
type RecordEvidence struct {
	// found_by: what found the record, which sets where its likelihood starts. agrees: a value the record shares with what the investigator gave or with a joined record. clashes: a birth date, full name, gender or country that differs from theirs.
	Signal RecordEvidenceSignal `json:"signal"`
	// For found_by, what the search looked up (email, phone, username, handle, name). Otherwise the kind of value compared, such as email, username, person_name, born, city or crypto_address.
	Kind *string `json:"kind,omitzero"`
	// The record's own field that agreed or clashed, or for found_by the field the search matched.
	Field *string `json:"field,omitzero"`
	// given for what the investigator gave, or the id of the person's record it agreed or clashed with, or for found_by the record whose value the search looked up.
	With *string `json:"with,omitzero"`
	// The field on the other side: the field of the given identity or of the record named by with.
	WithField *string `json:"with_field,omitzero"`
	// The item's weight in bits, to two places: log2 of how much likelier the evidence is for the same person than for a stranger. Positive raises the likelihood and negative lowers it.
	Bits float64 `json:"bits"`
}

func (r *RecordEvidence) GetSignal() RecordEvidenceSignal {
	if r == nil {
		return RecordEvidenceSignal("")
	}
	return r.Signal
}

func (r *RecordEvidence) GetKind() *string {
	if r == nil {
		return nil
	}
	return r.Kind
}

func (r *RecordEvidence) GetField() *string {
	if r == nil {
		return nil
	}
	return r.Field
}

func (r *RecordEvidence) GetWith() *string {
	if r == nil {
		return nil
	}
	return r.With
}

func (r *RecordEvidence) GetWithField() *string {
	if r == nil {
		return nil
	}
	return r.WithField
}

func (r *RecordEvidence) GetBits() float64 {
	if r == nil {
		return 0.0
	}
	return r.Bits
}
