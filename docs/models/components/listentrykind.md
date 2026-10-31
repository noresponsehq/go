# ListEntryKind

email matches one address, whatever the case of its local part. email_domain matches every address of the registrable domain and of its subdomains.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.ListEntryKindEmail

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntryKind("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ListEntryKindEmail`       | email                      |
| `ListEntryKindEmailDomain` | email_domain               |