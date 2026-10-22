# Transport

Where this context was derived. Backend is the trusted integration path; browser values are treated as weak hints. provider_hook is a hosted provider hook (an Auth0 or WorkOS Action) calling at hook_point; it needs hook_point and verified false, and its authentication facts are not scored.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.TransportBackend
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `TransportBackend`      | backend                 |
| `TransportBrowser`      | browser                 |
| `TransportProviderHook` | provider_hook           |