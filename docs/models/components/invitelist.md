# InviteList

GET /v1/workspaces/{workspace_id}/invites: the workspace's pending invites, oldest first, at most 50, with no token.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Items`                                                  | [][components.Invite](../../models/components/invite.md) | :heavy_check_mark:                                       | N/A                                                      |