package components

import (
	"encoding/json"
	"fmt"
)

type ListPeopleBand string

const (
	ListPeopleBandCritical ListPeopleBand = "critical"
	ListPeopleBandHigh     ListPeopleBand = "high"
	ListPeopleBandModerate ListPeopleBand = "moderate"
	ListPeopleBandLow      ListPeopleBand = "low"
	ListPeopleBandInfo     ListPeopleBand = "info"
	ListPeopleBandNone     ListPeopleBand = "none"
)

func (e ListPeopleBand) ToPointer() *ListPeopleBand {
	return &e
}
func (e *ListPeopleBand) UnmarshalJSON(data []byte) error {
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
		*e = ListPeopleBand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPeopleBand: %v", v)
	}
}
