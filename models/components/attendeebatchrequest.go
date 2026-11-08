package components

// AttendeeBatchRequest - Adds or updates 1 to 500 attendees. A row naming an existing attendee of the event (the same external_id, else the same email, else the same phone) updates it.
type AttendeeBatchRequest struct {
	Attendees []AttendeeIdentity `json:"attendees"`
}

func (a *AttendeeBatchRequest) GetAttendees() []AttendeeIdentity {
	if a == nil {
		return []AttendeeIdentity{}
	}
	return a.Attendees
}
