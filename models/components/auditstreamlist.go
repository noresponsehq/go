package components

// AuditStreamList - Read audit stream subscriptions.
type AuditStreamList struct {
	Items []AuditStream `json:"items"`
}

func (a *AuditStreamList) GetItems() []AuditStream {
	if a == nil {
		return []AuditStream{}
	}
	return a.Items
}
