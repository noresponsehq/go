package components

type RiskDetailsBreachName struct {
	Value  string `json:"value"`
	SeenIn int64  `json:"seen_in"`
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
