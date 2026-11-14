package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// ScreeningPost - One post a collector read, with its permalink and time.
type ScreeningPost struct {
	ID string `json:"id"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	kind      string  `const:"post" json:"kind"`
	Text      string  `json:"text"`
	ReplyTo   *string `json:"reply_to,omitzero"`
	ProfileID string  `json:"profile_id"`
	Permalink string  `json:"permalink"`
	PostedAt  string  `json:"posted_at"`
	// SHA-256 of the screenshot taken of a flagged item.
	ScreenshotSha256 *string `json:"screenshot_sha256,omitzero"`
	// When the screenshot was taken.
	CapturedAt *string `json:"captured_at,omitzero"`
}

func (s ScreeningPost) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningPost) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningPost) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *ScreeningPost) GetKind() string {
	return "post"
}

func (s *ScreeningPost) GetText() string {
	if s == nil {
		return ""
	}
	return s.Text
}

func (s *ScreeningPost) GetReplyTo() *string {
	if s == nil {
		return nil
	}
	return s.ReplyTo
}

func (s *ScreeningPost) GetProfileID() string {
	if s == nil {
		return ""
	}
	return s.ProfileID
}

func (s *ScreeningPost) GetPermalink() string {
	if s == nil {
		return ""
	}
	return s.Permalink
}

func (s *ScreeningPost) GetPostedAt() string {
	if s == nil {
		return ""
	}
	return s.PostedAt
}

func (s *ScreeningPost) GetScreenshotSha256() *string {
	if s == nil {
		return nil
	}
	return s.ScreenshotSha256
}

func (s *ScreeningPost) GetCapturedAt() *string {
	if s == nil {
		return nil
	}
	return s.CapturedAt
}
