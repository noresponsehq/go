package components

import (
	"github.com/noresponsehq/go/v13/internal/utils"
)

// ScreeningPost - One post a collector read, with its permalink and time.
type ScreeningPost struct {
	ID string `json:"id"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	kind    string  `const:"post" json:"kind"`
	Text    string  `json:"text"`
	ReplyTo *string `json:"reply_to,omitzero"`
	// The permalink of the post this one reposts. text then holds that post's author's words, not the person's.
	RepostOf *string `json:"repost_of,omitzero"`
	// The permalink of the post this one quotes. text holds the person's own words alone.
	QuoteOf *string `json:"quote_of,omitzero"`
	// A place a post names, as the platform shows it.
	Place *PostPlace `json:"place,omitzero"`
	// The handles of the accounts the post tags, on its own platform.
	Tagged []string `json:"tagged,omitzero"`
	// How many likes the platform shows for the post.
	Likes *int64 `json:"likes,omitzero"`
	// How many comments the platform shows for the post.
	Comments *int64 `json:"comments,omitzero"`
	// How many views or plays the platform shows for the post.
	Views *int64 `json:"views,omitzero"`
	// How many times the platform shows the post was shared or reposted.
	Shares *int64 `json:"shares,omitzero"`
	// The language the platform says the post is in, as the platform codes it, such as en or pt-BR.
	Language *string `json:"language,omitzero"`
	// true when the platform marks the post as edited after it went up. Absent when it does not mark it.
	Edited *bool `json:"edited,omitzero"`
	// true when the platform marks the post as a paid partnership or an ad. Absent when it does not mark it.
	Sponsored *bool `json:"sponsored,omitzero"`
	// true when the platform labels the post as made or changed with AI. Absent when it does not label it.
	AiGenerated *bool `json:"ai_generated,omitzero"`
	// The app or device the platform says the post was made with.
	PostedWith *string `json:"posted_with,omitzero"`
	// Everything else the platform sent about the post, under the platform's own names and as it typed them: counts, ids of linked posts and accounts, links, media details, labels and notices. The fields above are not repeated here. Nothing that describes the account the screening read with, such as whether it liked or saved the post, is ever here.
	Stated    map[string]any `json:"stated,omitzero"`
	ProfileID string         `json:"profile_id"`
	Permalink string         `json:"permalink"`
	PostedAt  string         `json:"posted_at"`
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

func (s *ScreeningPost) GetRepostOf() *string {
	if s == nil {
		return nil
	}
	return s.RepostOf
}

func (s *ScreeningPost) GetQuoteOf() *string {
	if s == nil {
		return nil
	}
	return s.QuoteOf
}

func (s *ScreeningPost) GetPlace() *PostPlace {
	if s == nil {
		return nil
	}
	return s.Place
}

func (s *ScreeningPost) GetTagged() []string {
	if s == nil {
		return nil
	}
	return s.Tagged
}

func (s *ScreeningPost) GetLikes() *int64 {
	if s == nil {
		return nil
	}
	return s.Likes
}

func (s *ScreeningPost) GetComments() *int64 {
	if s == nil {
		return nil
	}
	return s.Comments
}

func (s *ScreeningPost) GetViews() *int64 {
	if s == nil {
		return nil
	}
	return s.Views
}

func (s *ScreeningPost) GetShares() *int64 {
	if s == nil {
		return nil
	}
	return s.Shares
}

func (s *ScreeningPost) GetLanguage() *string {
	if s == nil {
		return nil
	}
	return s.Language
}

func (s *ScreeningPost) GetEdited() *bool {
	if s == nil {
		return nil
	}
	return s.Edited
}

func (s *ScreeningPost) GetSponsored() *bool {
	if s == nil {
		return nil
	}
	return s.Sponsored
}

func (s *ScreeningPost) GetAiGenerated() *bool {
	if s == nil {
		return nil
	}
	return s.AiGenerated
}

func (s *ScreeningPost) GetPostedWith() *string {
	if s == nil {
		return nil
	}
	return s.PostedWith
}

func (s *ScreeningPost) GetStated() map[string]any {
	if s == nil {
		return nil
	}
	return s.Stated
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
