package components

// WebhookEndpoint - One of the workspace's registered webhook endpoints. Signing secrets are not part of it: a secret is shown once, when it is issued (POST /v1/webhooks and POST /v1/webhooks/{endpoint_id}/rotate).
type WebhookEndpoint struct {
	// Opaque id of the endpoint registration.
	EndpointID string `json:"endpoint_id"`
	// The destination URL, as registered.
	URL string `json:"url"`
	// unverified until the endpoint echoes an endpoint.verification challenge (send one with POST /v1/webhooks/{endpoint_id}/test); verified receives assessment.updated events; disabled after 72 hours of continuous delivery failure, and verified again by the same challenge.
	State WebhookEndpointState `json:"state"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode ResponseMode `json:"response_mode"`
	CreatedAt    string       `json:"created_at"`
	// When the endpoint last echoed a challenge. Absent until it first does.
	VerifiedAt *string `json:"verified_at,omitzero"`
	// The first failed delivery attempt of the endpoint's current failure streak. Absent while its latest attempt succeeded, before any attempt failed, and after it is verified again.
	FailingSince *string `json:"failing_since,omitzero"`
	// The latest failed delivery attempt of the current failure streak; present exactly when failing_since is. A failure more than 2 hours after it starts a new streak.
	LastFailureAt *string `json:"last_failure_at,omitzero"`
	// When the endpoint was disabled after failing continuously for 72 hours, as in the webhook_endpoint.disabled event. Present while state is disabled; absent before and after it is verified again.
	DisabledAt *string `json:"disabled_at,omitzero"`
	// The keys a receiver may see in NR-Webhook-Key-Id, oldest first. The last one signs every new delivery; an older one keeps verifying until its expires_at.
	SigningKeys []SigningKey `json:"signing_keys"`
}

func (w *WebhookEndpoint) GetEndpointID() string {
	if w == nil {
		return ""
	}
	return w.EndpointID
}

func (w *WebhookEndpoint) GetURL() string {
	if w == nil {
		return ""
	}
	return w.URL
}

func (w *WebhookEndpoint) GetState() WebhookEndpointState {
	if w == nil {
		return WebhookEndpointState("")
	}
	return w.State
}

func (w *WebhookEndpoint) GetResponseMode() ResponseMode {
	if w == nil {
		return ResponseMode("")
	}
	return w.ResponseMode
}

func (w *WebhookEndpoint) GetCreatedAt() string {
	if w == nil {
		return ""
	}
	return w.CreatedAt
}

func (w *WebhookEndpoint) GetVerifiedAt() *string {
	if w == nil {
		return nil
	}
	return w.VerifiedAt
}

func (w *WebhookEndpoint) GetFailingSince() *string {
	if w == nil {
		return nil
	}
	return w.FailingSince
}

func (w *WebhookEndpoint) GetLastFailureAt() *string {
	if w == nil {
		return nil
	}
	return w.LastFailureAt
}

func (w *WebhookEndpoint) GetDisabledAt() *string {
	if w == nil {
		return nil
	}
	return w.DisabledAt
}

func (w *WebhookEndpoint) GetSigningKeys() []SigningKey {
	if w == nil {
		return []SigningKey{}
	}
	return w.SigningKeys
}
