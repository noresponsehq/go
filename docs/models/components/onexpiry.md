# OnExpiry

The decision a held assessment takes when hold_seconds pass before its investigation completes, unless the evidence so far gives a stricter one (default review). It is announced as account.risk_changed like any other change, so an expired hold is always announced. A hosted provider hook, which cannot wait, answers a held account by it at once: deny for block, allow otherwise. Kept while the flow is immediate.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.OnExpiryAllow

// Open enum: custom values can be created with a direct type cast
custom := components.OnExpiry("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `OnExpiryAllow`  | allow            |
| `OnExpiryReview` | review           |
| `OnExpiryBlock`  | block            |