# ScreeningCoverageSources

How many sources this search asked and how many answered, across the lookup stage (sources asked whether an account exists for an identifier) and the platform stage (profiles read on each platform). A source that answered may have found nothing; one that did not answer failed or ran out of time.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Asked`                                                          | `int64`                                                          | :heavy_check_mark:                                               | Lookup sources plus platforms asked.                             |
| `Answered`                                                       | `int64`                                                          | :heavy_check_mark:                                               | Lookup sources plus platforms that answered.                     |
| `Lookups`                                                        | [components.SourceCount](../../models/components/sourcecount.md) | :heavy_check_mark:                                               | N/A                                                              |
| `Platforms`                                                      | [components.SourceCount](../../models/components/sourcecount.md) | :heavy_check_mark:                                               | N/A                                                              |