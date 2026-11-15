# Verdict

A reviewer's overrule (POST .../people/{person_id}/records), which decides joined whatever the likelihood and sticks for as long as a lookup finds the same record. Absent when nobody overruled it.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v13/models/components"
)

value := components.VerdictJoined

// Open enum: custom values can be created with a direct type cast
custom := components.Verdict("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `VerdictJoined`   | joined            |
| `VerdictRejected` | rejected          |