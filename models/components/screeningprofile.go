package components

import (
	"github.com/noresponsehq/go/v12/internal/utils"
)

// ScreeningProfile - An account the screening matched to the person, how it was matched, and how much of it was read.
type ScreeningProfile struct {
	ID string `json:"id"`
	// A platform a screening can cover. github and scratch through flickr are each one read of the public profile a handle names, with no posts; chesscom also reads the profile's stats, clubs and last two months of games, for its activity and contacts. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
	Platform         HandlePlatform `json:"platform"`
	Handle           string         `json:"handle"`
	URL              string         `json:"url"`
	DisplayName      *string        `json:"display_name,omitzero"`
	Bio              *string        `json:"bio,omitzero"`
	AvatarURL        *string        `json:"avatar_url,omitzero"`
	Followers        *int64         `json:"followers,omitzero"`
	AccountCreatedAt *string        `json:"account_created_at,omitzero"`
	// public profiles were read; a private or locked one exists but could not be read.
	Access Access `json:"access"`
	// How a profile was matched to the person: the state its evidence earns, and the evidence.
	Match ScreeningMatch `json:"match"`
	// The resolver, collector or analyser that produced a record, by its catalogue name.
	Collector string `json:"collector"`
	ItemsRead int64  `json:"items_read"`
	// Why fewer posts were read than the screening's depth asks for while the account holds more. platform_limit: the platform shows the screening no more of them. page_limit: the pages a screening reads ran out first. rate_limited or unavailable: a later page was held back or failed; the posts before it are kept and the person is partial. Absent when the account was read to its end or to the depth.
	PostsStopped *PostsStopped `json:"posts_stopped,omitzero"`
	// The oldest item read.
	PeriodFrom *string `json:"period_from,omitzero"`
	// The newest item read.
	PeriodTo    *string `json:"period_to,omitzero"`
	CollectedAt string  `json:"collected_at"`
	// One answer from one source: where it came from, what found it and every field it holds. A person's lookup records also say how likely each is the person.
	Record *Record `json:"record,omitzero"`
}

func (s ScreeningProfile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningProfile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningProfile) GetID() string {
	if s == nil {
		return ""
	}
	return s.ID
}

func (s *ScreeningProfile) GetPlatform() HandlePlatform {
	if s == nil {
		return HandlePlatform("")
	}
	return s.Platform
}

func (s *ScreeningProfile) GetHandle() string {
	if s == nil {
		return ""
	}
	return s.Handle
}

func (s *ScreeningProfile) GetURL() string {
	if s == nil {
		return ""
	}
	return s.URL
}

func (s *ScreeningProfile) GetDisplayName() *string {
	if s == nil {
		return nil
	}
	return s.DisplayName
}

func (s *ScreeningProfile) GetBio() *string {
	if s == nil {
		return nil
	}
	return s.Bio
}

func (s *ScreeningProfile) GetAvatarURL() *string {
	if s == nil {
		return nil
	}
	return s.AvatarURL
}

func (s *ScreeningProfile) GetFollowers() *int64 {
	if s == nil {
		return nil
	}
	return s.Followers
}

func (s *ScreeningProfile) GetAccountCreatedAt() *string {
	if s == nil {
		return nil
	}
	return s.AccountCreatedAt
}

func (s *ScreeningProfile) GetAccess() Access {
	if s == nil {
		return Access("")
	}
	return s.Access
}

func (s *ScreeningProfile) GetMatch() ScreeningMatch {
	if s == nil {
		return ScreeningMatch{}
	}
	return s.Match
}

func (s *ScreeningProfile) GetCollector() string {
	if s == nil {
		return ""
	}
	return s.Collector
}

func (s *ScreeningProfile) GetItemsRead() int64 {
	if s == nil {
		return 0
	}
	return s.ItemsRead
}

func (s *ScreeningProfile) GetPostsStopped() *PostsStopped {
	if s == nil {
		return nil
	}
	return s.PostsStopped
}

func (s *ScreeningProfile) GetPeriodFrom() *string {
	if s == nil {
		return nil
	}
	return s.PeriodFrom
}

func (s *ScreeningProfile) GetPeriodTo() *string {
	if s == nil {
		return nil
	}
	return s.PeriodTo
}

func (s *ScreeningProfile) GetCollectedAt() string {
	if s == nil {
		return ""
	}
	return s.CollectedAt
}

func (s *ScreeningProfile) GetRecord() *Record {
	if s == nil {
		return nil
	}
	return s.Record
}
