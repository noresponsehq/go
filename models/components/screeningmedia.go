package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// ScreeningMedia - One image or video a collector read, with the post that carried it.
type ScreeningMedia struct {
	ID string `json:"id"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	kind      string    `const:"media" json:"kind"`
	PostID    *string   `json:"post_id,omitzero"`
	MediaType MediaType `json:"media_type"`
	// Where the platform serves the image or video.
	MediaURL string `json:"media_url"`
	// The caption or alt text.
	Text      *string `json:"text,omitzero"`
	ProfileID string  `json:"profile_id"`
	Permalink string  `json:"permalink"`
	PostedAt  string  `json:"posted_at"`
	// SHA-256 of the screenshot taken of a flagged item.
	ScreenshotSha256 *string `json:"screenshot_sha256,omitzero"`
	// When the screenshot was taken.
	CapturedAt *string `json:"captured_at,omitzero"`
}

func (s ScreeningMedia) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningMedia) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningMedia) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *ScreeningMedia) GetKind() string {
	return "media"
}

func (s *ScreeningMedia) GetPostID() *string {
	if s == nil {
		return nil
	}
	return s.PostID
}

func (s *ScreeningMedia) GetMediaType() MediaType {
	if s == nil {
		return MediaType("")
	}
	return s.MediaType
}

func (s *ScreeningMedia) GetMediaURL() string {
	if s == nil {
		return ""
	}
	return s.MediaURL
}

func (s *ScreeningMedia) GetText() *string {
	if s == nil {
		return nil
	}
	return s.Text
}

func (s *ScreeningMedia) GetProfileID() string {
	if s == nil {
		return ""
	}
	return s.ProfileID
}

func (s *ScreeningMedia) GetPermalink() string {
	if s == nil {
		return ""
	}
	return s.Permalink
}

func (s *ScreeningMedia) GetPostedAt() string {
	if s == nil {
		return ""
	}
	return s.PostedAt
}

func (s *ScreeningMedia) GetScreenshotSha256() *string {
	if s == nil {
		return nil
	}
	return s.ScreenshotSha256
}

func (s *ScreeningMedia) GetCapturedAt() *string {
	if s == nil {
		return nil
	}
	return s.CapturedAt
}
