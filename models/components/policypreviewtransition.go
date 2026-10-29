package components

type PolicyPreviewTransition struct {
	From  From  `json:"from"`
	To    From  `json:"to"`
	Count int64 `json:"count"`
}

func (p *PolicyPreviewTransition) GetFrom() From {
	if p == nil {
		return From("")
	}
	return p.From
}

func (p *PolicyPreviewTransition) GetTo() From {
	if p == nil {
		return From("")
	}
	return p.To
}

func (p *PolicyPreviewTransition) GetCount() int64 {
	if p == nil {
		return 0
	}
	return p.Count
}
