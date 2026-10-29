package components

type ReviewComment struct {
	ID            string `json:"id"`
	ReviewVersion int64  `json:"review_version"`
	Body          string `json:"body"`
	CreatedBy     string `json:"created_by"`
	CreatedAt     string `json:"created_at"`
}

func (r *ReviewComment) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *ReviewComment) GetReviewVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ReviewVersion
}

func (r *ReviewComment) GetBody() string {
	if r == nil {
		return ""
	}
	return r.Body
}

func (r *ReviewComment) GetCreatedBy() string {
	if r == nil {
		return ""
	}
	return r.CreatedBy
}

func (r *ReviewComment) GetCreatedAt() string {
	if r == nil {
		return ""
	}
	return r.CreatedAt
}
