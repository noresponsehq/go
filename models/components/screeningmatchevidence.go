package components

// ScreeningMatchEvidence - One identifier that ties a profile to the attendee.
type ScreeningMatchEvidence struct {
	// What matched. email, phone, supplied_handle and source_fact are hard identifiers; company, city and school are soft.
	Kind ScreeningMatchEvidenceKind `json:"kind"`
	// Where the match was seen, such as the profile field that carried it.
	Detail *string `json:"detail,omitzero"`
	// The resolver, collector or analyser that produced a record, by its catalogue name.
	Resolver string `json:"resolver"`
}

func (s *ScreeningMatchEvidence) GetKind() ScreeningMatchEvidenceKind {
	if s == nil {
		return ScreeningMatchEvidenceKind("")
	}
	return s.Kind
}

func (s *ScreeningMatchEvidence) GetDetail() *string {
	if s == nil {
		return nil
	}
	return s.Detail
}

func (s *ScreeningMatchEvidence) GetResolver() string {
	if s == nil {
		return ""
	}
	return s.Resolver
}
