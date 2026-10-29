# ConsoleLookups

The live environment's daily cap on assessments members make through the console (the playground), as spent on the current UTC day. Each console submission with a new Idempotency-Key spends one lookup; a replay spends nothing, and API keys and the test environment spend none. When used reaches limit, the next lookup answers 429 lookup_limit_reached until resets_at. The count is the one the cap enforces.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Used`                                                           | `int64`                                                          | :heavy_check_mark:                                               | Lookups spent since the last reset.                              |                                                                  |
| `Limit`                                                          | `int64`                                                          | :heavy_check_mark:                                               | Lookups allowed per UTC day.                                     |                                                                  |
| `ResetsAt`                                                       | `string`                                                         | :heavy_check_mark:                                               | The next reset: the coming midnight UTC, when used returns to 0. | 2026-01-02T00:00:00.000Z                                         |