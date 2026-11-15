package components

// WorkspacePersonResult - Create or reuse a workspace person.
type WorkspacePersonResult struct {
	Person           WorkspacePerson `json:"person"`
	SuggestedMatches []PersonMatch   `json:"suggested_matches"`
	Created          bool            `json:"created"`
}

func (w *WorkspacePersonResult) GetPerson() WorkspacePerson {
	if w == nil {
		return WorkspacePerson{}
	}
	return w.Person
}

func (w *WorkspacePersonResult) GetSuggestedMatches() []PersonMatch {
	if w == nil {
		return []PersonMatch{}
	}
	return w.SuggestedMatches
}

func (w *WorkspacePersonResult) GetCreated() bool {
	if w == nil {
		return false
	}
	return w.Created
}
