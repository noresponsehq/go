package components

// ScreeningProfile - An account the screening matched to the person, how it was matched, and how much of it was read.
type ScreeningProfile struct {
	ID string `json:"id"`
	// A platform a screening can cover. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
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
	// The oldest item read.
	PeriodFrom *string `json:"period_from,omitzero"`
	// The newest item read.
	PeriodTo    *string `json:"period_to,omitzero"`
	CollectedAt string  `json:"collected_at"`
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
