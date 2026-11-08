package components

import (
	"encoding/json"
	"fmt"
)

// ScreeningCompletedStatus - completed reached every attendee; partial left some not fully covered (partial_coverage); failed covered nobody.
type ScreeningCompletedStatus string

const (
	ScreeningCompletedStatusCompleted ScreeningCompletedStatus = "completed"
	ScreeningCompletedStatusPartial   ScreeningCompletedStatus = "partial"
	ScreeningCompletedStatusFailed    ScreeningCompletedStatus = "failed"
)

func (e ScreeningCompletedStatus) ToPointer() *ScreeningCompletedStatus {
	return &e
}
func (e *ScreeningCompletedStatus) UnmarshalJSON(data []byte) error {
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
		*e = ScreeningCompletedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScreeningCompletedStatus: %v", v)
	}
}
