package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AttendeeDecided - A reviewer recorded a decision on an attendee. The decision's note stays in the API. Sent once per recorded review; an attendee's latest review is its decision, so order by decided_at.
type AttendeeDecided struct {
	// The attendee's event.
	EventID string `json:"event_id"`
	// The attendee decided on.
	AttendeeID string `json:"attendee_id"`
	// The recorded review.
	ReviewID string           `json:"review_id"`
	Decision AttendeeDecision `json:"decision"`
	// Who decided: the member's external user id, or the API key's label, as the attendee's decided_by.
	DecidedBy string `json:"decided_by"`
	// When the decision was recorded; equal to the event's created_at.
	DecidedAt string `json:"decided_at"`
}

func (a AttendeeDecided) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttendeeDecided) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttendeeDecided) GetEventID() string {
	if a == nil {
		return ""
	}
	return a.EventID
}

func (a *AttendeeDecided) GetAttendeeID() string {
	if a == nil {
		return ""
	}
	return a.AttendeeID
}

func (a *AttendeeDecided) GetReviewID() string {
	if a == nil {
		return ""
	}
	return a.ReviewID
}

func (a *AttendeeDecided) GetDecision() AttendeeDecision {
	if a == nil {
		return AttendeeDecision("")
	}
	return a.Decision
}

func (a *AttendeeDecided) GetDecidedBy() string {
	if a == nil {
		return ""
	}
	return a.DecidedBy
}

func (a *AttendeeDecided) GetDecidedAt() string {
	if a == nil {
		return ""
	}
	return a.DecidedAt
}
