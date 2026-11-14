package components

// SearchStatus - A password or hash search is counting, then counted: rows says how many breach rows hold the value, and nothing runs until a person runs it (POST .../searches/{search_id}/run). Any other search is queued, then running. completed: every source it asked answered. partial: some did not. failed: none answered or none could be asked, which says nothing about whether the value is held anywhere.
type SearchStatus string

const (
	SearchStatusCounting  SearchStatus = "counting"
	SearchStatusCounted   SearchStatus = "counted"
	SearchStatusQueued    SearchStatus = "queued"
	SearchStatusRunning   SearchStatus = "running"
	SearchStatusCompleted SearchStatus = "completed"
	SearchStatusPartial   SearchStatus = "partial"
	SearchStatusFailed    SearchStatus = "failed"
)

func (e SearchStatus) ToPointer() *SearchStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SearchStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "counting", "counted", "queued", "running", "completed", "partial", "failed":
			return true
		}
	}
	return false
}
