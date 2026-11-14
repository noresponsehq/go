# RecordSource

Where a record came from.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Connector`                                                           | `string`                                                              | :heavy_check_mark:                                                    | The id of the connector that read the record.                         | bluesky                                                               |
| `Kind`                                                                | `string`                                                              | :heavy_check_mark:                                                    | The kind of source: platform for an account read off a live platform. | platform                                                              |
| `Site`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | The host a person knows the source's site by.                         | bsky.app                                                              |
| `Table`                                                               | `*string`                                                             | :heavy_minus_sign:                                                    | The table a leaked row came from.                                     | users                                                                 |
| `Date`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | When the source's data dates from, at the precision it is known.      | 2019-03                                                               |