package components

// ReviewSubjectChange - The link and resulting review version.
type ReviewSubjectChange struct {
	ReviewVersion int64 `json:"review_version"`
	// A tenant-local investigation link. risk_version is the account risk version captured when the link was added. It does not assert a relationship or change feedback or enforcement.
	Subject ReviewSubject `json:"subject"`
}

func (r *ReviewSubjectChange) GetReviewVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ReviewVersion
}

func (r *ReviewSubjectChange) GetSubject() ReviewSubject {
	if r == nil {
		return ReviewSubject{}
	}
	return r.Subject
}
