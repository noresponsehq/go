package components

type ResultStatus string

const (
	ResultStatusNoChange               ResultStatus = "no_change"
	ResultStatusProposalRequiresReview ResultStatus = "proposal_requires_review"
)

func (e ResultStatus) ToPointer() *ResultStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ResultStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "no_change", "proposal_requires_review":
			return true
		}
	}
	return false
}
