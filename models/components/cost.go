package components

type Cost struct {
	Allow  float64 `json:"allow"`
	StepUp float64 `json:"step_up"`
	Review float64 `json:"review"`
	Block  float64 `json:"block"`
}

func (c *Cost) GetAllow() float64 {
	if c == nil {
		return 0.0
	}
	return c.Allow
}

func (c *Cost) GetStepUp() float64 {
	if c == nil {
		return 0.0
	}
	return c.StepUp
}

func (c *Cost) GetReview() float64 {
	if c == nil {
		return 0.0
	}
	return c.Review
}

func (c *Cost) GetBlock() float64 {
	if c == nil {
		return 0.0
	}
	return c.Block
}
