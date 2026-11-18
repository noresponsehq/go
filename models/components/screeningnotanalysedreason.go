package components

// ScreeningNotAnalysedReason - Why an item read was not analysed. video: a video with no cover frame the analyser could read, whose caption alone is read. image_unavailable: the image could not be fetched from the platform. image_unsupported: not a JPEG, PNG, GIF or WebP image, or too large to send. image_limit: past the most images one person's analysis reads. refused: the model declined to read it. unfinished: the model cut its batch short or answered outside the schema, or the request was rejected. unverified: a finding cited it with words it does not contain or an image it did not describe. unspecified: the analyser gave no reason. analysis_off: no analyser is switched on, so the investigation found, read and typed everything without AI and analysed nothing.
type ScreeningNotAnalysedReason string

const (
	ScreeningNotAnalysedReasonVideo            ScreeningNotAnalysedReason = "video"
	ScreeningNotAnalysedReasonImageUnavailable ScreeningNotAnalysedReason = "image_unavailable"
	ScreeningNotAnalysedReasonImageUnsupported ScreeningNotAnalysedReason = "image_unsupported"
	ScreeningNotAnalysedReasonImageLimit       ScreeningNotAnalysedReason = "image_limit"
	ScreeningNotAnalysedReasonRefused          ScreeningNotAnalysedReason = "refused"
	ScreeningNotAnalysedReasonUnfinished       ScreeningNotAnalysedReason = "unfinished"
	ScreeningNotAnalysedReasonUnverified       ScreeningNotAnalysedReason = "unverified"
	ScreeningNotAnalysedReasonUnspecified      ScreeningNotAnalysedReason = "unspecified"
	ScreeningNotAnalysedReasonAnalysisOff      ScreeningNotAnalysedReason = "analysis_off"
)

func (e ScreeningNotAnalysedReason) ToPointer() *ScreeningNotAnalysedReason {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningNotAnalysedReason) IsExact() bool {
	if e != nil {
		switch *e {
		case "video", "image_unavailable", "image_unsupported", "image_limit", "refused", "unfinished", "unverified", "unspecified", "analysis_off":
			return true
		}
	}
	return false
}
