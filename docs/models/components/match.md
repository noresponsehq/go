# Match

exact compares the original submitted value; canonical uses identifier-specific normalisation; mailbox explicitly selects the email alias rules used by lists.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.MatchExact

// Open enum: custom values can be created with a direct type cast
custom := components.Match("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `MatchExact`     | exact            |
| `MatchCanonical` | canonical        |
| `MatchMailbox`   | mailbox          |