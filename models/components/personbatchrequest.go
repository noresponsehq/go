package components

// PersonBatchRequest - Adds or updates 1 to 500 people. A row naming an existing person of the group (the same external_id, else the same email, else the same phone) updates it.
type PersonBatchRequest struct {
	People []PersonIdentity `json:"people"`
}

func (p *PersonBatchRequest) GetPeople() []PersonIdentity {
	if p == nil {
		return []PersonIdentity{}
	}
	return p.People
}
