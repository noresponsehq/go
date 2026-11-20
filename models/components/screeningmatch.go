package components

// ScreeningMatch - How a profile was matched to the person: the state its evidence earns, and the evidence.
type ScreeningMatch struct {
	// How sure the screening is that a profile is the person's, from the likelihood its record scores. Confirmed reaches the join bar of 0.9. Probable reaches 0.85, which takes a shared value beyond the one that found the account. Possible is the rest.
	State    ProfileMatchState        `json:"state"`
	Evidence []ScreeningMatchEvidence `json:"evidence"`
}

func (s *ScreeningMatch) GetState() ProfileMatchState {
	if s == nil {
		return ProfileMatchState("")
	}
	return s.State
}

func (s *ScreeningMatch) GetEvidence() []ScreeningMatchEvidence {
	if s == nil {
		return []ScreeningMatchEvidence{}
	}
	return s.Evidence
}
