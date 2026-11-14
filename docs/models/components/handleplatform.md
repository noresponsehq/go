# HandlePlatform

A platform a screening can cover. synthetic is the deterministic connector's own platform, for exercising the product end to end with no network.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.HandlePlatformInstagram

// Open enum: custom values can be created with a direct type cast
custom := components.HandlePlatform("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `HandlePlatformInstagram` | instagram                 |
| `HandlePlatformX`         | x                         |
| `HandlePlatformBluesky`   | bluesky                   |
| `HandlePlatformYoutube`   | youtube                   |
| `HandlePlatformTiktok`    | tiktok                    |
| `HandlePlatformFacebook`  | facebook                  |
| `HandlePlatformThreads`   | threads                   |
| `HandlePlatformReddit`    | reddit                    |
| `HandlePlatformTelegram`  | telegram                  |
| `HandlePlatformLinkedin`  | linkedin                  |
| `HandlePlatformSynthetic` | synthetic                 |