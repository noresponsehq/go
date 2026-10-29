package components

type ReviewViewRequest struct {
	Name            string `json:"name"`
	Shared          bool   `json:"shared"`
	ExpectedVersion *int64 `json:"expected_version,omitzero"`
	Filter          Filter `json:"filter"`
}

func (r *ReviewViewRequest) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r *ReviewViewRequest) GetShared() bool {
	if r == nil {
		return false
	}
	return r.Shared
}

func (r *ReviewViewRequest) GetExpectedVersion() *int64 {
	if r == nil {
		return nil
	}
	return r.ExpectedVersion
}

func (r *ReviewViewRequest) GetFilter() Filter {
	if r == nil {
		return Filter{}
	}
	return r.Filter
}
