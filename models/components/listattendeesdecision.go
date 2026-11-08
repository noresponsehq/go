package components

import (
	"encoding/json"
	"fmt"
)

type ListAttendeesDecision string

const (
	ListAttendeesDecisionCleared   ListAttendeesDecision = "cleared"
	ListAttendeesDecisionWatch     ListAttendeesDecision = "watch"
	ListAttendeesDecisionEscalated ListAttendeesDecision = "escalated"
	ListAttendeesDecisionUndecided ListAttendeesDecision = "undecided"
)

func (e ListAttendeesDecision) ToPointer() *ListAttendeesDecision {
	return &e
}
func (e *ListAttendeesDecision) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cleared":
		fallthrough
	case "watch":
		fallthrough
	case "escalated":
		fallthrough
	case "undecided":
		*e = ListAttendeesDecision(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAttendeesDecision: %v", v)
	}
}
