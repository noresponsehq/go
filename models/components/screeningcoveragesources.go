package components

// ScreeningCoverageSources - How many sources this search asked and how many answered, across the lookup stage (sources asked whether an account exists for an identifier) and the platform stage (profiles read on each platform). A source that answered may have found nothing; one that did not answer failed or ran out of time.
type ScreeningCoverageSources struct {
	// Lookup sources plus platforms asked.
	Asked int64 `json:"asked"`
	// Lookup sources plus platforms that answered.
	Answered  int64       `json:"answered"`
	Lookups   SourceCount `json:"lookups"`
	Platforms SourceCount `json:"platforms"`
}

func (s *ScreeningCoverageSources) GetAsked() int64 {
	if s == nil {
		return 0
	}
	return s.Asked
}

func (s *ScreeningCoverageSources) GetAnswered() int64 {
	if s == nil {
		return 0
	}
	return s.Answered
}

func (s *ScreeningCoverageSources) GetLookups() SourceCount {
	if s == nil {
		return SourceCount{}
	}
	return s.Lookups
}

func (s *ScreeningCoverageSources) GetPlatforms() SourceCount {
	if s == nil {
		return SourceCount{}
	}
	return s.Platforms
}
