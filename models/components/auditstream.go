package components

type AuditStream struct {
	EndpointID  string      `json:"endpoint_id"`
	Environment Environment `json:"environment"`
	Actions     []string    `json:"actions"`
	CreatedAt   string      `json:"created_at"`
	// unverified until the endpoint echoes an endpoint.verification challenge (send one with POST /v1/webhooks/{endpoint_id}/test); verified receives assessment.updated events; disabled after 72 hours of continuous delivery failure, and verified again by the same challenge.
	EndpointState   WebhookEndpointState `json:"endpoint_state"`
	LastAttemptAt   *string              `json:"last_attempt_at,omitzero"`
	LastDeliveredAt *string              `json:"last_delivered_at,omitzero"`
}

func (a *AuditStream) GetEndpointID() string {
	if a == nil {
		return ""
	}
	return a.EndpointID
}

func (a *AuditStream) GetEnvironment() Environment {
	if a == nil {
		return Environment("")
	}
	return a.Environment
}

func (a *AuditStream) GetActions() []string {
	if a == nil {
		return []string{}
	}
	return a.Actions
}

func (a *AuditStream) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AuditStream) GetEndpointState() WebhookEndpointState {
	if a == nil {
		return WebhookEndpointState("")
	}
	return a.EndpointState
}

func (a *AuditStream) GetLastAttemptAt() *string {
	if a == nil {
		return nil
	}
	return a.LastAttemptAt
}

func (a *AuditStream) GetLastDeliveredAt() *string {
	if a == nil {
		return nil
	}
	return a.LastDeliveredAt
}
