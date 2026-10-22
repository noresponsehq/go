package components

// SigningKey - One signing key, by id. The secret is not included.
type SigningKey struct {
	// The signing key id, sent as NR-Webhook-Key-Id on every delivery it signs.
	KeyID     string `json:"key_id"`
	CreatedAt string `json:"created_at"`
	// Present on a key a rotation replaced: it stops verifying at this instant, 7 days after the rotation. Absent on the key that signs.
	ExpiresAt *string `json:"expires_at,omitzero"`
}

func (s *SigningKey) GetKeyID() string {
	if s == nil {
		return ""
	}
	return s.KeyID
}

func (s *SigningKey) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}

func (s *SigningKey) GetExpiresAt() *string {
	if s == nil {
		return nil
	}
	return s.ExpiresAt
}
