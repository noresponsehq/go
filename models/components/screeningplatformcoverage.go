package components

// ScreeningPlatformCoverage - What a screening read on one platform.
type ScreeningPlatformCoverage struct {
	// A platform a screening can cover, such as instagram, x, bluesky, youtube, tiktok or github. The screening registry owns the set: a request that names a platform it does not cover answers 422. A new platform adds an id without changing this contract, so read a platform you do not know in a response as one you have not seen before, not as an error. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
	Platform string `json:"platform"`
	// covered read at least one profile; none_found looked and found no account; inaccessible found accounts that are private or locked; failed could not finish looking; unavailable has no resolver or collector enabled.
	Status ScreeningPlatformCoverageStatus `json:"status"`
	// Accounts found on the platform: at most 5, the best candidates the search kept.
	AccountsFound int64 `json:"accounts_found"`
	ItemsRead     int64 `json:"items_read"`
}

func (s *ScreeningPlatformCoverage) GetPlatform() string {
	if s == nil {
		return ""
	}
	return s.Platform
}

func (s *ScreeningPlatformCoverage) GetStatus() ScreeningPlatformCoverageStatus {
	if s == nil {
		return ScreeningPlatformCoverageStatus("")
	}
	return s.Status
}

func (s *ScreeningPlatformCoverage) GetAccountsFound() int64 {
	if s == nil {
		return 0
	}
	return s.AccountsFound
}

func (s *ScreeningPlatformCoverage) GetItemsRead() int64 {
	if s == nil {
		return 0
	}
	return s.ItemsRead
}
