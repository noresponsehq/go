# FootprintTimeline

One identifier's dated footprint across every source: the first and last year any event was seen, the years between them, the latest year a source dated activity (omitted when none did), how many distinct sites and which kinds of evidence saw it.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `FirstSeenYear`                                                            | `int64`                                                                    | :heavy_check_mark:                                                         | N/A                                                                        |
| `LastSeenYear`                                                             | `int64`                                                                    | :heavy_check_mark:                                                         | N/A                                                                        |
| `SpanYears`                                                                | `int64`                                                                    | :heavy_check_mark:                                                         | N/A                                                                        |
| `LastActiveYear`                                                           | `*int64`                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Sites`                                                                    | `int64`                                                                    | :heavy_check_mark:                                                         | N/A                                                                        |
| `Families`                                                                 | [][components.FootprintFamily](../../models/components/footprintfamily.md) | :heavy_check_mark:                                                         | N/A                                                                        |