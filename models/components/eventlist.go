package components

// EventList - Events, newest first.
type EventList struct {
	Items []Event `json:"items"`
	// Pass as before to read the next, older page.
	NextBefore *string `json:"next_before,omitzero"`
}

func (e *EventList) GetItems() []Event {
	if e == nil {
		return []Event{}
	}
	return e.Items
}

func (e *EventList) GetNextBefore() *string {
	if e == nil {
		return nil
	}
	return e.NextBefore
}
