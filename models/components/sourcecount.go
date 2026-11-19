package components

type SourceCount struct {
	Asked    int64 `json:"asked"`
	Answered int64 `json:"answered"`
}

func (s *SourceCount) GetAsked() int64 {
	if s == nil {
		return 0
	}
	return s.Asked
}

func (s *SourceCount) GetAnswered() int64 {
	if s == nil {
		return 0
	}
	return s.Answered
}
