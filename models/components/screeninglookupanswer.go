package components

type ScreeningLookupAnswer struct {
	// The id of the connector asked: breach for the breach search.
	Connector  string     `json:"connector"`
	Identifier Identifier `json:"identifier"`
	// Per-signal outcome. A source that timed out, was blocked or errored reports blocked, unavailable or inconclusive with a reason and is listed in incomplete_sources; not_found is a definite absence.
	Outcome    SignalOutcome `json:"outcome"`
	AnsweredAt *string       `json:"answered_at,omitzero"`
}

func (s *ScreeningLookupAnswer) GetConnector() string {
	if s == nil {
		return ""
	}
	return s.Connector
}

func (s *ScreeningLookupAnswer) GetIdentifier() Identifier {
	if s == nil {
		return Identifier("")
	}
	return s.Identifier
}

func (s *ScreeningLookupAnswer) GetOutcome() SignalOutcome {
	if s == nil {
		return SignalOutcome("")
	}
	return s.Outcome
}

func (s *ScreeningLookupAnswer) GetAnsweredAt() *string {
	if s == nil {
		return nil
	}
	return s.AnsweredAt
}
