package components

// KeyUsageBatch - The 202 answer of POST /v1/key_usage/batches. A replay with the same Idempotency-Key and body answers the same record and counts nothing again.
type KeyUsageBatch struct {
	// Opaque, server-assigned id of the batch: 'kub_' and a 26-character Crockford base32 ULID.
	ID string `json:"id"`
	// When NoResponse first recorded the batch.
	ReceivedAt string `json:"received_at"`
	// Events counted.
	Accepted int64 `json:"accepted"`
	// Events that added nothing: a private, reserved or unreadable ip, or an occurred_at more than 24 hours old or more than 5 minutes ahead.
	Ignored int64 `json:"ignored"`
}

func (k *KeyUsageBatch) GetID() string {
	if k == nil {
		return ""
	}
	return k.ID
}

func (k *KeyUsageBatch) GetReceivedAt() string {
	if k == nil {
		return ""
	}
	return k.ReceivedAt
}

func (k *KeyUsageBatch) GetAccepted() int64 {
	if k == nil {
		return 0
	}
	return k.Accepted
}

func (k *KeyUsageBatch) GetIgnored() int64 {
	if k == nil {
		return 0
	}
	return k.Ignored
}
