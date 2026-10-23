package components

// KeyUsageBatchRequest - POST /v1/key_usage/batches body: facts about requests made with API keys your signup gate issued to your accounts, one event per request. Each event names the key by the SHA-256 that the assessment's issued_credential.sha256 carried, not the key itself. The SDKs buffer events and send them in the background. A batch holds 1 to 500 events and at most 256 KiB. Idempotency-Key is required and binds the body as on POST /v1/assessments, so a retried batch is counted once.
type KeyUsageBatchRequest struct {
	Events []KeyUsageEvent `json:"events"`
}

func (k *KeyUsageBatchRequest) GetEvents() []KeyUsageEvent {
	if k == nil {
		return []KeyUsageEvent{}
	}
	return k.Events
}
