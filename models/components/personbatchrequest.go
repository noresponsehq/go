package components

// PersonBatchRequest - Adds 1 to 500 workspace people to the group. Reuses an explicit person_id or external_id. Matching email or phone values only suggest people; automatic merging is disabled.
type PersonBatchRequest struct {
	People []PersonIdentity `json:"people"`
}

func (p *PersonBatchRequest) GetPeople() []PersonIdentity {
	if p == nil {
		return []PersonIdentity{}
	}
	return p.People
}
