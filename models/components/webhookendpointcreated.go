package components

// WebhookEndpointCreated - The 201 answer of POST /v1/webhooks: the new endpoint, unverified, and its first signing secret, shown once. Store the secret where the receiver reads it, then send endpoint.verification with POST /v1/webhooks/{endpoint_id}/test.
type WebhookEndpointCreated struct {
	// One of the tenant's registered webhook endpoints. Signing secrets are not part of it: a secret is shown once, when it is issued (POST /v1/webhooks and POST /v1/webhooks/{endpoint_id}/rotate).
	Endpoint WebhookEndpoint `json:"endpoint"`
	// A newly issued signing key with its secret, shown only in this answer. A receiver verifies NR-Webhook-Signature with the secret's UTF-8 bytes as the HMAC-SHA256 key.
	SigningSecret SigningSecret `json:"signing_secret"`
}

func (w *WebhookEndpointCreated) GetEndpoint() WebhookEndpoint {
	if w == nil {
		return WebhookEndpoint{}
	}
	return w.Endpoint
}

func (w *WebhookEndpointCreated) GetSigningSecret() SigningSecret {
	if w == nil {
		return SigningSecret{}
	}
	return w.SigningSecret
}
