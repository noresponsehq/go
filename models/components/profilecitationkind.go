package components

// ProfileCitationKind - record: a record in lookup.records or profiles[].record. post: an item in posts. check: an item in lookup.checks.
type ProfileCitationKind string

const (
	ProfileCitationKindRecord ProfileCitationKind = "record"
	ProfileCitationKindPost   ProfileCitationKind = "post"
	ProfileCitationKindCheck  ProfileCitationKind = "check"
)

func (e ProfileCitationKind) ToPointer() *ProfileCitationKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ProfileCitationKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "record", "post", "check":
			return true
		}
	}
	return false
}
