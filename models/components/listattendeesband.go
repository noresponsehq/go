package components

import (
	"encoding/json"
	"fmt"
)

type ListAttendeesBand string

const (
	ListAttendeesBandCritical ListAttendeesBand = "critical"
	ListAttendeesBandHigh     ListAttendeesBand = "high"
	ListAttendeesBandModerate ListAttendeesBand = "moderate"
	ListAttendeesBandLow      ListAttendeesBand = "low"
	ListAttendeesBandInfo     ListAttendeesBand = "info"
	ListAttendeesBandNone     ListAttendeesBand = "none"
)

func (e ListAttendeesBand) ToPointer() *ListAttendeesBand {
	return &e
}
func (e *ListAttendeesBand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "critical":
		fallthrough
	case "high":
		fallthrough
	case "moderate":
		fallthrough
	case "low":
		fallthrough
	case "info":
		fallthrough
	case "none":
		*e = ListAttendeesBand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAttendeesBand: %v", v)
	}
}
