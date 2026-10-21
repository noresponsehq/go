package components

// KeyRotationRequest - POST /v1/keys/rotate body. The tenant and the key being rotated are the API key that signs the request, so neither is a body member. An empty object takes the default overlap.
type KeyRotationRequest struct {
	// How long the presented key keeps working beside the new one, from 60 seconds to 7 days. Default 86400 (24 hours). At the end of the overlap the presented key stops authenticating.
	OverlapSeconds *int64 `json:"overlap_seconds,omitzero"`
}

func (k *KeyRotationRequest) GetOverlapSeconds() *int64 {
	if k == nil {
		return nil
	}
	return k.OverlapSeconds
}
