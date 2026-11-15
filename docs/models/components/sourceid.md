# SourceID

A NoResponse source: github (GitHub accounts), phone (the phone lookup), breach (breach exposure), rdap (the email domain's registration), mx (the email domain's mail servers), ipintel (IP intelligence) or an email account source that searches one platform for an account by email address (adobe, nvidia, wix, life360, zapier, x_email, snapchat_email, amazon, spotify_email, google, duolingo_email). x_email, snapchat_email, amazon and spotify_email only say whether the email has an account on X, Snapchat, Amazon or Spotify. google runs only for Gmail addresses. scratch, dailymotion, substack, duolingo, dockerhub, wikipedia, internetarchive, telegram, spotify, patreon, wordpress, okru and flickr each search that platform for an account by username; telegram also searches by phone number, through a logged-in Telegram lookup account. whatsapp searches WhatsApp for an account by phone number. youtube searches YouTube for a channel by username.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v12/models/components"
)

value := components.SourceIDGithub

// Open enum: custom values can be created with a direct type cast
custom := components.SourceID("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `SourceIDGithub`          | github                    |
| `SourceIDPhone`           | phone                     |
| `SourceIDBreach`          | breach                    |
| `SourceIDRdap`            | rdap                      |
| `SourceIDMx`              | mx                        |
| `SourceIDIpintel`         | ipintel                   |
| `SourceIDAdobe`           | adobe                     |
| `SourceIDNvidia`          | nvidia                    |
| `SourceIDWix`             | wix                       |
| `SourceIDLife360`         | life360                   |
| `SourceIDZapier`          | zapier                    |
| `SourceIDXEmail`          | x_email                   |
| `SourceIDSnapchatEmail`   | snapchat_email            |
| `SourceIDAmazon`          | amazon                    |
| `SourceIDSpotifyEmail`    | spotify_email             |
| `SourceIDGoogle`          | google                    |
| `SourceIDScratch`         | scratch                   |
| `SourceIDDailymotion`     | dailymotion               |
| `SourceIDSubstack`        | substack                  |
| `SourceIDDuolingo`        | duolingo                  |
| `SourceIDDuolingoEmail`   | duolingo_email            |
| `SourceIDDockerhub`       | dockerhub                 |
| `SourceIDWikipedia`       | wikipedia                 |
| `SourceIDInternetarchive` | internetarchive           |
| `SourceIDTelegram`        | telegram                  |
| `SourceIDYoutube`         | youtube                   |
| `SourceIDSpotify`         | spotify                   |
| `SourceIDPatreon`         | patreon                   |
| `SourceIDWordpress`       | wordpress                 |
| `SourceIDOkru`            | okru                      |
| `SourceIDFlickr`          | flickr                    |
| `SourceIDWhatsapp`        | whatsapp                  |