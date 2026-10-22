# WebhookEndpoints

## Overview

Your webhook endpoints: register, list, rotate signing keys, send test events and delete.

### Available Operations

* [List](#list) - List the tenant's webhook endpoints.
* [Create](#create) - Register a webhook endpoint.
* [Delete](#delete) - Delete a webhook endpoint.
* [RotateSigningKey](#rotatesigningkey) - Rotate an endpoint's signing key.
* [SendTestEvent](#sendtestevent) - Send a test event to an endpoint.

## List

Lists your webhook endpoints in every state, ordered by endpoint_id, with the ids of their signing keys and without secrets. Requires the webhooks scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="listWebhookEndpoints" method="get" path="/v1/webhooks" example="endpoints" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.List(ctx)
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

**[*operations.ListWebhookEndpointsResponse](../../models/operations/listwebhookendpointsresponse.md), error**

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

Registers a webhook endpoint and returns it with its first signing secret, shown only in this answer: store it where your receiver reads it. The URL is https on port 443 or 8443 with a DNS host whose every address is public; it is checked now and on every delivery. The endpoint starts unverified and receives assessment events once it echoes an endpoint.verification challenge, which sendTestEvent sends. You can hold 5 endpoints. Requires the webhooks scope, and response_mode detailed also the detailed scope. Registration is not idempotent: a retry after a lost answer registers a second endpoint, which you can list and delete.

### Example Usage: created

<!-- UsageSnippet language="go" operationID="createWebhookEndpoint" method="post" path="/v1/webhooks" example="created" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.Create(ctx, components.WebhookEndpointRequest{
        URL: "https://hooks.example.com/noresponse",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: detailed

<!-- UsageSnippet language="go" operationID="createWebhookEndpoint" method="post" path="/v1/webhooks" example="detailed" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.Create(ctx, components.WebhookEndpointRequest{
        URL: "https://hooks.example.com:8443/noresponse/detailed",
        ResponseMode: components.ResponseModeDetailed.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: standard

<!-- UsageSnippet language="go" operationID="createWebhookEndpoint" method="post" path="/v1/webhooks" example="standard" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.Create(ctx, components.WebhookEndpointRequest{
        URL: "https://hooks.example.com/noresponse",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.WebhookEndpointRequest](../../models/components/webhookendpointrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |

### Response

**[*operations.CreateWebhookEndpointResponse](../../models/operations/createwebhookendpointresponse.md), error**

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

Deletes the endpoint and its signing keys. Deliveries still queued for it are dropped. Requires the webhooks scope. An endpoint that is already gone answers 404.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteWebhookEndpoint" method="delete" path="/v1/webhooks/{endpoint_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.Delete(ctx, "whe_01JA2ZD0E1F2G3H4J5K6M7N8P9")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `endpointID`                                                                              | `string`                                                                                  | :heavy_check_mark:                                                                        | The endpoint's whe_ id. An id that is unknown, malformed or another tenant's answers 404. | whe_01JA2ZD0E1F2G3H4J5K6M7N8P9                                                            |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.DeleteWebhookEndpointResponse](../../models/operations/deletewebhookendpointresponse.md), error**

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

## RotateSigningKey

Issues a new signing key for the endpoint and returns its secret once. Every delivery from now on is signed with it; the key it replaces keeps verifying for 7 days, so your receiver can hold both while it switches. Requires the webhooks scope. Rotation is not idempotent: a retry after a lost answer issues another key, which replaces the first.

### Example Usage: empty

<!-- UsageSnippet language="go" operationID="rotateWebhookSigningKey" method="post" path="/v1/webhooks/{endpoint_id}/rotate" example="empty" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.RotateSigningKey(ctx, "whe_01JA2ZD0E1F2G3H4J5K6M7N8P9", components.WebhookKeyRotationRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: rotated

<!-- UsageSnippet language="go" operationID="rotateWebhookSigningKey" method="post" path="/v1/webhooks/{endpoint_id}/rotate" example="rotated" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.RotateSigningKey(ctx, "whe_01JA2ZD0E1F2G3H4J5K6M7N8P9", components.WebhookKeyRotationRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `endpointID`                                                                                 | `string`                                                                                     | :heavy_check_mark:                                                                           | The endpoint's whe_ id. An id that is unknown, malformed or another tenant's answers 404.    | whe_01JA2ZD0E1F2G3H4J5K6M7N8P9                                                               |
| `body`                                                                                       | [components.WebhookKeyRotationRequest](../../models/components/webhookkeyrotationrequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |

### Response

**[*operations.RotateWebhookSigningKeyResponse](../../models/operations/rotatewebhooksigningkeyresponse.md), error**

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

## SendTestEvent

endpoint.verification sends a signed challenge now and answers whether the endpoint echoed it within 5 seconds; an echo verifies the endpoint, including a disabled one. assessment.updated queues one signed event with test: true and a fixed example assessment in the endpoint's response mode, delivered and retried like any other event; the endpoint must be verified. Both use the endpoint's current signing key. Requires the webhooks scope.

### Example Usage: assessment_updated

<!-- UsageSnippet language="go" operationID="sendWebhookTestEvent" method="post" path="/v1/webhooks/{endpoint_id}/test" example="assessment_updated" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.SendTestEvent(ctx, "whe_01JA2ZD0E1F2G3H4J5K6M7N8P9", components.WebhookTestRequest{
        Event: components.WebhookTestRequestEventAssessmentUpdated,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: verification

<!-- UsageSnippet language="go" operationID="sendWebhookTestEvent" method="post" path="/v1/webhooks/{endpoint_id}/test" example="verification" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.WebhookEndpoints.SendTestEvent(ctx, "whe_01JA2ZD0E1F2G3H4J5K6M7N8P9", components.WebhookTestRequest{
        Event: components.WebhookTestRequestEventEndpointVerification,
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

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `endpointID`                                                                              | `string`                                                                                  | :heavy_check_mark:                                                                        | The endpoint's whe_ id. An id that is unknown, malformed or another tenant's answers 404. | whe_01JA2ZD0E1F2G3H4J5K6M7N8P9                                                            |
| `body`                                                                                    | [components.WebhookTestRequest](../../models/components/webhooktestrequest.md)            | :heavy_check_mark:                                                                        | N/A                                                                                       |                                                                                           |

### Response

**[*operations.SendWebhookTestEventResponse](../../models/operations/sendwebhooktesteventresponse.md), error**

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