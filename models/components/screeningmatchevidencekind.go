package components

// ScreeningMatchEvidenceKind - What matched. email, phone, supplied_handle and source_fact are hard identifiers; company, city and school are soft. supplied_handle is a handle on the attendee row; source_fact is an account a profile found by the attendee's lookups lists.
type ScreeningMatchEvidenceKind string

const (
	ScreeningMatchEvidenceKindEmail          ScreeningMatchEvidenceKind = "email"
	ScreeningMatchEvidenceKindPhone          ScreeningMatchEvidenceKind = "phone"
	ScreeningMatchEvidenceKindSuppliedHandle ScreeningMatchEvidenceKind = "supplied_handle"
	ScreeningMatchEvidenceKindSourceFact     ScreeningMatchEvidenceKind = "source_fact"
	ScreeningMatchEvidenceKindName           ScreeningMatchEvidenceKind = "name"
	ScreeningMatchEvidenceKindCompany        ScreeningMatchEvidenceKind = "company"
	ScreeningMatchEvidenceKindCity           ScreeningMatchEvidenceKind = "city"
	ScreeningMatchEvidenceKindSchool         ScreeningMatchEvidenceKind = "school"
)

func (e ScreeningMatchEvidenceKind) ToPointer() *ScreeningMatchEvidenceKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningMatchEvidenceKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone", "supplied_handle", "source_fact", "name", "company", "city", "school":
			return true
		}
	}
	return false
}
