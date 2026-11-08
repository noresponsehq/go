package components

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
)

func (e HandlePlatform) ToPointer() *HandlePlatform {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *HandlePlatform) IsExact() bool {
	if e != nil {
		switch *e {
		case "instagram", "x", "bluesky", "youtube", "tiktok", "facebook", "threads", "reddit", "telegram", "linkedin":
			return true
		}
	}
	return false
}
