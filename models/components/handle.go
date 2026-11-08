package components

// Handle - An account the attendee is known to hold.
type Handle struct {
	Platform HandlePlatform `json:"platform"`
	// The account's handle, without a leading @.
	Handle string `json:"handle"`
}

func (h *Handle) GetPlatform() HandlePlatform {
	if h == nil {
		return HandlePlatform("")
	}
	return h.Platform
}

func (h *Handle) GetHandle() string {
	if h == nil {
		return ""
	}
	return h.Handle
}
