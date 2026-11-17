# EnrichmentMode

off: enrichment starts no search. suggest: enrichment starts no search by itself. automatic: after a workspace person's own investigation, and each time a search on that person finishes, enrichment starts searches by itself.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v15/models/components"
)

value := components.EnrichmentModeOff

// Open enum: custom values can be created with a direct type cast
custom := components.EnrichmentMode("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `EnrichmentModeOff`       | off                       |
| `EnrichmentModeSuggest`   | suggest                   |
| `EnrichmentModeAutomatic` | automatic                 |