package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// PersonFlagged - A screening recorded a finding on a person. It says which person and how severe, not what they posted, who they are or why: GET the person report for the findings. Sent once per screening that flags the person, so a re-screen of someone on watch announces them again.
type PersonFlagged struct {
	// The person's group.
	GroupID string `json:"group_id"`
	// The flagged person.
	PersonID string `json:"person_id"`
	// The screening that flagged them.
	ScreeningID string      `json:"screening_id"`
	HighestBand HighestBand `json:"highest_band"`
	// How sure the screening is that the accounts it read are the person's. A finding on a possible match is at most moderate, on a probable match at most high.
	MatchState PersonFlaggedMatchState `json:"match_state"`
	// When the screening recorded the finding; equal to the event's created_at.
	FlaggedAt string `json:"flagged_at"`
}

func (p PersonFlagged) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonFlagged) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonFlagged) GetGroupID() string {
	if p == nil {
		return ""
	}
	return p.GroupID
}

func (p *PersonFlagged) GetPersonID() string {
	if p == nil {
		return ""
	}
	return p.PersonID
}

func (p *PersonFlagged) GetScreeningID() string {
	if p == nil {
		return ""
	}
	return p.ScreeningID
}

func (p *PersonFlagged) GetHighestBand() HighestBand {
	if p == nil {
		return HighestBand("")
	}
	return p.HighestBand
}

func (p *PersonFlagged) GetMatchState() PersonFlaggedMatchState {
	if p == nil {
		return PersonFlaggedMatchState("")
	}
	return p.MatchState
}

func (p *PersonFlagged) GetFlaggedAt() string {
	if p == nil {
		return ""
	}
	return p.FlaggedAt
}
