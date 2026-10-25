# SourceID

A NoResponse source: github (GitHub accounts), phone (the phone lookup), breach (breach exposure), rdap (the email domain's registration), mx (the email domain's mail servers) or ipintel (IP intelligence).

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.SourceIDGithub

// Open enum: custom values can be created with a direct type cast
custom := components.SourceID("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `SourceIDGithub`  | github            |
| `SourceIDPhone`   | phone             |
| `SourceIDBreach`  | breach            |
| `SourceIDRdap`    | rdap              |
| `SourceIDMx`      | mx                |
| `SourceIDIpintel` | ipintel           |