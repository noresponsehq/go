package components

// RecordSource - Where a record came from.
type RecordSource struct {
	// The id of the connector that read the record.
	Connector string `json:"connector"`
	// The kind of source: platform for an account read off a live platform.
	Kind string `json:"kind"`
	// The host a person knows the source's site by.
	Site *string `json:"site,omitzero"`
	// The table a leaked row came from.
	Table *string `json:"table,omitzero"`
	// When the source's data dates from, at the precision it is known.
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
