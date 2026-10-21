package components

type RiskDetailsPhoneName struct {
	Value  string `json:"value"`
	SeenBy int64  `json:"seen_by"`
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
