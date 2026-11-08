package components

import (
	"encoding/json"
	"fmt"
)

// AttendeeFlaggedMatchState - How sure the screening is that the accounts it read are the attendee's. A finding on a possible match is at most moderate, on a probable match at most high.
type AttendeeFlaggedMatchState string

const (
	AttendeeFlaggedMatchStateConfirmed    AttendeeFlaggedMatchState = "confirmed"
	AttendeeFlaggedMatchStateProbable     AttendeeFlaggedMatchState = "probable"
	AttendeeFlaggedMatchStatePossible     AttendeeFlaggedMatchState = "possible"
	AttendeeFlaggedMatchStateNoneFound    AttendeeFlaggedMatchState = "none_found"
	AttendeeFlaggedMatchStateInaccessible AttendeeFlaggedMatchState = "inaccessible"
)

func (e AttendeeFlaggedMatchState) ToPointer() *AttendeeFlaggedMatchState {
	return &e
}
func (e *AttendeeFlaggedMatchState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "confirmed":
		fallthrough
	case "probable":
		fallthrough
	case "possible":
		fallthrough
	case "none_found":
		fallthrough
	case "inaccessible":
		*e = AttendeeFlaggedMatchState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttendeeFlaggedMatchState: %v", v)
	}
}
