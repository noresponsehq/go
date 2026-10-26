# Settings

## Overview

### Available Operations

* [Get](#get) - Read an environment's settings.
* [Update](#update) - Change an environment's settings.

## Get

Reads the settings of the signing key's environment: retention in days; per flow (account_creation, login, purchase) the step_up, review and block thresholds, the default depth and the admission; and the sources the environment refuses. A flow without thresholds of its own shows account_creation's with inherits_account_creation true. Requires the management scope; the other environment answers 404 not_found. Read-only and safe to retry.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSettings" method="get" path="/v1/workspaces/{workspace_id}/settings/{environment}" example="default" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
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

Changes only what the body names, in one transaction: retention_days, refused_sources, and per flow step_up, review, block (null turns blocking off), default_depth, admission, hold_seconds, on_expiry, or inherits_account_creation to return login or purchase to account_creation's thresholds. After the change each flow keeps 0 < step_up < review <= 100 and block off or review <= block <= 100; a change that breaks it answers 400 naming the flow. A policy change is one new policy_version with the signing key's prefix as updated_by, and every change is recorded in the workspace's audit log with the key as its actor. Requires the management scope; the other environment answers 404 not_found. A change reaches assessments within a minute.

### Example Usage: changed

<!-- UsageSnippet language="go" operationID="updateSettings" method="patch" path="/v1/workspaces/{workspace_id}/settings/{environment}" example="changed" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.SettingsRequest{})
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
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
	"github.com/noresponsehq/go/v8/optionalnullable"
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
                DefaultDepth: components.DefaultDepthFull.ToPointer(),
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
	noresponse "github.com/noresponsehq/go/v8"
	"github.com/noresponsehq/go/v8/models/components"
	"github.com/noresponsehq/go/v8/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Settings.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, components.SettingsRequest{
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
| `body`                                                                                                                                                                 | [components.SettingsRequest](../../models/components/settingsrequest.md)                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
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