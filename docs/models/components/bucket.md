# Bucket

The period each count covers: an hour, or a day or a week (from Monday) in time_zone. Hour buckets are whole UTC hours in any time zone.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.BucketHour

// Open enum: custom values can be created with a direct type cast
custom := components.Bucket("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `BucketHour` | hour         |
| `BucketDay`  | day          |
| `BucketWeek` | week         |