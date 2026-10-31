# Settings

## Overview

Environment policy, history and previews.

### Available Operations

* [Get](#get) - Read an environment's settings.
* [Update](#update) - Change an environment's settings.
* [ListHistory](#listhistory) - Read immutable configuration versions.
* [Restore](#restore) - Restore captured settings as a new guarded version.
* [Preview](#preview) - Preview exact retained-score threshold transitions without side effects.

## Get

Read this environment retention, source selection and effective flow policies. A flow without its own thresholds inherits account_creation. Requires settings or management. Another environment answers404. Read-only and safe to retry.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSettings" method="get" path="/v1/workspaces/{workspace_id}/settings/{environment}" example="default" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Get(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.GetSettingsResponse](../../models/operations/getsettingsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.BadRequestError       | 400                             | application/problem+json        |
| apierrors.AuthenticationError   | 401                             | application/problem+json        |
| apierrors.PermissionDeniedError | 403                             | application/problem+json        |
| apierrors.NotFoundError         | 404                             | application/problem+json        |
| apierrors.ConflictError         | 409                             | application/problem+json        |
| apierrors.RateLimitError        | 429                             | application/problem+json        |
| apierrors.APIStatusError        | 4XX, 5XX                        | \*/\*                           |

## Update

Change named retention, source or flow settings in one transaction. Requires settings or management; changes to active policy also require policy_publish, which legacy management includes. Use expected_policy_version to reject concurrent changes and record a change_reason. Each accepted change creates a fresh version and an atomic audit event. Retention-only edits do not require policy publication authority. Flow thresholds preserve their declared ordering and bounds.

### Example Usage: changed

<!-- UsageSnippet language="go" operationID="updateSettings" method="patch" path="/v1/workspaces/{workspace_id}/settings/{environment}" example="changed" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.SettingsRequest{
        RetentionDays: noresponse.Pointer[int64](30),
        ExpectedPolicyVersion: noresponse.Pointer[int64](1),
        ChangeReason: noresponse.Pointer("Reviewed retention policy"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: flows

<!-- UsageSnippet language="go" operationID="updateSettings" method="patch" path="/v1/workspaces/{workspace_id}/settings/{environment}" example="flows" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"github.com/noresponsehq/go/v10/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.SettingsRequest{
        RetentionDays: noresponse.Pointer[int64](90),
        RefusedSources: []components.SourceID{
            components.SourceIDPhone,
        },
        Flows: &components.SettingsRequestFlows{
            AccountCreation: &components.FlowChange{
                Admission: components.AdmissionHeld.ToPointer(),
                HoldSeconds: noresponse.Pointer[int64](600),
                OnExpiry: components.OnExpiryReview.ToPointer(),
            },
            Login: &components.FlowChange{
                StepUp: noresponse.Pointer[int64](30),
                Review: noresponse.Pointer[int64](60),
                Block: optionalnullable.From[int64](nil),
                DefaultDepth: components.UsageCountDepthFull.ToPointer(),
            },
            Purchase: &components.FlowChange{
                InheritsAccountCreation: noresponse.Pointer(true),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: signup

<!-- UsageSnippet language="go" operationID="updateSettings" method="patch" path="/v1/workspaces/{workspace_id}/settings/{environment}" example="signup" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"github.com/noresponsehq/go/v10/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.SettingsRequest{
        Flows: &components.SettingsRequestFlows{
            AccountCreation: &components.FlowChange{
                Block: optionalnullable.From(noresponse.Pointer[int64](95)),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.SettingsRequest](../../models/components/settingsrequest.md)                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | {<br/>"expected_policy_version": 1,<br/>"change_reason": "Reviewed retention policy",<br/>"retention_days": 30<br/>}                                                   |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.UpdateSettingsResponse](../../models/operations/updatesettingsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.BadRequestError       | 400                             | application/problem+json        |
| apierrors.AuthenticationError   | 401                             | application/problem+json        |
| apierrors.PermissionDeniedError | 403                             | application/problem+json        |
| apierrors.NotFoundError         | 404                             | application/problem+json        |
| apierrors.ConflictError         | 409                             | application/problem+json        |
| apierrors.RateLimitError        | 429                             | application/problem+json        |
| apierrors.APIStatusError        | 4XX, 5XX                        | \*/\*                           |

## ListHistory

Read immutable configuration versions in one environment with settings or management. Follow before_version for older records.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPolicyHistory" method="get" path="/v1/workspaces/{workspace_id}/settings/{environment}/history" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.ListHistory(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, nil, noresponse.Pointer[int64](50))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `beforeVersion`                                                                                                                                                        | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `limit`                                                                                                                                                                | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListPolicyHistoryResponse](../../models/operations/listpolicyhistoryresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.BadRequestError       | 400                             | application/problem+json        |
| apierrors.AuthenticationError   | 401                             | application/problem+json        |
| apierrors.PermissionDeniedError | 403                             | application/problem+json        |
| apierrors.NotFoundError         | 404                             | application/problem+json        |
| apierrors.ConflictError         | 409                             | application/problem+json        |
| apierrors.RateLimitError        | 429                             | application/problem+json        |
| apierrors.APIStatusError        | 4XX, 5XX                        | \*/\*                           |

## Restore

Restore a retained configuration as a fresh version with the expected current version and a change reason. Requires settings plus policy_publish, or legacy management. Restore preserves the historical record and does not undo customer enforcement.

### Example Usage

<!-- UsageSnippet language="go" operationID="restorePolicy" method="post" path="/v1/workspaces/{workspace_id}/settings/{environment}/restore" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Restore(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, components.PolicyRestoreRequest{
        Version: 157178,
        ExpectedPolicyVersion: 260899,
        ChangeReason: "Adjust thresholds after reviewing retained evidence",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.PolicyRestoreRequest](../../models/components/policyrestorerequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.RestorePolicyResponse](../../models/operations/restorepolicyresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.BadRequestError       | 400                             | application/problem+json        |
| apierrors.AuthenticationError   | 401                             | application/problem+json        |
| apierrors.PermissionDeniedError | 403                             | application/problem+json        |
| apierrors.NotFoundError         | 404                             | application/problem+json        |
| apierrors.ConflictError         | 409                             | application/problem+json        |
| apierrors.RateLimitError        | 429                             | application/problem+json        |
| apierrors.APIStatusError        | 4XX, 5XX                        | \*/\*                           |

## Preview

Compare proposed policy with bounded retained first answers, including baseline, coverage and uncertainty. Requires settings or management. This read starts no provider work and publishes no changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="previewPolicy" method="post" path="/v1/workspaces/{workspace_id}/settings/{environment}/preview" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"github.com/noresponsehq/go/v10/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Preview(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.PolicyPreviewRequest{
        ExpectedPolicyVersion: 789985,
        From: "2026-01-01T00:00:00Z",
        To: "2026-01-01T00:00:00Z",
        Flows: components.PolicyPreviewRequestFlows{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.PolicyPreviewRequest](../../models/components/policypreviewrequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.PreviewPolicyResponse](../../models/operations/previewpolicyresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.BadRequestError       | 400                             | application/problem+json        |
| apierrors.AuthenticationError   | 401                             | application/problem+json        |
| apierrors.PermissionDeniedError | 403                             | application/problem+json        |
| apierrors.NotFoundError         | 404                             | application/problem+json        |
| apierrors.ConflictError         | 409                             | application/problem+json        |
| apierrors.RateLimitError        | 429                             | application/problem+json        |
| apierrors.APIStatusError        | 4XX, 5XX                        | \*/\*                           |