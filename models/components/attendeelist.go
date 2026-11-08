package components

// AttendeeList - Attendees in upload order.
type AttendeeList struct {
	Items []Attendee `json:"items"`
	// Pass as after to read the next page.
	NextAfter *string `json:"next_after,omitzero"`
}

func (a *AttendeeList) GetItems() []Attendee {
	if a == nil {
		return []Attendee{}
	}
	return a.Items
}

func (a *AttendeeList) GetNextAfter() *string {
	if a == nil {
		return nil
	}
	return a.NextAfter
}
