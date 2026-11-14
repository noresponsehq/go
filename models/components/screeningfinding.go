package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// ScreeningFinding - Something the analyser flagged. band is capped by the strongest match among the profiles of its evidence: possible at most moderate, probable at most high. capped_from is the band the analyser gave when the cap lowered it.
type ScreeningFinding struct {
	ID string `json:"id"`
	// What a finding is about: a direct threat to the principal or the group, intent to harm without a named target, fixation on the principal, a grievance against the host, research into the venue, route or security, weapons tied to any of these, disruption or protest planning toward this group, violent extremist content or identification, prior harassment or stalking of staff or principals, or posting the principal's addresses or movements.
	Category FindingCategory `json:"category"`
	// How severe a finding is, most severe first.
	Band FindingBand `json:"band"`
	// How severe a finding is, most severe first.
	CappedFrom *FindingBand `json:"capped_from,omitzero"`
	Confidence float64      `json:"confidence"`
	Rationale  string       `json:"rationale"`
	// The exact words of the first evidence item that the finding rests on, as the analyser quoted them.
	Quote *string `json:"quote,omitzero"`
	// The posts and media that support the finding.
	Evidence []string `json:"evidence"`
	// What the analyser saw in each image among evidence that the finding rests on. Unlike quote, an observation cannot be checked against the record: judge it against the image at the media item's permalink.
	MediaObservations []ScreeningMediaObservation `json:"media_observations,omitzero"`
	// The resolver, collector or analyser that produced a record, by its catalogue name.
	Analyser string `json:"analyser"`
	Model    string `json:"model"`
	// What the analysis cost, in US cents.
	SpendMinor int64  `json:"spend_minor"`
	CreatedAt  string `json:"created_at"`
}

func (s ScreeningFinding) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningFinding) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
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

func (s *ScreeningFinding) GetQuote() *string {
	if s == nil {
		return nil
	}
	return s.Quote
}

func (s *ScreeningFinding) GetEvidence() []string {
	if s == nil {
		return []string{}
	}
	return s.Evidence
}

func (s *ScreeningFinding) GetMediaObservations() []ScreeningMediaObservation {
	if s == nil {
		return nil
	}
	return s.MediaObservations
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
