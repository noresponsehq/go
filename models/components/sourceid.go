package components

// SourceID - A NoResponse source: github (GitHub accounts), phone (the phone lookup), breach (breach exposure), rdap (the email domain's registration), mx (the email domain's mail servers), ipintel (IP intelligence) or an email account source that searches one platform for an account by email address (adobe, nvidia, wix, life360, zapier, x_email, snapchat_email, amazon, spotify_email, google, duolingo_email). x_email, snapchat_email, amazon and spotify_email only say whether the email has an account on X, Snapchat, Amazon or Spotify. google runs only for Gmail addresses. scratch, dailymotion, substack, duolingo, dockerhub, wikipedia, internetarchive, telegram, spotify, patreon, wordpress, okru and flickr each search that platform for an account by username; telegram also searches by phone number, through a logged-in Telegram lookup account. whatsapp searches WhatsApp for an account by phone number. opplysningen (1881, Norway), ratsit (Sweden) and reverseaustralia (Australia) each look up a phone number of their own country on that site. youtube searches YouTube for a channel by username.
type SourceID string

const (
	SourceIDGithub           SourceID = "github"
	SourceIDPhone            SourceID = "phone"
	SourceIDBreach           SourceID = "breach"
	SourceIDRdap             SourceID = "rdap"
	SourceIDMx               SourceID = "mx"
	SourceIDIpintel          SourceID = "ipintel"
	SourceIDAdobe            SourceID = "adobe"
	SourceIDNvidia           SourceID = "nvidia"
	SourceIDWix              SourceID = "wix"
	SourceIDLife360          SourceID = "life360"
	SourceIDZapier           SourceID = "zapier"
	SourceIDXEmail           SourceID = "x_email"
	SourceIDSnapchatEmail    SourceID = "snapchat_email"
	SourceIDAmazon           SourceID = "amazon"
	SourceIDSpotifyEmail     SourceID = "spotify_email"
	SourceIDGoogle           SourceID = "google"
	SourceIDScratch          SourceID = "scratch"
	SourceIDDailymotion      SourceID = "dailymotion"
	SourceIDSubstack         SourceID = "substack"
	SourceIDDuolingo         SourceID = "duolingo"
	SourceIDDuolingoEmail    SourceID = "duolingo_email"
	SourceIDDockerhub        SourceID = "dockerhub"
	SourceIDWikipedia        SourceID = "wikipedia"
	SourceIDInternetarchive  SourceID = "internetarchive"
	SourceIDTelegram         SourceID = "telegram"
	SourceIDYoutube          SourceID = "youtube"
	SourceIDSpotify          SourceID = "spotify"
	SourceIDPatreon          SourceID = "patreon"
	SourceIDWordpress        SourceID = "wordpress"
	SourceIDOkru             SourceID = "okru"
	SourceIDFlickr           SourceID = "flickr"
	SourceIDWhatsapp         SourceID = "whatsapp"
	SourceIDOpplysningen     SourceID = "opplysningen"
	SourceIDRatsit           SourceID = "ratsit"
	SourceIDReverseaustralia SourceID = "reverseaustralia"
)

func (e SourceID) ToPointer() *SourceID {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SourceID) IsExact() bool {
	if e != nil {
		switch *e {
		case "github", "phone", "breach", "rdap", "mx", "ipintel", "adobe", "nvidia", "wix", "life360", "zapier", "x_email", "snapchat_email", "amazon", "spotify_email", "google", "scratch", "dailymotion", "substack", "duolingo", "duolingo_email", "dockerhub", "wikipedia", "internetarchive", "telegram", "youtube", "spotify", "patreon", "wordpress", "okru", "flickr", "whatsapp", "opplysningen", "ratsit", "reverseaustralia":
			return true
		}
	}
	return false
}
