package components

import (
	"github.com/noresponsehq/go/v18/internal/utils"
)

type PersonIdentity struct {
	// Your reference for this person within the workspace environment. Only an explicit reference or person_id can reuse an existing person.
	ExternalID *string `json:"external_id,omitzero"`
	// The person's full name.
	Name *string `json:"name,omitzero"`
	// The person's email.
	Email *string `json:"email,omitzero"`
	// The person's phone in E.164 form.
	Phone *string `json:"phone,omitzero"`
	// Where the person works.
	Company *string `json:"company,omitzero"`
	// Where the person lives.
	City     *string  `json:"city,omitzero"`
	Handles  []Handle `json:"handles,omitzero"`
	PersonID *string  `json:"person_id,omitzero"`
	// The person's role in the group: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: a person without one has no role, not guest. A later row naming the same person replaces it, and a change of role alone keeps the person's screening results.
	Role *PersonRole `json:"role,omitzero"`
}

func (p PersonIdentity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonIdentity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PersonIdentity) GetExternalID() *string {
	if p == nil {
		return nil
	}
	return p.ExternalID
}

func (p *PersonIdentity) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PersonIdentity) GetEmail() *string {
	if p == nil {
		return nil
	}
	return p.Email
}

func (p *PersonIdentity) GetPhone() *string {
	if p == nil {
		return nil
	}
	return p.Phone
}

func (p *PersonIdentity) GetCompany() *string {
	if p == nil {
		return nil
	}
	return p.Company
}

func (p *PersonIdentity) GetCity() *string {
	if p == nil {
		return nil
	}
	return p.City
}

func (p *PersonIdentity) GetHandles() []Handle {
	if p == nil {
		return nil
	}
	return p.Handles
}

func (p *PersonIdentity) GetPersonID() *string {
	if p == nil {
		return nil
	}
	return p.PersonID
}

func (p *PersonIdentity) GetRole() *PersonRole {
	if p == nil {
		return nil
	}
	return p.Role
}
