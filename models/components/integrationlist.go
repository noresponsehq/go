package components

// IntegrationList - GET /v1/integrations: every integration of the tenant, oldest first, at most 10. No secret is ever listed.
type IntegrationList struct {
	Items []Integration `json:"items"`
}

func (i *IntegrationList) GetItems() []Integration {
	if i == nil {
		return []Integration{}
	}
	return i.Items
}
