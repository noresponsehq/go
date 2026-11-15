package components

// Handle - An account the person is known to hold.
type Handle struct {
	// A platform a screening can cover, such as instagram, x, bluesky, youtube, tiktok or github. The screening registry owns the set: a request that names a platform it does not cover answers 422. A new platform adds an id without changing this contract, so read a platform you do not know in a response as one you have not seen before, not as an error. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
	Platform string `json:"platform"`
	// The account's handle, without a leading @.
	Handle string `json:"handle"`
}

func (h *Handle) GetPlatform() string {
	if h == nil {
		return ""
	}
	return h.Platform
}

func (h *Handle) GetHandle() string {
	if h == nil {
		return ""
	}
	return h.Handle
}
