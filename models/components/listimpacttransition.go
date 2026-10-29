package components

type ListImpactTransition struct {
	From  ListImpactTransitionFrom `json:"from"`
	To    ListImpactTransitionFrom `json:"to"`
	Count int64                    `json:"count"`
}

func (l *ListImpactTransition) GetFrom() ListImpactTransitionFrom {
	if l == nil {
		return ListImpactTransitionFrom("")
	}
	return l.From
}

func (l *ListImpactTransition) GetTo() ListImpactTransitionFrom {
	if l == nil {
		return ListImpactTransitionFrom("")
	}
	return l.To
}

func (l *ListImpactTransition) GetCount() int64 {
	if l == nil {
		return 0
	}
	return l.Count
}
