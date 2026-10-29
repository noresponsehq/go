package components

type Conclusion string

const (
	ConclusionBenign       Conclusion = "benign"
	ConclusionAbuse        Conclusion = "abuse"
	ConclusionInconclusive Conclusion = "inconclusive"
)

func (e Conclusion) ToPointer() *Conclusion {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Conclusion) IsExact() bool {
	if e != nil {
		switch *e {
		case "benign", "abuse", "inconclusive":
			return true
		}
	}
	return false
}
