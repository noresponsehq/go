package components

// ScreeningRequest - Screens the group's people no screening has reached.
type ScreeningRequest struct {
	Depth ScreeningDepth `json:"depth"`
}

func (s *ScreeningRequest) GetDepth() ScreeningDepth {
	if s == nil {
		return ScreeningDepth("")
	}
	return s.Depth
}
