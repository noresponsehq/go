package components

import (
	"github.com/noresponsehq/go/v11/internal/utils"
)

// PersonDecided - A reviewer recorded a decision on a person. The decision's note stays in the API. Sent once per recorded review; a person's latest review is its decision, so order by decided_at.
type PersonDecided struct {
	// The person's group.
	GroupID string `json:"group_id"`
	// The person decided on.
	PersonID string `json:"person_id"`
	// The recorded review.
	ReviewID string         `json:"review_id"`
	Decision PersonDecision `json:"decision"`
	// Who decided: the member's external user id, or the API key's label, as the person's decided_by.
	DecidedBy string `json:"decided_by"`
	// When the decision was recorded; equal to the event's created_at.
	DecidedAt string `json:"decided_at"`
}

func (p PersonDecided) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonDecided) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonDecided) GetGroupID() string {
	if p == nil {
		return ""
	}
	return p.GroupID
}

func (p *PersonDecided) GetPersonID() string {
	if p == nil {
		return ""
	}
	return p.PersonID
}

func (p *PersonDecided) GetReviewID() string {
	if p == nil {
		return ""
	}
	return p.ReviewID
}

func (p *PersonDecided) GetDecision() PersonDecision {
	if p == nil {
		return PersonDecision("")
	}
	return p.Decision
}

func (p *PersonDecided) GetDecidedBy() string {
	if p == nil {
		return ""
	}
	return p.DecidedBy
}

func (p *PersonDecided) GetDecidedAt() string {
	if p == nil {
		return ""
	}
	return p.DecidedAt
}
