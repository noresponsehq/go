package components

// ReviewsConclusions - Current resolution of each resolved retained review; prior resolutions and reopened reviews do not count. Missing keys mean zero. These investigation conclusions are independent of account feedback.
type ReviewsConclusions struct {
	Benign       *int64 `json:"benign,omitzero"`
	Abuse        *int64 `json:"abuse,omitzero"`
	Inconclusive *int64 `json:"inconclusive,omitzero"`
}

func (r *ReviewsConclusions) GetBenign() *int64 {
	if r == nil {
		return nil
	}
	return r.Benign
}

func (r *ReviewsConclusions) GetAbuse() *int64 {
	if r == nil {
		return nil
	}
	return r.Abuse
}

func (r *ReviewsConclusions) GetInconclusive() *int64 {
	if r == nil {
		return nil
	}
	return r.Inconclusive
}
