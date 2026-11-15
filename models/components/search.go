package components

// Search - One search a person started on a value of one of a person's records. Nothing searches again by itself.
type Search struct {
	SearchID string `json:"search_id"`
	// The record whose value led to the search.
	RecordID string `json:"record_id"`
	// The field of that record that holds the value.
	Field string `json:"field"`
	// The one breach table searched, when the search was of one table.
	Table *string `json:"table,omitzero"`
	// The kind of the value searched. A domain search finds the leaked rows of every email address at the domain.
	Kind SearchKind `json:"kind"`
	// The value searched: the field's cleaned value, else as the source wrote it.
	Value string `json:"value"`
	// A password or hash search is counting, then counted: rows says how many breach rows hold the value, and nothing runs until a person runs it (POST .../searches/{search_id}/run). Any other search is queued, then running. completed: every source it asked answered. partial: some did not. failed: none answered or none could be asked, which says nothing about whether the value is held anywhere.
	Status SearchStatus `json:"status"`
	// How many breach rows hold the password or hash, counted before the search runs.
	Rows *int64 `json:"rows,omitzero"`
	// How many sources the search asked: the breach search and the account and phone lookups that take the value's kind.
	Sources int64 `json:"sources"`
	// How many of them gave a definite answer, found or not found.
	Answered int64 `json:"answered"`
	// How many records the search found. They are among the person report's lookup records, each match naming this search.
	Found int64 `json:"found"`
	// What the search may spend, in US cents, held against the environment's daily screening cap while it runs.
	EstimatedSpend int64 `json:"estimated_spend"`
	// What the search spent, in US cents, metered as screening work.
	Spent int64 `json:"spent"`
	// Why a failed search found nothing: its sources did not answer, none could be asked, or it could not run before its attempts were spent.
	ErrorCode *ErrorCode `json:"error_code,omitzero"`
	// Who started the search.
	RequestedBy string  `json:"requested_by"`
	CreatedAt   string  `json:"created_at"`
	FinishedAt  *string `json:"finished_at,omitzero"`
}

func (s *Search) GetSearchID() string {
	if s == nil {
		return ""
	}
	return s.SearchID
}

func (s *Search) GetRecordID() string {
	if s == nil {
		return ""
	}
	return s.RecordID
}

func (s *Search) GetField() string {
	if s == nil {
		return ""
	}
	return s.Field
}

func (s *Search) GetTable() *string {
	if s == nil {
		return nil
	}
	return s.Table
}

func (s *Search) GetKind() SearchKind {
	if s == nil {
		return SearchKind("")
	}
	return s.Kind
}

func (s *Search) GetValue() string {
	if s == nil {
		return ""
	}
	return s.Value
}

func (s *Search) GetStatus() SearchStatus {
	if s == nil {
		return SearchStatus("")
	}
	return s.Status
}

func (s *Search) GetRows() *int64 {
	if s == nil {
		return nil
	}
	return s.Rows
}

func (s *Search) GetSources() int64 {
	if s == nil {
		return 0
	}
	return s.Sources
}

func (s *Search) GetAnswered() int64 {
	if s == nil {
		return 0
	}
	return s.Answered
}

func (s *Search) GetFound() int64 {
	if s == nil {
		return 0
	}
	return s.Found
}

func (s *Search) GetEstimatedSpend() int64 {
	if s == nil {
		return 0
	}
	return s.EstimatedSpend
}

func (s *Search) GetSpent() int64 {
	if s == nil {
		return 0
	}
	return s.Spent
}

func (s *Search) GetErrorCode() *ErrorCode {
	if s == nil {
		return nil
	}
	return s.ErrorCode
}

func (s *Search) GetRequestedBy() string {
	if s == nil {
		return ""
	}
	return s.RequestedBy
}

func (s *Search) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}

func (s *Search) GetFinishedAt() *string {
	if s == nil {
		return nil
	}
	return s.FinishedAt
}
