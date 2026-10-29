package components

// PolicyHistory - Immutable settings snapshots, newest first. history_complete is false when earlier versions predate capture. A restored version is published with a fresh version number. History is removed with the tenant.
type PolicyHistory struct {
	Items             []PolicyHistoryItem `json:"items"`
	HasMore           bool                `json:"has_more"`
	NextBeforeVersion *int64              `json:"next_before_version,omitzero"`
	HistoryComplete   bool                `json:"history_complete"`
}

func (p *PolicyHistory) GetItems() []PolicyHistoryItem {
	if p == nil {
		return []PolicyHistoryItem{}
	}
	return p.Items
}

func (p *PolicyHistory) GetHasMore() bool {
	if p == nil {
		return false
	}
	return p.HasMore
}

func (p *PolicyHistory) GetNextBeforeVersion() *int64 {
	if p == nil {
		return nil
	}
	return p.NextBeforeVersion
}

func (p *PolicyHistory) GetHistoryComplete() bool {
	if p == nil {
		return false
	}
	return p.HistoryComplete
}
