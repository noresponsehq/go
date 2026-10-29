# Bucket

The period each count covers, in UTC; a week starts on Monday.

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