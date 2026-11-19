package components

// RecordEvidenceSignal - found_by: what found the record, which sets where its likelihood starts. agrees: a value the record shares with what the investigator gave or with a joined record. clashes: a birth date, full name, gender or country that differs from theirs.
type RecordEvidenceSignal string

const (
	RecordEvidenceSignalFoundBy RecordEvidenceSignal = "found_by"
	RecordEvidenceSignalAgrees  RecordEvidenceSignal = "agrees"
	RecordEvidenceSignalClashes RecordEvidenceSignal = "clashes"
)

func (e RecordEvidenceSignal) ToPointer() *RecordEvidenceSignal {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RecordEvidenceSignal) IsExact() bool {
	if e != nil {
		switch *e {
		case "found_by", "agrees", "clashes":
			return true
		}
	}
	return false
}
