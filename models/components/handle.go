package components

// Handle - An account the person is known to hold.
type Handle struct {
	// A platform a screening can cover. github and scratch through flickr are each one read of the public profile a handle names, with no posts. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
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
