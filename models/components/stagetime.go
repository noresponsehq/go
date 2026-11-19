package components

type StageTime struct {
	// ai_typing and analysis appear only on searches recorded before the engine stopped analysing results.
	Stage Stage `json:"stage"`
	Ms    int64 `json:"ms"`
}

func (s *StageTime) GetStage() Stage {
	if s == nil {
		return Stage("")
	}
	return s.Stage
}

func (s *StageTime) GetMs() int64 {
	if s == nil {
		return 0
	}
	return s.Ms
}
