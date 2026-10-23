package components

// WebhookEndpointRequest - POST /v1/webhooks body. The workspace is the API key that signs the request, so it is not a body member. The URL must be https on port 443 or 8443 with a DNS name, not an IP literal, and every address the name resolves to must be public.
type WebhookEndpointRequest struct {
	// Where deliveries are sent: an https URL on port 443 or 8443, with a DNS host and no userinfo or fragment, at most 2048 characters.
	URL string `json:"url"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode *ResponseMode `json:"response_mode,omitzero"`
}

func (w *WebhookEndpointRequest) GetURL() string {
	if w == nil {
		return ""
	}
	return w.URL
}

func (w *WebhookEndpointRequest) GetResponseMode() *ResponseMode {
	if w == nil {
		return nil
	}
	return w.ResponseMode
}
