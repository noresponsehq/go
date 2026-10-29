package components

// Export - Retained review snapshot.
type Export struct {
	ExportedAt  string          `json:"exported_at"`
	Review      Review          `json:"review"`
	Comments    []ReviewComment `json:"comments"`
	Resolutions []Resolution    `json:"resolutions"`
	Subjects    []ReviewSubject `json:"subjects"`
}

func (e *Export) GetExportedAt() string {
	if e == nil {
		return ""
	}
	return e.ExportedAt
}

func (e *Export) GetReview() Review {
	if e == nil {
		return Review{}
	}
	return e.Review
}

func (e *Export) GetComments() []ReviewComment {
	if e == nil {
		return []ReviewComment{}
	}
	return e.Comments
}

func (e *Export) GetResolutions() []Resolution {
	if e == nil {
		return []Resolution{}
	}
	return e.Resolutions
}

func (e *Export) GetSubjects() []ReviewSubject {
	if e == nil {
		return []ReviewSubject{}
	}
	return e.Subjects
}
