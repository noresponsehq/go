package components

import (
	"encoding/json"
	"fmt"
)

// PersonFlaggedMatchState - How sure the screening is that the accounts it read are the person's. A finding on a possible match is at most moderate, on a probable match at most high.
type PersonFlaggedMatchState string

const (
	PersonFlaggedMatchStateConfirmed    PersonFlaggedMatchState = "confirmed"
	PersonFlaggedMatchStateProbable     PersonFlaggedMatchState = "probable"
	PersonFlaggedMatchStatePossible     PersonFlaggedMatchState = "possible"
	PersonFlaggedMatchStateNoneFound    PersonFlaggedMatchState = "none_found"
	PersonFlaggedMatchStateInaccessible PersonFlaggedMatchState = "inaccessible"
)

func (e PersonFlaggedMatchState) ToPointer() *PersonFlaggedMatchState {
	return &e
}
func (e *PersonFlaggedMatchState) UnmarshalJSON(data []byte) error {
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
		*e = PersonFlaggedMatchState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PersonFlaggedMatchState: %v", v)
	}
}
