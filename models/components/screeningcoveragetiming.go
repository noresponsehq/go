package components

// ScreeningCoverageTiming - How long this search's latest attempt took, stage by stage. started_at minus the investigation's creation is the time spent queued or retrying.
type ScreeningCoverageTiming struct {
	// When the attempt that wrote this result started.
	StartedAt string `json:"started_at"`
	// Milliseconds from started_at to the end of the last stage listed.
	TotalMs int64 `json:"total_ms"`
	// Stages finished so far, in order. The result first written lists stages up to written; the settled result lists them all.
	Stages []StageTime `json:"stages"`
}

func (s *ScreeningCoverageTiming) GetStartedAt() string {
	if s == nil {
		return ""
	}
	return s.StartedAt
}

func (s *ScreeningCoverageTiming) GetTotalMs() int64 {
	if s == nil {
		return 0
	}
	return s.TotalMs
}

func (s *ScreeningCoverageTiming) GetStages() []StageTime {
	if s == nil {
		return []StageTime{}
	}
	return s.Stages
}
