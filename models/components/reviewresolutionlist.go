package components

// ReviewResolutionList - Read immutable review resolution history.
type ReviewResolutionList struct {
	Items             []Resolution `json:"items"`
	NextBeforeVersion *int64       `json:"next_before_version,omitzero"`
}

func (r *ReviewResolutionList) GetItems() []Resolution {
	if r == nil {
		return []Resolution{}
	}
	return r.Items
}

func (r *ReviewResolutionList) GetNextBeforeVersion() *int64 {
	if r == nil {
		return nil
	}
	return r.NextBeforeVersion
}
