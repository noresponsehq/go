package components

import (
	"github.com/noresponsehq/go/v17/internal/utils"
)

type WorkspacePersonRequest struct {
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
	City *string `json:"city,omitzero"`
	// A username the person uses, with no platform. A search looks it up on every platform that finds accounts by username.
	Username *string  `json:"username,omitzero"`
	Handles  []Handle `json:"handles,omitzero"`
	PersonID *string  `json:"person_id,omitzero"`
}

func (w WorkspacePersonRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspacePersonRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WorkspacePersonRequest) GetExternalID() *string {
	if w == nil {
		return nil
	}
	return w.ExternalID
}

func (w *WorkspacePersonRequest) GetName() *string {
	if w == nil {
		return nil
	}
	return w.Name
}

func (w *WorkspacePersonRequest) GetEmail() *string {
	if w == nil {
		return nil
	}
	return w.Email
}

func (w *WorkspacePersonRequest) GetPhone() *string {
	if w == nil {
		return nil
	}
	return w.Phone
}

func (w *WorkspacePersonRequest) GetCompany() *string {
	if w == nil {
		return nil
	}
	return w.Company
}

func (w *WorkspacePersonRequest) GetCity() *string {
	if w == nil {
		return nil
	}
	return w.City
}

func (w *WorkspacePersonRequest) GetUsername() *string {
	if w == nil {
		return nil
	}
	return w.Username
}

func (w *WorkspacePersonRequest) GetHandles() []Handle {
	if w == nil {
		return nil
	}
	return w.Handles
}

func (w *WorkspacePersonRequest) GetPersonID() *string {
	if w == nil {
		return nil
	}
	return w.PersonID
}
