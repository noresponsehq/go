package components

// ScreeningNotAnalysedCount - How many items read were not analysed for one reason.
type ScreeningNotAnalysedCount struct {
	// Why an item read was not analysed. video: a video, whose caption alone is read. image_unavailable: the image could not be fetched from the platform. image_unsupported: not a JPEG, PNG, GIF or WebP image, or too large to send. image_limit: past the most images one person's analysis reads. refused: the model declined to read it. unfinished: the model cut its batch short or answered outside the schema, or the request was rejected. unverified: a finding cited it with words it does not contain or an image it did not describe. unspecified: the analyser gave no reason.
	Reason ScreeningNotAnalysedReason `json:"reason"`
	Items  int64                      `json:"items"`
}

func (s *ScreeningNotAnalysedCount) GetReason() ScreeningNotAnalysedReason {
	if s == nil {
		return ScreeningNotAnalysedReason("")
	}
	return s.Reason
}

func (s *ScreeningNotAnalysedCount) GetItems() int64 {
	if s == nil {
		return 0
	}
	return s.Items
}
