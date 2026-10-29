package components

// ReviewSubjectList - Retained tenant-local subject links and the current review version.
type ReviewSubjectList struct {
	ReviewVersion int64           `json:"review_version"`
	Items         []ReviewSubject `json:"items"`
}

func (r *ReviewSubjectList) GetReviewVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ReviewVersion
}

func (r *ReviewSubjectList) GetItems() []ReviewSubject {
	if r == nil {
		return []ReviewSubject{}
	}
	return r.Items
}
