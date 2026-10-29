# ListEntryList

Newest first. Without a query, returns every active entry for compatibility. With limit or cursor, returns up to 100 entries and an authenticated workspace/environment-bound cursor. Expired entries are omitted.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Items`                                                                    | [][components.ListEntry](../../models/components/listentry.md)             | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `NextCursor`                                                               | `*string`                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        | eyJ0IjoxNzY3MjI1NjAwMDAwMDAwLCJpIjoiYXNtX2V4YW1wbGUiLCJlIjoxNzY3MjI5MjAwfQ |