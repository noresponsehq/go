package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type RiskDetailsBreachName struct {
	Value  string `json:"value"`
	SeenIn int64  `json:"seen_in"`
}

func (r RiskDetailsBreachName) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskDetailsBreachName) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskDetailsBreachName) GetValue() string {
	if r == nil {
		return ""
	}
	return r.Value
}

func (r *RiskDetailsBreachName) GetSeenIn() int64 {
	if r == nil {
		return 0
	}
	return r.SeenIn
}
