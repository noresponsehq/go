package components

import (
	"github.com/noresponsehq/go/v5/internal/utils"
)

// Breach - One breach table that matched, detailed view only. name is the catalogue's display name for the table's site, not the upstream provider. No member carries a record, value, password or hash.
type Breach struct {
	Name string `json:"name"`
	// The breached site's domain, lower case. Absent when the table has none (a credential list, for example).
	Domain *string `json:"domain,omitzero"`
	// The table's category from its name, lower case (social, gaming, combolist, ...).
	Category *string `json:"category,omitzero"`
	// Who is in the table: site_breach, a site's own users; credential_list, a combolist or compilation; stealer_log, malware or botnet victims; scrape; hacking_forum, a criminal hacking, cracking or carding community; abuse_service, the customers of DDoS-for-hire, account generators, dump or card shops, darknet markets, stalkerware or breach-lookup services; darkweb_site, an onion site of unknown purpose.
	Kind BreachKind `json:"kind"`
	// A breach table's date year.
	Year *int64 `json:"year,omitzero"`
	// Absent when only the year is known.
	Month *int64 `json:"month,omitzero"`
	// The submitted identifier a breach table matched on: email or username. Phone and name are not breach selectors.
	MatchedOn BreachSelector  `json:"matched_on"`
	Exposed   []BreachExposed `json:"exposed,omitzero"`
}

func (b Breach) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Breach) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (b *Breach) GetName() string {
	if b == nil {
		return ""
	}
	return b.Name
}

func (b *Breach) GetDomain() *string {
	if b == nil {
		return nil
	}
	return b.Domain
}

func (b *Breach) GetCategory() *string {
	if b == nil {
		return nil
	}
	return b.Category
}

func (b *Breach) GetKind() BreachKind {
	if b == nil {
		return BreachKind("")
	}
	return b.Kind
}

func (b *Breach) GetYear() *int64 {
	if b == nil {
		return nil
	}
	return b.Year
}

func (b *Breach) GetMonth() *int64 {
	if b == nil {
		return nil
	}
	return b.Month
}

func (b *Breach) GetMatchedOn() BreachSelector {
	if b == nil {
		return BreachSelector("")
	}
	return b.MatchedOn
}

func (b *Breach) GetExposed() []BreachExposed {
	if b == nil {
		return nil
	}
	return b.Exposed
}
