package components

import (
	"encoding/json"
	"fmt"
)

// ListMatch - The list action captured at the first answer. Requires revision first. Older answers without captured list state do not match, including none.
type ListMatch string

const (
	ListMatchAllow ListMatch = "allow"
	ListMatchBlock ListMatch = "block"
	ListMatchNone  ListMatch = "none"
)

func (e ListMatch) ToPointer() *ListMatch {
	return &e
}
func (e *ListMatch) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		fallthrough
	case "block":
		fallthrough
	case "none":
		*e = ListMatch(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListMatch: %v", v)
	}
}
