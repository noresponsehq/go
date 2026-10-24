package components

import (
	"github.com/noresponsehq/go/v6/internal/utils"
)

type EndpointVerification struct {
	// Opaque id of the endpoint registration being verified.
	EndpointID string `json:"endpoint_id"`
	// Endpoint verification nonce: 32 random bytes, unpadded base64url (43 characters). Fresh per verification.
	Challenge string `json:"challenge"`
}

func (e EndpointVerification) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EndpointVerification) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (e *EndpointVerification) GetEndpointID() string {
	if e == nil {
		return ""
	}
	return e.EndpointID
}

func (e *EndpointVerification) GetChallenge() string {
	if e == nil {
		return ""
	}
	return e.Challenge
}
