package components

type PersonMatchMatchedOn string

const (
	PersonMatchMatchedOnEmail PersonMatchMatchedOn = "email"
	PersonMatchMatchedOnPhone PersonMatchMatchedOn = "phone"
)

func (e PersonMatchMatchedOn) ToPointer() *PersonMatchMatchedOn {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PersonMatchMatchedOn) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "phone":
			return true
		}
	}
	return false
}

type PersonMatch struct {
	PersonID  string                 `json:"person_id"`
	MatchedOn []PersonMatchMatchedOn `json:"matched_on"`
}

func (p *PersonMatch) GetPersonID() string {
	if p == nil {
		return ""
	}
	return p.PersonID
}

func (p *PersonMatch) GetMatchedOn() []PersonMatchMatchedOn {
	if p == nil {
		return []PersonMatchMatchedOn{}
	}
	return p.MatchedOn
}
