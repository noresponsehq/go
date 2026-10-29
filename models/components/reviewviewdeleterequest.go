package components

type ReviewViewDeleteRequest struct {
	ExpectedVersion int64 `json:"expected_version"`
}

func (r *ReviewViewDeleteRequest) GetExpectedVersion() int64 {
	if r == nil {
		return 0
	}
	return r.ExpectedVersion
}
