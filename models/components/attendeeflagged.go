package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AttendeeFlagged - A screening recorded a finding on an attendee. It says which attendee and how severe, not what they posted, who they are or why: GET the attendee report for the findings. Sent once per screening that flags the attendee, so a re-screen of someone on watch announces them again.
type AttendeeFlagged struct {
	// The attendee's event.
	EventID string `json:"event_id"`
	// The flagged attendee.
	AttendeeID string `json:"attendee_id"`
	// The screening that flagged them.
	ScreeningID string      `json:"screening_id"`
	HighestBand HighestBand `json:"highest_band"`
	// How sure the screening is that the accounts it read are the attendee's. A finding on a possible match is at most moderate, on a probable match at most high.
	MatchState AttendeeFlaggedMatchState `json:"match_state"`
	// When the screening recorded the finding; equal to the event's created_at.
	FlaggedAt string `json:"flagged_at"`
}

func (a AttendeeFlagged) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttendeeFlagged) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttendeeFlagged) GetEventID() string {
	if a == nil {
		return ""
	}
	return a.EventID
}

func (a *AttendeeFlagged) GetAttendeeID() string {
	if a == nil {
		return ""
	}
	return a.AttendeeID
}

func (a *AttendeeFlagged) GetScreeningID() string {
	if a == nil {
		return ""
	}
	return a.ScreeningID
}

func (a *AttendeeFlagged) GetHighestBand() HighestBand {
	if a == nil {
		return HighestBand("")
	}
	return a.HighestBand
}

func (a *AttendeeFlagged) GetMatchState() AttendeeFlaggedMatchState {
	if a == nil {
		return AttendeeFlaggedMatchState("")
	}
	return a.MatchState
}

func (a *AttendeeFlagged) GetFlaggedAt() string {
	if a == nil {
		return ""
	}
	return a.FlaggedAt
}
