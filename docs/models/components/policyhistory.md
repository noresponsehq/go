# PolicyHistory

Immutable settings snapshots, newest first. history_complete is false when earlier versions predate capture. A restored version is published with a fresh version number. History is removed with the tenant.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Items`                                                                        | [][components.PolicyHistoryItem](../../models/components/policyhistoryitem.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `HasMore`                                                                      | `bool`                                                                         | :heavy_check_mark:                                                             | N/A                                                                            |
| `NextBeforeVersion`                                                            | `*int64`                                                                       | :heavy_minus_sign:                                                             | N/A                                                                            |
| `HistoryComplete`                                                              | `bool`                                                                         | :heavy_check_mark:                                                             | N/A                                                                            |