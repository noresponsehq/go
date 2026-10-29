package components

type ReviewCommentRequest struct {
	ExpectedVersion int64  `json:"expected_version"`
	Body            string `json:"body"`
}

func (r *ReviewCommentRequest) GetExpectedVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ExpectedVersion
}

func (r *ReviewCommentRequest) GetBody() string {
	if r == nil {
		return ""
	}
	return r.Body
}
