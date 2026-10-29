package components

// ConsoleLookups - The live environment's daily cap on assessments members make through the console (the playground), as spent on the current UTC day. Each console submission with a new Idempotency-Key spends one lookup; a replay spends nothing, and API keys and the test environment spend none. When used reaches limit, the next lookup answers 429 lookup_limit_reached until resets_at. The count is the one the cap enforces.
type ConsoleLookups struct {
	// Lookups spent since the last reset.
	Used int64 `json:"used"`
	// Lookups allowed per UTC day.
	Limit int64 `json:"limit"`
	// The next reset: the coming midnight UTC, when used returns to 0.
	ResetsAt string `json:"resets_at"`
}

func (c *ConsoleLookups) GetUsed() int64 {
	if c == nil {
		return 0
	}
	return c.Used
}

func (c *ConsoleLookups) GetLimit() int64 {
	if c == nil {
		return 0
	}
	return c.Limit
}

func (c *ConsoleLookups) GetResetsAt() string {
	if c == nil {
		return ""
	}
	return c.ResetsAt
}
