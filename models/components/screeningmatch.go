package components

// ScreeningMatch - How a profile was matched to the attendee: the state its evidence earns, and the evidence.
type ScreeningMatch struct {
	// How sure the screening is that a profile is the attendee's, from its evidence: confirmed is a handle on the attendee row, or two or more hard identifiers (email, phone, a supplied handle, or a fact an existing source returned); probable is one identifier beyond the name; possible is the name alone.
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
