# Keys

## Overview

Your own API keys.

### Available Operations

* [Rotate](#rotate) - Rotate the API key that signs the request.

## Rotate

Mints a new key with the same scopes as the key that signs the request and returns it once. The presented key keeps working for overlap_seconds (default 24 hours, from 60 seconds to 7 days), then answers 401. A key accepts one rotation: rotating again with a key already in its overlap answers 409 key_rotation_in_progress, so rotate with the new key. Rotation is not idempotent: a retry after a lost answer mints another key.

### Example Usage: default

<!-- UsageSnippet language="go" operationID="rotateApiKey" method="post" path="/v1/keys/rotate" example="default" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v7"
	"github.com/noresponsehq/go/v7/models/components"
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
	noresponse "github.com/noresponsehq/go/v7"
	"github.com/noresponsehq/go/v7/models/components"
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
	noresponse "github.com/noresponsehq/go/v7"
	"github.com/noresponsehq/go/v7/models/components"
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