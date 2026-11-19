package components

type Between string

const (
	BetweenEmail Between = "email"
	BetweenPhone Between = "phone"
)

func (e Between) ToPointer() *Between {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Between) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone":
			return true
		}
	}
	return false
}

// RecordConflict - Two identifiers you gave that lead to records of different people.
type RecordConflict struct {
	// anchors_disagree: records found by two of the identifiers you gave clash with each other, so those identifiers may belong to different people.
	Signal RecordConflictSignal `json:"signal"`
	// The kind of value that clashes: person_name or born.
	Kind string `json:"kind"`
	// The two identifiers you gave whose records disagree.
	Between []Between `json:"between"`
	// The ids of the two records that clash, in the order of between.
	Records []string `json:"records"`
}

func (r *RecordConflict) GetSignal() RecordConflictSignal {
	if r == nil {
		return RecordConflictSignal("")
	}
	return r.Signal
}

func (r *RecordConflict) GetKind() string {
	if r == nil {
		return ""
	}
	return r.Kind
}

func (r *RecordConflict) GetBetween() []Between {
	if r == nil {
		return []Between{}
	}
	return r.Between
}

func (r *RecordConflict) GetRecords() []string {
	if r == nil {
		return []string{}
	}
	return r.Records
}
