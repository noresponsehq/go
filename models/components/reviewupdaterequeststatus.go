package components

import (
	"encoding/json"
	"fmt"
)

type ReviewUpdateRequestStatus string

const (
	ReviewUpdateRequestStatusOpen       ReviewUpdateRequestStatus = "open"
	ReviewUpdateRequestStatusInProgress ReviewUpdateRequestStatus = "in_progress"
)

func (e ReviewUpdateRequestStatus) ToPointer() *ReviewUpdateRequestStatus {
	return &e
}
func (e *ReviewUpdateRequestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "in_progress":
		*e = ReviewUpdateRequestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReviewUpdateRequestStatus: %v", v)
	}
}
