# Failure

Why the attempt failed: the URL is no longer valid, the host did not resolve, it resolved to an address that is not public, no answer within 5 seconds, the connection failed, a 3xx (redirects are not followed) or another non-2xx status. Absent for a delivered attempt.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.FailureInvalidURL

// Open enum: custom values can be created with a direct type cast
custom := components.Failure("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `FailureInvalidURL`       | invalid_url               |
| `FailureResolveFailed`    | resolve_failed            |
| `FailureAddressRefused`   | address_refused           |
| `FailureTimeout`          | timeout                   |
| `FailureConnectionFailed` | connection_failed         |
| `FailureRedirect`         | redirect                  |
| `FailureNon2xx`           | non_2xx                   |