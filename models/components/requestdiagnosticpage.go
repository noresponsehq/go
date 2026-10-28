package components

// RequestDiagnosticPage - Authenticated attempts in a bounded interval, newest first. Keep the same filters when paging.
type RequestDiagnosticPage struct {
	Items      []RequestDiagnostic `json:"items"`
	NextCursor *string             `json:"next_cursor,omitzero"`
}

func (r *RequestDiagnosticPage) GetItems() []RequestDiagnostic {
	if r == nil {
		return []RequestDiagnostic{}
	}
	return r.Items
}

func (r *RequestDiagnosticPage) GetNextCursor() *string {
	if r == nil {
		return nil
	}
	return r.NextCursor
}
