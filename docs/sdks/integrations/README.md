# Integrations

## Overview

Hosted endpoints that answer your sign-up and sign-in hooks.

### Available Operations

* [Create](#create) - Register a hosted provider integration.
* [List](#list) - List the workspace's integrations.
* [Delete](#delete) - Delete an integration.
* [Rotate](#rotate) - Replace an integration's secret.

## Create

Registers a WorkOS, Auth0 or Supabase integration and returns its hosted endpoint, url, so a sign-up or sign-in is scored with no server of your own. For WorkOS, send the Actions signing secret from the WorkOS dashboard and point the user registration and authentication Actions at url. For Supabase, send the hook secret Supabase generates (v1,whsec_...) and set url as the Before User Created hook's HTTPS endpoint. For Auth0, the answer's secret is the shared secret of the NoResponse Auth0 Action, shown once: store it as an Action secret. hook_points defaults to every hook point the platform has (Supabase has only pre_registration); an event that is not listed is allowed without an assessment. The endpoint waits for the full assessment for at most 2.5 seconds (WorkOS, Supabase) or 4.5 seconds (Auth0), denies only when risk.block is true and allows on any failure. Requires the integrations scope; at most 10 per account. Not idempotent: a retry after a lost answer registers another integration.

### Example Usage: auth0

<!-- UsageSnippet language="go" operationID="createIntegration" method="post" path="/v1/integrations" example="auth0" -->
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

    res, err := s.Integrations.Create(ctx, components.IntegrationRequest{
        Platform: components.IntegrationPlatformAuth0,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: supabase

<!-- UsageSnippet language="go" operationID="createIntegration" method="post" path="/v1/integrations" example="supabase" -->
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

    res, err := s.Integrations.Create(ctx, components.IntegrationRequest{
        Platform: components.IntegrationPlatformSupabase,
        SigningSecret: nil,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: workos

<!-- UsageSnippet language="go" operationID="createIntegration" method="post" path="/v1/integrations" example="workos" -->
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

    res, err := s.Integrations.Create(ctx, components.IntegrationRequest{
        Platform: components.IntegrationPlatformWorkos,
        HookPoints: []components.IntegrationHookPoint{
            components.IntegrationHookPointPreRegistration,
            components.IntegrationHookPointPostLogin,
        },
        SigningSecret: nil,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.IntegrationRequest](../../models/components/integrationrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |

### Response

**[*operations.CreateIntegrationResponse](../../models/operations/createintegrationresponse.md), error**

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

## List

Every integration, oldest first. Listings carry no secrets. Requires the integrations scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="listIntegrations" method="get" path="/v1/integrations" example="list" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Integrations.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListIntegrationsResponse](../../models/operations/listintegrationsresponse.md), error**

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

## Delete

Removes the integration and its secrets; its endpoint answers 404 from then on, so remove the Action from the provider first. Requires the integrations scope. An unknown or already deleted integration answers 404.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteIntegration" method="delete" path="/v1/integrations/{integration_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v10"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Integrations.Delete(ctx, "int_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           | Example                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |                                                                                                                                       |
| `integrationID`                                                                                                                       | `string`                                                                                                                              | :heavy_check_mark:                                                                                                                    | Path segment, matched case-sensitively. A value that does not match the integration_id shape, or another workspace's id, answers 404. | int_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                        |
| `opts`                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                              | :heavy_minus_sign:                                                                                                                    | The options for this request.                                                                                                         |                                                                                                                                       |

### Response

**[*operations.DeleteIntegrationResponse](../../models/operations/deleteintegrationresponse.md), error**

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

## Rotate

For Auth0, send {} and the answer's secret is the new shared secret, shown once. For WorkOS, regenerate the Actions secret in the WorkOS dashboard and send it as signing_secret; for Supabase, generate a new hook secret in the Supabase dashboard and send it as signing_secret. The replaced secret is still accepted for 24 hours (previous_secret_expires_at). Requires the integrations scope. Not idempotent: a retry after a lost answer rotates again.

### Example Usage: auth0

<!-- UsageSnippet language="go" operationID="rotateIntegrationSecret" method="post" path="/v1/integrations/{integration_id}/rotate" example="auth0" -->
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

    res, err := s.Integrations.Rotate(ctx, "int_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.IntegrationRotationRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: workos

<!-- UsageSnippet language="go" operationID="rotateIntegrationSecret" method="post" path="/v1/integrations/{integration_id}/rotate" example="workos" -->
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

    res, err := s.Integrations.Rotate(ctx, "int_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.IntegrationRotationRequest{
        SigningSecret: nil,
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

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           | Example                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |                                                                                                                                       |
| `integrationID`                                                                                                                       | `string`                                                                                                                              | :heavy_check_mark:                                                                                                                    | Path segment, matched case-sensitively. A value that does not match the integration_id shape, or another workspace's id, answers 404. | int_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                        |
| `body`                                                                                                                                | [components.IntegrationRotationRequest](../../models/components/integrationrotationrequest.md)                                        | :heavy_check_mark:                                                                                                                    | N/A                                                                                                                                   |                                                                                                                                       |

### Response

**[*operations.RotateIntegrationSecretResponse](../../models/operations/rotateintegrationsecretresponse.md), error**

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