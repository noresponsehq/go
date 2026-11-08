package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AttendeeIdentity - What you know about an attendee. Stored encrypted under the workspace's key.
type AttendeeIdentity struct {
	// Your own id for the attendee, unique within the event. A later row with the same external id updates this attendee.
	ExternalID *string `json:"external_id,omitzero"`
	// The attendee's full name.
	Name string `json:"name"`
	// The attendee's email.
	Email *string `json:"email,omitzero"`
	// The attendee's phone in E.164 form.
	Phone *string `json:"phone,omitzero"`
	// Where the attendee works.
	Company *string `json:"company,omitzero"`
	// Where the attendee lives.
	City *string `json:"city,omitzero"`
	// Accounts the attendee is known to hold, at most 20.
	Handles []Handle `json:"handles,omitzero"`
}

func (a AttendeeIdentity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttendeeIdentity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AttendeeIdentity) GetExternalID() *string {
	if a == nil {
		return nil
	}
	return a.ExternalID
}

func (a *AttendeeIdentity) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *AttendeeIdentity) GetEmail() *string {
	if a == nil {
		return nil
	}
	return a.Email
}

func (a *AttendeeIdentity) GetPhone() *string {
	if a == nil {
		return nil
	}
	return a.Phone
}

func (a *AttendeeIdentity) GetCompany() *string {
	if a == nil {
		return nil
	}
	return a.Company
}

func (a *AttendeeIdentity) GetCity() *string {
	if a == nil {
		return nil
	}
	return a.City
}

func (a *AttendeeIdentity) GetHandles() []Handle {
	if a == nil {
		return nil
	}
	return a.Handles
}
