package components

// ScreeningSpend - What screening cost, in US cents, in total and per implementation.
type ScreeningSpend struct {
	// The total, in US cents.
	TotalMinor       int64                `json:"total_minor"`
	ByImplementation []ScreeningSpendLine `json:"by_implementation"`
}

func (s *ScreeningSpend) GetTotalMinor() int64 {
	if s == nil {
		return 0
	}
	return s.TotalMinor
}

func (s *ScreeningSpend) GetByImplementation() []ScreeningSpendLine {
	if s == nil {
		return []ScreeningSpendLine{}
	}
	return s.ByImplementation
}
