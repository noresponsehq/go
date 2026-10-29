# ErasureReceiptTargetName

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.ErasureReceiptTargetNameServing

// Open enum: custom values can be created with a direct type cast
custom := components.ErasureReceiptTargetName("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `ErasureReceiptTargetNameServing`              | serving                                        |
| `ErasureReceiptTargetNameOutbox`               | outbox                                         |
| `ErasureReceiptTargetNameCache`                | cache                                          |
| `ErasureReceiptTargetNameDownstreamProjection` | downstream_projection                          |