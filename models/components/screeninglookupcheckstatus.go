package components

// ScreeningLookupCheckStatus - found: a source found the identifier. not_found: every source answered and none found it. partial: the sources that answered found nothing and the others did not answer. not_covered: no source answered. A lookup that failed reads partial or not_covered, not not_found.
type ScreeningLookupCheckStatus string

const (
	ScreeningLookupCheckStatusFound      ScreeningLookupCheckStatus = "found"
	ScreeningLookupCheckStatusNotFound   ScreeningLookupCheckStatus = "not_found"
	ScreeningLookupCheckStatusPartial    ScreeningLookupCheckStatus = "partial"
	ScreeningLookupCheckStatusNotCovered ScreeningLookupCheckStatus = "not_covered"
)

func (e ScreeningLookupCheckStatus) ToPointer() *ScreeningLookupCheckStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningLookupCheckStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "found", "not_found", "partial", "not_covered":
			return true
		}
	}
	return false
}
