package components

// ScreeningPlatformCoverage - What a screening read on one platform.
type ScreeningPlatformCoverage struct {
	// A platform a screening can cover. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
	Platform HandlePlatform `json:"platform"`
	// covered read at least one profile; none_found looked and found no account; inaccessible found accounts that are private or locked; failed could not finish looking; unavailable has no resolver or collector enabled.
	Status        ScreeningPlatformCoverageStatus `json:"status"`
	AccountsFound int64                           `json:"accounts_found"`
	ItemsRead     int64                           `json:"items_read"`
}

func (s *ScreeningPlatformCoverage) GetPlatform() HandlePlatform {
	if s == nil {
		return HandlePlatform("")
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
