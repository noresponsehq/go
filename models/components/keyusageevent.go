package components

import (
	"github.com/noresponsehq/go/v3/internal/utils"
)

// KeyUsageEvent - One request made with an issued key. An event whose ip is not a public address, or whose occurred_at is more than 24 hours old or more than 5 minutes ahead, is counted in ignored and adds nothing.
type KeyUsageEvent struct {
	// SHA-256 of the issued key's UTF-8 bytes, as 64 lower-case hex characters: the value sent as issued_credential.sha256 when the account was assessed. Each SDK's issued-credential helper computes it.
	KeySha256 string `json:"key_sha256"`
	// The client IP address the request came from, IPv4 or IPv6, as your server saw it. Stored only as a keyed digest and its country.
	IP string `json:"ip"`
	// When the request was made.
	OccurredAt string `json:"occurred_at"`
	// The request's User-Agent header, at most 512 characters. Stored only as a keyed digest.
	UserAgent *string `json:"user_agent,omitzero"`
	// The model or product the request used, such as gpt-4o-mini. Stored only as a keyed digest.
	Model *string `json:"model,omitzero"`
	// Input tokens the request consumed.
	InputTokens *int64 `json:"input_tokens,omitzero"`
	// Output tokens the request produced.
	OutputTokens *int64 `json:"output_tokens,omitzero"`
	// The share of requests your sampler records, when it records only some: 0.1 means this event stands for 10 requests. Request and token totals are weighted by its inverse. Distinct IP and country counts see only the events sent.
	SampleRate *float64 `default:"1" json:"sample_rate"`
}

func (k KeyUsageEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyUsageEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (k *KeyUsageEvent) GetKeySha256() string {
	if k == nil {
		return ""
	}
	return k.KeySha256
}

func (k *KeyUsageEvent) GetIP() string {
	if k == nil {
		return ""
	}
	return k.IP
}

func (k *KeyUsageEvent) GetOccurredAt() string {
	if k == nil {
		return ""
	}
	return k.OccurredAt
}

func (k *KeyUsageEvent) GetUserAgent() *string {
	if k == nil {
		return nil
	}
	return k.UserAgent
}

func (k *KeyUsageEvent) GetModel() *string {
	if k == nil {
		return nil
	}
	return k.Model
}

func (k *KeyUsageEvent) GetInputTokens() *int64 {
	if k == nil {
		return nil
	}
	return k.InputTokens
}

func (k *KeyUsageEvent) GetOutputTokens() *int64 {
	if k == nil {
		return nil
	}
	return k.OutputTokens
}

func (k *KeyUsageEvent) GetSampleRate() *float64 {
	if k == nil {
		return nil
	}
	return k.SampleRate
}
