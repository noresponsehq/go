package components

// WebhookKeyRotation - The 201 answer of POST /v1/webhooks/{endpoint_id}/rotate: the new signing secret, shown once, which signs every delivery from now on, and when the key it replaced stops verifying. Until then a receiver accepts both.
type WebhookKeyRotation struct {
	// Opaque id of the endpoint registration.
	EndpointID string `json:"endpoint_id"`
	// A newly issued signing key with its secret, shown only in this answer. A receiver verifies NR-Webhook-Signature with the secret's UTF-8 bytes as the HMAC-SHA256 key.
	SigningSecret SigningSecret `json:"signing_secret"`
	// The key that signed until this rotation.
	PreviousKeyID string `json:"previous_key_id"`
	// When the previous key stops verifying: 7 days after the rotation.
	PreviousKeyExpiresAt string `json:"previous_key_expires_at"`
}

func (w *WebhookKeyRotation) GetEndpointID() string {
	if w == nil {
		return ""
	}
	return w.EndpointID
}

func (w *WebhookKeyRotation) GetSigningSecret() SigningSecret {
	if w == nil {
		return SigningSecret{}
	}
	return w.SigningSecret
}

func (w *WebhookKeyRotation) GetPreviousKeyID() string {
	if w == nil {
		return ""
	}
	return w.PreviousKeyID
}

func (w *WebhookKeyRotation) GetPreviousKeyExpiresAt() string {
	if w == nil {
		return ""
	}
	return w.PreviousKeyExpiresAt
}
