package components

import (
	"encoding/json"
	"fmt"
)

type ListPeopleDecision string

const (
	ListPeopleDecisionCleared   ListPeopleDecision = "cleared"
	ListPeopleDecisionWatch     ListPeopleDecision = "watch"
	ListPeopleDecisionEscalated ListPeopleDecision = "escalated"
	ListPeopleDecisionUndecided ListPeopleDecision = "undecided"
)

func (e ListPeopleDecision) ToPointer() *ListPeopleDecision {
	return &e
}
func (e *ListPeopleDecision) UnmarshalJSON(data []byte) error {
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
		*e = ListPeopleDecision(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPeopleDecision: %v", v)
	}
}
