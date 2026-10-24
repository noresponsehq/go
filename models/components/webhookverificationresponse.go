package components

// WebhookVerificationResponse - The body an endpoint returns, with status 200 and Content-Type application/json within 5 seconds, to an endpoint.verification event. A 2xx without the echoed challenge does not verify the endpoint.
type WebhookVerificationResponse struct {
	// Endpoint verification nonce: 32 random bytes, unpadded base64url (43 characters). Fresh per verification.
	Challenge string `json:"challenge"`
}

func (w *WebhookVerificationResponse) GetChallenge() string {
	if w == nil {
		return ""
	}
	return w.Challenge
}
