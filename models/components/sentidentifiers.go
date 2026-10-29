package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// SentIdentifiers - The identifiers an assessment was sent: each member of the request's identifiers object exactly as received, before the server normalised it (the email domain keeps its case, the name its Unicode form). A member the request did not send is absent. NoResponse keeps them encrypted under the workspace's own key, does not index or log them, and deletes them with the assessment: on DELETE of the assessment or its account, at its expires_at, and with the workspace.
type SentIdentifiers struct {
	// identifiers.email as the request sent it.
	Email *string `json:"email,omitzero"`
	// identifiers.username as the request sent it.
	Username *string `json:"username,omitzero"`
	// identifiers.phone as the request sent it.
	Phone *string `json:"phone,omitzero"`
	// identifiers.name as the request sent it, before NFC normalisation.
	Name *string `json:"name,omitzero"`
	// identifiers.ip as the request sent it.
	IP *string `json:"ip,omitzero"`
}

func (s SentIdentifiers) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SentIdentifiers) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *SentIdentifiers) GetEmail() *string {
	if s == nil {
		return nil
	}
	return s.Email
}

func (s *SentIdentifiers) GetUsername() *string {
	if s == nil {
		return nil
	}
	return s.Username
}

func (s *SentIdentifiers) GetPhone() *string {
	if s == nil {
		return nil
	}
	return s.Phone
}

func (s *SentIdentifiers) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SentIdentifiers) GetIP() *string {
	if s == nil {
		return nil
	}
	return s.IP
}
