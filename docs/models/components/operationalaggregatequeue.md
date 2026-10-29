# OperationalAggregateQueue

Current retained delivery queue, independent of the requested attempt window and status filter. Endpoint filter applies. Does not prove customer enforcement execution.


## Fields

| Field                | Type                 | Required             | Description          | Example              |
| -------------------- | -------------------- | -------------------- | -------------------- | -------------------- |
| `Pending`            | `int64`              | :heavy_check_mark:   | N/A                  |                      |
| `Leased`             | `int64`              | :heavy_check_mark:   | N/A                  |                      |
| `DeadRetained`       | `int64`              | :heavy_check_mark:   | N/A                  |                      |
| `OldestPendingAt`    | `*string`            | :heavy_minus_sign:   | N/A                  | 2026-01-01T00:00:00Z |