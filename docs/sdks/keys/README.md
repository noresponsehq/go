# Keys

## Overview

Your own API keys.

### Available Operations

* [Rotate](#rotate) - Rotate the API key that signs the request.
* [Create](#create) - Create an API key.
* [List](#list) - List API keys.
* [Revoke](#revoke) - Revoke an API key.

## Rotate

Mints a new key with the same scopes as the key that signs the request and returns it once. The presented key keeps working for overlap_seconds (default 24 hours, from 60 seconds to 7 days), then answers 401. A key accepts one rotation: rotating again with a key already in its overlap answers 409 key_rotation_in_progress, so rotate with the new key. Rotation is not idempotent: a retry after a lost answer mints another key.

### Example Usage: default

<!-- UsageSnippet language="go" operationID="rotateApiKey" method="post" path="/v1/keys/rotate" example="default" -->
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

    res, err := s.Keys.Rotate(ctx, components.KeyRotationRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: one_hour

<!-- UsageSnippet language="go" operationID="rotateApiKey" method="post" path="/v1/keys/rotate" example="one_hour" -->
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

    res, err := s.Keys.Rotate(ctx, components.KeyRotationRequest{
        OverlapSeconds: noresponse.Pointer[int64](3600),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: rotated

<!-- UsageSnippet language="go" operationID="rotateApiKey" method="post" path="/v1/keys/rotate" example="rotated" -->
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

    res, err := s.Keys.Rotate(ctx, components.KeyRotationRequest{})
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
| `request`                                                                      | [components.KeyRotationRequest](../../models/components/keyrotationrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |

### Response

**[*operations.RotateAPIKeyResponse](../../models/operations/rotateapikeyresponse.md), error**

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

## Create

Create an API key in the signing key environment with authority no greater than the issuer holds. Requires management. Optional name, purpose, kind and expiry describe its lifecycle. Send request_id for safe issuance retries: the first successful response returns the secret once; a repeat returns existing metadata without the secret, and changed parameters conflict. Without request_id a retry can create another key. A service credential is independent of its issuer; personal credentials require an authenticated member and retire when that member loses authority.

### Example Usage: live

<!-- UsageSnippet language="go" operationID="createKey" method="post" path="/v1/workspaces/{workspace_id}/keys" example="live" -->
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

    res, err := s.Keys.Create(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyRequest{
        Environment: components.KeyEnvironmentLive,
        Scopes: []components.KeyRequestScope{
            components.KeyRequestScopeAssess,
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
### Example Usage: test

<!-- UsageSnippet language="go" operationID="createKey" method="post" path="/v1/workspaces/{workspace_id}/keys" example="test" -->
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

    res, err := s.Keys.Create(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyRequest{
        Environment: components.KeyEnvironmentTest,
        Scopes: []components.KeyRequestScope{
            components.KeyRequestScopeAssess,
            components.KeyRequestScopeHistory,
            components.KeyRequestScopeWebhooks,
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `workspaceID`                                                                                                                      | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404. | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `body`                                                                                                                             | [components.KeyRequest](../../models/components/keyrequest.md)                                                                     | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |

### Response

**[*operations.CreateKeyResponse](../../models/operations/createkeyresponse.md), error**

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

Lists the keys of the signing key's environment, newest first, at most 200, revoked and expired keys included, without secrets: each key's id, prefix, scopes, status, and who created and revoked it. Requires the management scope. Read-only and safe to retry.

### Example Usage

<!-- UsageSnippet language="go" operationID="listKeys" method="get" path="/v1/workspaces/{workspace_id}/keys" example="list" -->
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

    res, err := s.Keys.List(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `workspaceID`                                                                                                                      | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404. | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.ListKeysResponse](../../models/operations/listkeysresponse.md), error**

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

## Revoke

Revokes a key of the signing key's environment at once: its next request answers 401. Requires the management scope. A key of the other environment or of another workspace answers 404 not_found. A key cannot revoke itself (409 self_revocation): rotate it with keys.rotate, or revoke it with another management key. Revoking a revoked key changes nothing and answers the key as it was revoked, so a retry is safe.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokeKey" method="delete" path="/v1/workspaces/{workspace_id}/keys/{key_id}" example="key" -->
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

    res, err := s.Keys.Revoke(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "01JA2ZE4K5M6N7P8Q9R0S1T2V3")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `workspaceID`                                                                                                                      | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404. | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                     |
| `keyID`                                                                                                                            | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | The key's id. A key of another workspace answers 404, as does a key of the other environment to an API key.                        | 01JA2ZE4K5M6N7P8Q9R0S1T2V3                                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.RevokeKeyResponse](../../models/operations/revokekeyresponse.md), error**

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