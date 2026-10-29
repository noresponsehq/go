package components

type ReviewSubjectKind string

const (
	ReviewSubjectKindAccount          ReviewSubjectKind = "account"
	ReviewSubjectKindAssessment       ReviewSubjectKind = "assessment"
	ReviewSubjectKindIssuedCredential ReviewSubjectKind = "issued_credential"
)

func (e ReviewSubjectKind) ToPointer() *ReviewSubjectKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ReviewSubjectKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "account", "assessment", "issued_credential":
			return true
		}
	}
	return false
}
