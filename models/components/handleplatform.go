package components

// HandlePlatform - A platform a screening can cover. github and scratch through flickr are each one read of the public profile a handle names, with no posts. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.
type HandlePlatform string

const (
	HandlePlatformInstagram       HandlePlatform = "instagram"
	HandlePlatformX               HandlePlatform = "x"
	HandlePlatformBluesky         HandlePlatform = "bluesky"
	HandlePlatformYoutube         HandlePlatform = "youtube"
	HandlePlatformTiktok          HandlePlatform = "tiktok"
	HandlePlatformFacebook        HandlePlatform = "facebook"
	HandlePlatformThreads         HandlePlatform = "threads"
	HandlePlatformReddit          HandlePlatform = "reddit"
	HandlePlatformTelegram        HandlePlatform = "telegram"
	HandlePlatformLinkedin        HandlePlatform = "linkedin"
	HandlePlatformGithub          HandlePlatform = "github"
	HandlePlatformScratch         HandlePlatform = "scratch"
	HandlePlatformDailymotion     HandlePlatform = "dailymotion"
	HandlePlatformSubstack        HandlePlatform = "substack"
	HandlePlatformDuolingo        HandlePlatform = "duolingo"
	HandlePlatformDockerhub       HandlePlatform = "dockerhub"
	HandlePlatformWikipedia       HandlePlatform = "wikipedia"
	HandlePlatformInternetarchive HandlePlatform = "internetarchive"
	HandlePlatformSpotify         HandlePlatform = "spotify"
	HandlePlatformPatreon         HandlePlatform = "patreon"
	HandlePlatformWordpress       HandlePlatform = "wordpress"
	HandlePlatformOkru            HandlePlatform = "okru"
	HandlePlatformFlickr          HandlePlatform = "flickr"
	HandlePlatformSynthetic       HandlePlatform = "synthetic"
)

func (e HandlePlatform) ToPointer() *HandlePlatform {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *HandlePlatform) IsExact() bool {
	if e != nil {
		switch *e {
		case "instagram", "x", "bluesky", "youtube", "tiktok", "facebook", "threads", "reddit", "telegram", "linkedin", "github", "scratch", "dailymotion", "substack", "duolingo", "dockerhub", "wikipedia", "internetarchive", "spotify", "patreon", "wordpress", "okru", "flickr", "synthetic":
			return true
		}
	}
	return false
}
