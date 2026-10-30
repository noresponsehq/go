# SourceID

A NoResponse source: github (GitHub accounts), phone (the phone lookup), breach (breach exposure), rdap (the email domain's registration), mx (the email domain's mail servers), ipintel (IP intelligence) or an email account source that searches one platform for an account by email address (adobe, nvidia, wix, life360, zapier, google). google runs only for Gmail addresses.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
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
| `SourceIDAdobe`   | adobe             |
| `SourceIDNvidia`  | nvidia            |
| `SourceIDWix`     | wix               |
| `SourceIDLife360` | life360           |
| `SourceIDZapier`  | zapier            |
| `SourceIDGoogle`  | google            |