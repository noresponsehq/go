package components

// SigningSecret - A newly issued signing key with its secret, shown only in this answer. A receiver verifies NR-Webhook-Signature with the secret's UTF-8 bytes as the HMAC-SHA256 key.
type SigningSecret struct {
	// The signing key id, sent as NR-Webhook-Key-Id on every delivery it signs.
	KeyID string `json:"key_id"`
	// The signing secret: 32 random bytes as 43 characters of unpadded base64url. It is not stored in readable form or returned again.
	Secret    string `json:"secret"`
	CreatedAt string `json:"created_at"`
}

func (s *SigningSecret) GetKeyID() string {
	if s == nil {
		return ""
	}
	return s.KeyID
}

func (s *SigningSecret) GetSecret() string {
	if s == nil {
		return ""
	}
	return s.Secret
}

func (s *SigningSecret) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}
