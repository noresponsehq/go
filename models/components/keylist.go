package components

// KeyList - GET /v1/workspaces/{workspace_id}/keys: the keys of both of the workspace's environments, newest first, at most 200 of each, revoked and expired keys included. No item carries a secret.
type KeyList struct {
	Items []Key `json:"items"`
}

func (k *KeyList) GetItems() []Key {
	if k == nil {
		return []Key{}
	}
	return k.Items
}
