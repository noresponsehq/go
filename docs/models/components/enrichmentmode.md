# EnrichmentMode

off: enrichment starts no search and suggests none. suggest: enrichment starts no search by itself, and the person report lists the searches it would start in suggested_searches. automatic: after a workspace person's own investigation or a group's screening, and each time a search on that person finishes, enrichment starts searches by itself. A group takes suggest while the workspace has not set a mode.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v19/models/components"
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