# PostsStopped

Why fewer posts were read than the screening's depth asks for while the account holds more. platform_limit: the platform shows the screening no more of them. page_limit: the pages a screening reads ran out first. rate_limited or unavailable: a later page was held back or failed; the posts before it are kept and the person is partial. Absent when the account was read to its end or to the depth.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v13/models/components"
)

value := components.PostsStoppedPlatformLimit

// Open enum: custom values can be created with a direct type cast
custom := components.PostsStopped("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `PostsStoppedPlatformLimit` | platform_limit              |
| `PostsStoppedPageLimit`     | page_limit                  |
| `PostsStoppedRateLimited`   | rate_limited                |
| `PostsStoppedUnavailable`   | unavailable                 |