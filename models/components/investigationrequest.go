package components

// InvestigationRequest - Runs a new investigation against one retained workspace person.
type InvestigationRequest struct {
	Depth     ScreeningDepth `json:"depth"`
	Platforms []string       `json:"platforms"`
	PersonID  string         `json:"person_id"`
}

func (i *InvestigationRequest) GetDepth() ScreeningDepth {
	if i == nil {
		return ScreeningDepth("")
	}
	return i.Depth
}

func (i *InvestigationRequest) GetPlatforms() []string {
	if i == nil {
		return []string{}
	}
	return i.Platforms
}

func (i *InvestigationRequest) GetPersonID() string {
	if i == nil {
		return ""
	}
	return i.PersonID
}
