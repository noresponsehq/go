package components

// RecordConflictSignal - anchors_disagree: records found by two of the identifiers you gave clash with each other, so those identifiers may belong to different people.
type RecordConflictSignal string

const (
	RecordConflictSignalAnchorsDisagree RecordConflictSignal = "anchors_disagree"
)

func (e RecordConflictSignal) ToPointer() *RecordConflictSignal {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RecordConflictSignal) IsExact() bool {
	if e != nil {
		switch *e {
		case "anchors_disagree":
			return true
		}
	}
	return false
}
