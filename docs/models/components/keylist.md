# KeyList

GET /v1/workspaces/{workspace_id}/keys: the keys of both of the workspace's environments, newest first, at most 200 of each, revoked and expired keys included. No item carries a secret.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Items`                                            | [][components.Key](../../models/components/key.md) | :heavy_check_mark:                                 | N/A                                                |