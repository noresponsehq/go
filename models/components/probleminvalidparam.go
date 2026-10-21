package components

type ProblemInvalidParam struct {
	Field   string  `json:"field"`
	Reason  string  `json:"reason"`
	Example *string `json:"example,omitzero"`
}

func (p *ProblemInvalidParam) GetField() string {
	if p == nil {
		return ""
	}
	return p.Field
}

func (p *ProblemInvalidParam) GetReason() string {
	if p == nil {
		return ""
	}
	return p.Reason
}

func (p *ProblemInvalidParam) GetExample() *string {
	if p == nil {
		return nil
	}
	return p.Example
}
