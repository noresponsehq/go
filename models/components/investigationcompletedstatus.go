package components

import (
	"encoding/json"
	"fmt"
)

type InvestigationCompletedStatus string

const (
	InvestigationCompletedStatusCompleted InvestigationCompletedStatus = "completed"
	InvestigationCompletedStatusPartial   InvestigationCompletedStatus = "partial"
	InvestigationCompletedStatusFailed    InvestigationCompletedStatus = "failed"
	InvestigationCompletedStatusCancelled InvestigationCompletedStatus = "cancelled"
)

func (e InvestigationCompletedStatus) ToPointer() *InvestigationCompletedStatus {
	return &e
}
func (e *InvestigationCompletedStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "completed":
		fallthrough
	case "partial":
		fallthrough
	case "failed":
		fallthrough
	case "cancelled":
		*e = InvestigationCompletedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvestigationCompletedStatus: %v", v)
	}
}
