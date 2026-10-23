# Freshness

Freshness class of the evidence behind a signal. unknown whenever no evidence was obtained.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v4/models/components"
)

value := components.FreshnessFresh

// Open enum: custom values can be created with a direct type cast
custom := components.Freshness("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `FreshnessFresh`   | fresh              |
| `FreshnessStale`   | stale              |
| `FreshnessUnknown` | unknown            |