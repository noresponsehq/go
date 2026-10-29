package components

type Thresholds struct {
	StepUp int64  `json:"step_up"`
	Review int64  `json:"review"`
	Block  *int64 `json:"block"`
}

func (t *Thresholds) GetStepUp() int64 {
	if t == nil {
		return 0
	}
	return t.StepUp
}

func (t *Thresholds) GetReview() int64 {
	if t == nil {
		return 0
	}
	return t.Review
}

func (t *Thresholds) GetBlock() *int64 {
	if t == nil {
		return nil
	}
	return t.Block
}
