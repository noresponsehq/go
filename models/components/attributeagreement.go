package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// AttributeAgreement - How far independent sources agree on one attribute: values is the number of distinct values reported, support the tables and profiles reporting any, and agreement match (one value), partial (the most reported value holds more than half the support) or mismatch. The values themselves are not served.
type AttributeAgreement struct {
	Agreement Agreement `json:"agreement"`
	Values    int64     `json:"values"`
	Support   int64     `json:"support"`
}

func (a AttributeAgreement) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttributeAgreement) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttributeAgreement) GetAgreement() Agreement {
	if a == nil {
		return Agreement("")
	}
	return a.Agreement
}

func (a *AttributeAgreement) GetValues() int64 {
	if a == nil {
		return 0
	}
	return a.Values
}

func (a *AttributeAgreement) GetSupport() int64 {
	if a == nil {
		return 0
	}
	return a.Support
}
