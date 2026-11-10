package components

// ScreeningMediaObservation - What an analyser saw in one image that a finding rests on.
type ScreeningMediaObservation struct {
	// The image, one of the finding's evidence.
	MediaID string `json:"media_id"`
	// What the analyser saw in the image that supports the finding.
	Description string `json:"description"`
	// Words the analyser read in the image, as it read them. They come from the image's pixels, not from a post's text, so they are not a quote and were not checked against the record.
	ReadText *string `json:"read_text,omitzero"`
}

func (s *ScreeningMediaObservation) GetMediaID() string {
	if s == nil {
		return ""
	}
	return s.MediaID
}

func (s *ScreeningMediaObservation) GetDescription() string {
	if s == nil {
		return ""
	}
	return s.Description
}

func (s *ScreeningMediaObservation) GetReadText() *string {
	if s == nil {
		return nil
	}
	return s.ReadText
}
