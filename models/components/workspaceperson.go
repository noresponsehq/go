package components

import (
	"github.com/noresponsehq/go/v15/internal/utils"
)

type WorkspacePerson struct {
	// Your reference for this person within the workspace environment. Only an explicit reference or person_id can reuse an existing person.
	ExternalID *string `json:"external_id,omitzero"`
	Name       *string `json:"name,omitzero"`
	// The person's email.
	Email *string `json:"email,omitzero"`
	// The person's phone in E.164 form.
	Phone *string `json:"phone,omitzero"`
	// Where the person works.
	Company *string `json:"company,omitzero"`
	// Where the person lives.
	City *string `json:"city,omitzero"`
	// Accounts the person is known to hold, at most 20.
	Handles   []Handle `json:"handles,omitzero"`
	ID        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	// The time the workspace retention policy removes this person, their evidence and investigation history. Shortening workspace retention can move this time earlier.
	ExpiresAt string `json:"expires_at"`
}

func (w WorkspacePerson) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkspacePerson) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (w *WorkspacePerson) GetExternalID() *string {
	if w == nil {
		return nil
	}
	return w.ExternalID
}

func (w *WorkspacePerson) GetName() *string {
	if w == nil {
		return nil
	}
	return w.Name
}

func (w *WorkspacePerson) GetEmail() *string {
	if w == nil {
		return nil
	}
	return w.Email
}

func (w *WorkspacePerson) GetPhone() *string {
	if w == nil {
		return nil
	}
	return w.Phone
}

func (w *WorkspacePerson) GetCompany() *string {
	if w == nil {
		return nil
	}
	return w.Company
}

func (w *WorkspacePerson) GetCity() *string {
	if w == nil {
		return nil
	}
	return w.City
}

func (w *WorkspacePerson) GetHandles() []Handle {
	if w == nil {
		return nil
	}
	return w.Handles
}

func (w *WorkspacePerson) GetID() string {
	if w == nil {
		return ""
	}
	return w.ID
}

func (w *WorkspacePerson) GetCreatedAt() string {
	if w == nil {
		return ""
	}
	return w.CreatedAt
}

func (w *WorkspacePerson) GetUpdatedAt() string {
	if w == nil {
		return ""
	}
	return w.UpdatedAt
}

func (w *WorkspacePerson) GetExpiresAt() string {
	if w == nil {
		return ""
	}
	return w.ExpiresAt
}
