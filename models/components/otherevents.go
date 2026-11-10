package components

// OtherEvents - Where else in the environment the attendee appears, read live from those events: deleting one removes it here.
type OtherEvents struct {
	// How many attendees of other events match, of which items lists the 20 most recent.
	Total int64 `json:"total"`
	// The 20 attendees of the most recent other events, newest event first.
	Items []OtherEvent `json:"items"`
}

func (o *OtherEvents) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}

func (o *OtherEvents) GetItems() []OtherEvent {
	if o == nil {
		return []OtherEvent{}
	}
	return o.Items
}
