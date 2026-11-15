package components

// RecordSource - Where a record came from.
type RecordSource struct {
	// The id of the connector that read the record: breach for a breach row.
	Connector string `json:"connector"`
	// The kind of source: platform for an account read off a live platform; for a breach row, the kind of leak its table is: site_breach, credential_list, stealer_log, scrape, hacking_forum, abuse_service or darkweb_site; domain_registry for an email domain's registry record and domain_dns for its mail records.
	Kind string `json:"kind"`
	// The host a person knows the source's site by.
	Site *string `json:"site,omitzero"`
	// The table a leaked row came from.
	Table *string `json:"table,omitzero"`
	// When the source's data dates from, at the precision it is known: YYYY or YYYY-MM for a breach table, and YYYY-MM-DD, the day the computer was infected, for a row of a stealer log.
	Date *string `json:"date,omitzero"`
}

func (r *RecordSource) GetConnector() string {
	if r == nil {
		return ""
	}
	return r.Connector
}

func (r *RecordSource) GetKind() string {
	if r == nil {
		return ""
	}
	return r.Kind
}

func (r *RecordSource) GetSite() *string {
	if r == nil {
		return nil
	}
	return r.Site
}

func (r *RecordSource) GetTable() *string {
	if r == nil {
		return nil
	}
	return r.Table
}

func (r *RecordSource) GetDate() *string {
	if r == nil {
		return nil
	}
	return r.Date
}
