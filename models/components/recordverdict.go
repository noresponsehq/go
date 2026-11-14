package components

// RecordVerdict - A recorded verdict on one record. A later verdict on the same record replaces it.
type RecordVerdict struct {
	RecordID string `json:"record_id"`
	// A reviewer's overrule (POST .../people/{person_id}/records), which decides joined whatever the likelihood and sticks for as long as a lookup finds the same record. Absent when nobody overruled it.
	Verdict   Verdict `json:"verdict"`
	DecidedBy string  `json:"decided_by"`
	DecidedAt string  `json:"decided_at"`
}

func (r *RecordVerdict) GetRecordID() string {
	if r == nil {
		return ""
	}
	return r.RecordID
}

func (r *RecordVerdict) GetVerdict() Verdict {
	if r == nil {
		return Verdict("")
	}
	return r.Verdict
}

func (r *RecordVerdict) GetDecidedBy() string {
	if r == nil {
		return ""
	}
	return r.DecidedBy
}

func (r *RecordVerdict) GetDecidedAt() string {
	if r == nil {
		return ""
	}
	return r.DecidedAt
}
