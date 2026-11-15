package components

// SourceID - A NoResponse source: github (GitHub accounts), phone (the phone lookup), breach (breach exposure), rdap (the email domain's registration), mx (the email domain's mail servers), ipintel (IP intelligence) or an email account source that searches one platform for an account by email address (adobe, nvidia, wix, life360, zapier, google). google runs only for Gmail addresses. scratch, dailymotion, substack, duolingo, dockerhub, wikipedia, internetarchive, telegram, spotify, patreon, wordpress, okru and flickr each search that platform for an account by username.
type SourceID string

const (
	SourceIDGithub          SourceID = "github"
	SourceIDPhone           SourceID = "phone"
	SourceIDBreach          SourceID = "breach"
	SourceIDRdap            SourceID = "rdap"
	SourceIDMx              SourceID = "mx"
	SourceIDIpintel         SourceID = "ipintel"
	SourceIDAdobe           SourceID = "adobe"
	SourceIDNvidia          SourceID = "nvidia"
	SourceIDWix             SourceID = "wix"
	SourceIDLife360         SourceID = "life360"
	SourceIDZapier          SourceID = "zapier"
	SourceIDGoogle          SourceID = "google"
	SourceIDScratch         SourceID = "scratch"
	SourceIDDailymotion     SourceID = "dailymotion"
	SourceIDSubstack        SourceID = "substack"
	SourceIDDuolingo        SourceID = "duolingo"
	SourceIDDockerhub       SourceID = "dockerhub"
	SourceIDWikipedia       SourceID = "wikipedia"
	SourceIDInternetarchive SourceID = "internetarchive"
	SourceIDTelegram        SourceID = "telegram"
	SourceIDSpotify         SourceID = "spotify"
	SourceIDPatreon         SourceID = "patreon"
	SourceIDWordpress       SourceID = "wordpress"
	SourceIDOkru            SourceID = "okru"
	SourceIDFlickr          SourceID = "flickr"
)

func (e SourceID) ToPointer() *SourceID {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SourceID) IsExact() bool {
	if e != nil {
		switch *e {
		case "github", "phone", "breach", "rdap", "mx", "ipintel", "adobe", "nvidia", "wix", "life360", "zapier", "google", "scratch", "dailymotion", "substack", "duolingo", "dockerhub", "wikipedia", "internetarchive", "telegram", "spotify", "patreon", "wordpress", "okru", "flickr":
			return true
		}
	}
	return false
}
