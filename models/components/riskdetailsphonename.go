package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type RiskDetailsPhoneName struct {
	Value  string `json:"value"`
	SeenBy int64  `json:"seen_by"`
}

func (r RiskDetailsPhoneName) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskDetailsPhoneName) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskDetailsPhoneName) GetValue() string {
	if r == nil {
		return ""
	}
	return r.Value
}

func (r *RiskDetailsPhoneName) GetSeenBy() int64 {
	if r == nil {
		return 0
	}
	return r.SeenBy
}
