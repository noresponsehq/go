package components

// HandlePlatform - A platform a screening can cover. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
type HandlePlatform string

const (
	HandlePlatformInstagram HandlePlatform = "instagram"
	HandlePlatformX         HandlePlatform = "x"
	HandlePlatformBluesky   HandlePlatform = "bluesky"
	HandlePlatformYoutube   HandlePlatform = "youtube"
	HandlePlatformTiktok    HandlePlatform = "tiktok"
	HandlePlatformFacebook  HandlePlatform = "facebook"
	HandlePlatformThreads   HandlePlatform = "threads"
	HandlePlatformReddit    HandlePlatform = "reddit"
	HandlePlatformTelegram  HandlePlatform = "telegram"
	HandlePlatformLinkedin  HandlePlatform = "linkedin"
	HandlePlatformSynthetic HandlePlatform = "synthetic"
)

func (e HandlePlatform) ToPointer() *HandlePlatform {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *HandlePlatform) IsExact() bool {
	if e != nil {
		switch *e {
		case "instagram", "x", "bluesky", "youtube", "tiktok", "facebook", "threads", "reddit", "telegram", "linkedin", "synthetic":
			return true
		}
	}
	return false
}
