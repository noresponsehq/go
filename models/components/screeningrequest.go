package components

// ScreeningRequest - Screens the group's people no screening has reached and everyone on watch or escalated.
type ScreeningRequest struct {
	Depth     ScreeningDepth   `json:"depth"`
	Platforms []HandlePlatform `json:"platforms"`
}

func (s *ScreeningRequest) GetDepth() ScreeningDepth {
	if s == nil {
		return ScreeningDepth("")
	}
	return s.Depth
}

func (s *ScreeningRequest) GetPlatforms() []HandlePlatform {
	if s == nil {
		return []HandlePlatform{}
	}
	return s.Platforms
}
