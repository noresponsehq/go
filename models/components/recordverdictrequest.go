package components

// RecordVerdictRequest - A reviewer's verdict on whether one of a person's lookup records is theirs.
type RecordVerdictRequest struct {
	// One of the person's lookup records.
	RecordID string `json:"record_id"`
	// A reviewer's overrule (POST .../people/{person_id}/records), which decides joined whatever the likelihood and sticks for as long as a lookup finds the same record. Absent when nobody overruled it.
	Verdict Verdict `json:"verdict"`
}

func (r *RecordVerdictRequest) GetRecordID() string {
	if r == nil {
		return ""
	}
	return r.RecordID
}

func (r *RecordVerdictRequest) GetVerdict() Verdict {
	if r == nil {
		return Verdict("")
	}
	return r.Verdict
}
