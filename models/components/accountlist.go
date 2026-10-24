package components

// AccountList - Body of the 200 answer to GET /v1/accounts: the calling workspace's accounts whose risk changed after changed_since, in the environment of the calling key, oldest change first. An account is listed when its risk_version advanced and it has a current risk, and each item is the account exactly as GET /v1/accounts/{account_id} answers it. An account appears once per page, at its latest change; when its risk changes again later, it appears again on a later page with the newer risk_version. next_cursor is always present: pass it back as changed_since to read the changes made after this page.
type AccountList struct {
	Items []Account `json:"items"`
	// Opaque cursor after the last change this page read, bound to the workspace and the environment of the calling key. Pass it back unchanged as changed_since. It is present on every page, the last one included: when nothing changed after changed_since it is the cursor that was sent, so a poller keeps calling with the latest next_cursor. A tampered or foreign cursor answers 400 invalid_cursor. Consumers do not parse it.
	NextCursor string `json:"next_cursor"`
}

func (a *AccountList) GetItems() []Account {
	if a == nil {
		return []Account{}
	}
	return a.Items
}

func (a *AccountList) GetNextCursor() string {
	if a == nil {
		return ""
	}
	return a.NextCursor
}
