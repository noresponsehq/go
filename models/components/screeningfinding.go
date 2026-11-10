package components

// ScreeningFinding - Something the analyser flagged. band is capped by the strongest match among the profiles of its evidence: possible at most moderate, probable at most high. capped_from is the band the analyser gave when the cap lowered it.
type ScreeningFinding struct {
	ID string `json:"id"`
	// What a finding is about: a direct threat to the principal or the event, intent to harm without a named target, fixation on the principal, a grievance against the host, research into the venue, route or security, weapons tied to any of these, disruption or protest planning toward this event, violent extremist content or identification, prior harassment or stalking of staff or principals, or posting the principal's addresses or movements.
	Category FindingCategory `json:"category"`
	// How severe a finding is, most severe first.
	Band FindingBand `json:"band"`
	// How severe a finding is, most severe first.
	CappedFrom *FindingBand `json:"capped_from,omitzero"`
	Confidence float64      `json:"confidence"`
	Rationale  string       `json:"rationale"`
	// The posts and media that support the finding.
	Evidence []string `json:"evidence"`
	// The resolver, collector or analyser that produced a record, by its catalogue name.
	Analyser string `json:"analyser"`
	Model    string `json:"model"`
	// What the analysis cost, in US cents.
	SpendMinor int64  `json:"spend_minor"`
	CreatedAt  string `json:"created_at"`
}

func (s *ScreeningFinding) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *ScreeningFinding) GetCategory() FindingCategory {
	if s == nil {
		return FindingCategory("")
	}
	return s.Category
}

func (s *ScreeningFinding) GetBand() FindingBand {
	if s == nil {
		return FindingBand("")
	}
	return s.Band
}

func (s *ScreeningFinding) GetCappedFrom() *FindingBand {
	if s == nil {
		return nil
	}
	return s.CappedFrom
}

func (s *ScreeningFinding) GetConfidence() float64 {
	if s == nil {
		return 0.0
	}
	return s.Confidence
}

func (s *ScreeningFinding) GetRationale() string {
	if s == nil {
		return ""
	}
	return s.Rationale
}

func (s *ScreeningFinding) GetEvidence() []string {
	if s == nil {
		return []string{}
	}
	return s.Evidence
}

func (s *ScreeningFinding) GetAnalyser() string {
	if s == nil {
		return ""
	}
	return s.Analyser
}

func (s *ScreeningFinding) GetModel() string {
	if s == nil {
		return ""
	}
	return s.Model
}

func (s *ScreeningFinding) GetSpendMinor() int64 {
	if s == nil {
		return 0
	}
	return s.SpendMinor
}

func (s *ScreeningFinding) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}
