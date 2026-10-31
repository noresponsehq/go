# KeyStatus

active: the key authenticates. expired: a rotation's overlap window has ended (expires_at). revoked: the key was revoked (revoked_at) and answers 401.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.KeyStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.KeyStatus("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `KeyStatusActive`  | active             |
| `KeyStatusExpired` | expired            |
| `KeyStatusRevoked` | revoked            |