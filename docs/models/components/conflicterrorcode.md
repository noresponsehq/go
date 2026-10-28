# ConflictErrorCode

Machine codes an HTTP 409 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.ConflictErrorCodeIdempotencyKeyReused

// Open enum: custom values can be created with a direct type cast
custom := components.ConflictErrorCode("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `ConflictErrorCodeIdempotencyKeyReused`        | idempotency_key_reused                         |
| `ConflictErrorCodeKeyRotationInProgress`       | key_rotation_in_progress                       |
| `ConflictErrorCodeWebhookEndpointLimitReached` | webhook_endpoint_limit_reached                 |
| `ConflictErrorCodeWebhookEndpointUnverified`   | webhook_endpoint_unverified                    |
| `ConflictErrorCodeTooManyIntegrations`         | too_many_integrations                          |
| `ConflictErrorCodeRiskVersionAhead`            | risk_version_ahead                             |
| `ConflictErrorCodeKeyLimitReached`             | key_limit_reached                              |
| `ConflictErrorCodeSelfRevocation`              | self_revocation                                |
| `ConflictErrorCodeOutcomeConflict`             | outcome_conflict                               |
| `ConflictErrorCodeCredentialConflict`          | credential_conflict                            |
| `ConflictErrorCodeBindingConflict`             | binding_conflict                               |
| `ConflictErrorCodeListEntryExists`             | list_entry_exists                              |
| `ConflictErrorCodeListLimitReached`            | list_limit_reached                             |
| `ConflictErrorCodeListEntryTooBroad`           | list_entry_too_broad                           |