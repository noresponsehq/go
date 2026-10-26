# EnforcementAction

What the workspace did in its identity provider: revoke_sessions (signed the account out everywhere), restrict (limited what the account can do), unrestrict (lifted an earlier restriction), fail_onboarding (stopped an account that was still onboarding), reject_activation (refused to activate a provisional account), activate (let a held onboarding or provisional account in) or none (decided to take no action).

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.EnforcementActionRevokeSessions

// Open enum: custom values can be created with a direct type cast
custom := components.EnforcementAction("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `EnforcementActionRevokeSessions`   | revoke_sessions                     |
| `EnforcementActionRestrict`         | restrict                            |
| `EnforcementActionUnrestrict`       | unrestrict                          |
| `EnforcementActionFailOnboarding`   | fail_onboarding                     |
| `EnforcementActionRejectActivation` | reject_activation                   |
| `EnforcementActionActivate`         | activate                            |
| `EnforcementActionNone`             | none                                |