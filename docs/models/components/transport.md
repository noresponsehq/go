# Transport

Where the context was derived: backend (also when the request named no transport), browser, or provider_hook for a hosted provider hook.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.TransportBackend

// Open enum: custom values can be created with a direct type cast
custom := components.Transport("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `TransportBackend`      | backend                 |
| `TransportBrowser`      | browser                 |
| `TransportProviderHook` | provider_hook           |