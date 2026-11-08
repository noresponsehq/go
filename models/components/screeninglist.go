package components

// ScreeningList - Screenings, newest first.
type ScreeningList struct {
	Items []Screening `json:"items"`
	// Pass as before to read the next, older page.
	NextBefore *string `json:"next_before,omitzero"`
}

func (s *ScreeningList) GetItems() []Screening {
	if s == nil {
		return []Screening{}
	}
	return s.Items
}

func (s *ScreeningList) GetNextBefore() *string {
	if s == nil {
		return nil
	}
	return s.NextBefore
}
