package components

type Agreement string

const (
	AgreementMatch    Agreement = "match"
	AgreementPartial  Agreement = "partial"
	AgreementMismatch Agreement = "mismatch"
)

func (e Agreement) ToPointer() *Agreement {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Agreement) IsExact() bool {
	if e != nil {
		switch *e {
		case "match", "partial", "mismatch":
			return true
		}
	}
	return false
}
