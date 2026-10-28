package components

import (
	"encoding/json"
	"fmt"
)

// AnswerRevision - Which answer of each assessment the filters describe: latest, the latest stored revision, or first, the first answer, as served when the assessment was first answered (POST /v1/assessments/aggregate counts first answers). A cell of an aggregate drills down to POST /v1/assessments/search with the same filters, revision first and the cell's bucket as from and to.
type AnswerRevision string

const (
	AnswerRevisionLatest AnswerRevision = "latest"
	AnswerRevisionFirst  AnswerRevision = "first"
)

func (e AnswerRevision) ToPointer() *AnswerRevision {
	return &e
}
func (e *AnswerRevision) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "latest":
		fallthrough
	case "first":
		*e = AnswerRevision(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnswerRevision: %v", v)
	}
}
