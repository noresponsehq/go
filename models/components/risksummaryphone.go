package components

type RiskSummaryPhone struct {
	LineType *LineType `json:"line_type,omitzero"`
	Country  *string   `json:"country,omitzero"`
	// Highest report count any source gave for the number.
	ReportCount *int64 `json:"report_count,omitzero"`
	// Carrier-network metadata from privately weighted sources; not upstream provider attribution.
	Carrier *string `json:"carrier,omitzero"`
}

func (r *RiskSummaryPhone) GetLineType() *LineType {
	if r == nil {
		return nil
	}
	return r.LineType
}

func (r *RiskSummaryPhone) GetCountry() *string {
	if r == nil {
		return nil
	}
	return r.Country
}

func (r *RiskSummaryPhone) GetReportCount() *int64 {
	if r == nil {
		return nil
	}
	return r.ReportCount
}

func (r *RiskSummaryPhone) GetCarrier() *string {
	if r == nil {
		return nil
	}
	return r.Carrier
}
