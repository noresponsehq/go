package components

// BreachKind - Who is in the table: site_breach, a site's own users; credential_list, a combolist or compilation; stealer_log, malware or botnet victims; scrape; hacking_forum, a criminal hacking, cracking or carding community; abuse_service, the customers of DDoS-for-hire, account generators, dump or card shops, darknet markets, stalkerware or breach-lookup services; darkweb_site, an onion site of unknown purpose.
type BreachKind string

const (
	BreachKindSiteBreach     BreachKind = "site_breach"
	BreachKindCredentialList BreachKind = "credential_list"
	BreachKindStealerLog     BreachKind = "stealer_log"
	BreachKindScrape         BreachKind = "scrape"
	BreachKindHackingForum   BreachKind = "hacking_forum"
	BreachKindAbuseService   BreachKind = "abuse_service"
	BreachKindDarkwebSite    BreachKind = "darkweb_site"
)

func (e BreachKind) ToPointer() *BreachKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *BreachKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "site_breach", "credential_list", "stealer_log", "scrape", "hacking_forum", "abuse_service", "darkweb_site":
			return true
		}
	}
	return false
}
