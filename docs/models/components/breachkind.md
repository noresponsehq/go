# BreachKind

Who is in the table: site_breach, a site's own users; credential_list, a combolist or compilation; stealer_log, malware or botnet victims; scrape; hacking_forum, a criminal hacking, cracking or carding community; abuse_service, the customers of DDoS-for-hire, account generators, dump or card shops, darknet markets, stalkerware or breach-lookup services; darkweb_site, an onion site of unknown purpose.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.BreachKindSiteBreach

// Open enum: custom values can be created with a direct type cast
custom := components.BreachKind("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `BreachKindSiteBreach`     | site_breach                |
| `BreachKindCredentialList` | credential_list            |
| `BreachKindStealerLog`     | stealer_log                |
| `BreachKindScrape`         | scrape                     |
| `BreachKindHackingForum`   | hacking_forum              |
| `BreachKindAbuseService`   | abuse_service              |
| `BreachKindDarkwebSite`    | darkweb_site               |