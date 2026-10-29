# Reviews

## Overview

Collaborative review cases, comments and linked subjects.

### Available Operations

* [List](#list) - List retained account reviews.
* [Create](#create) - Open a review for a retained account.
* [Get](#get) - Read an account review.
* [Update](#update) - Assign, prioritize or reopen a guarded review.
* [Resolve](#resolve) - Resolve a review against an exact account risk version.
* [ListComments](#listcomments) - Read retained review comments.
* [Comment](#comment) - Append a comment against the current review version.
* [ListResolutions](#listresolutions) - Read immutable review resolution history.
* [ListViews](#listviews) - List your private and shared review filters.
* [CreateView](#createview) - Save a tenant-local review filter.
* [UpdateView](#updateview) - Replace your saved filter with optimistic concurrency.
* [DeleteView](#deleteview) - Delete your saved filter with optimistic concurrency.
* [Export](#export) - Export an account review
* [ListSubjects](#listsubjects) - List retained subjects linked to a review.
* [AddSubject](#addsubject) - Link retained tenant-local evidence to an investigation.
* [DeleteSubject](#deletesubject) - Remove a linked subject using the current review version.

## List

List retained investigation reviews using filters and an authenticated cursor.

### Example Usage

<!-- UsageSnippet language="go" operationID="listReviews" method="get" path="/v1/workspaces/{workspace_id}/reviews/{environment}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"github.com/noresponsehq/go/v9/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.List(ctx, operations.ListReviewsRequest{
        WorkspaceID: "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        Environment: components.KeyEnvironmentTest,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListReviewsRequest](../../models/operations/listreviewsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListReviewsResponse](../../models/operations/listreviewsresponse.md), error**

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

Open a review for one account. An account can have one active review in its environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createReview" method="post" path="/v1/workspaces/{workspace_id}/reviews/{environment}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.Create(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, components.ReviewCreateRequest{
        AccountID: "acct_7f3a9c",
        Assignee: noresponse.Pointer("usr_customer_admin"),
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
| `body`                                                                                                                                                                 | [components.ReviewCreateRequest](../../models/components/reviewcreaterequest.md)                                                                                       | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.CreateReviewResponse](../../models/operations/createreviewresponse.md), error**

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

## Get

Read a retained review and its current version.

### Example Usage

<!-- UsageSnippet language="go" operationID="getReview" method="get" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.Get(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>")
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.GetReviewResponse](../../models/operations/getreviewresponse.md), error**

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

Change review status, priority or assignment after checking the expected version.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateReview" method="patch" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"github.com/noresponsehq/go/v9/optionalnullable"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>", components.ReviewUpdateRequest{
        ExpectedVersion: 14577,
        Assignee: optionalnullable.From(noresponse.Pointer("usr_customer_admin")),
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ReviewUpdateRequest](../../models/components/reviewupdaterequest.md)                                                                                       | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.UpdateReviewResponse](../../models/operations/updatereviewresponse.md), error**

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

## Resolve

Resolve a review against the exact current review and account decision versions. A resolution does not report a factual outcome or execute enforcement.

### Example Usage

<!-- UsageSnippet language="go" operationID="resolveReview" method="post" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/resolve" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.Resolve(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "<id>", components.ReviewResolveRequest{
        ExpectedVersion: 902778,
        ExpectedRiskVersion: 385930,
        Conclusion: components.ConclusionBenign,
        Note: "Reviewed the retained account evidence and customer report.",
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ReviewResolveRequest](../../models/components/reviewresolverequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.ResolveReviewResponse](../../models/operations/resolvereviewresponse.md), error**

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

## ListComments

Page through retained review comments.

### Example Usage

<!-- UsageSnippet language="go" operationID="listReviewComments" method="get" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/comments" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.ListComments(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>", noresponse.Pointer[int64](50), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            | Example                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |                                                                                                                                                                        |
| `workspaceID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                     | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `environment`                                                                                                                                                          | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                 | :heavy_check_mark:                                                                                                                                                     | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404. |                                                                                                                                                                        |
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `limit`                                                                                                                                                                | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `cursor`                                                                                                                                                               | `*string`                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListReviewCommentsResponse](../../models/operations/listreviewcommentsresponse.md), error**

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

## Comment

Append a comment after checking the expected review version.

### Example Usage

<!-- UsageSnippet language="go" operationID="commentReview" method="post" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/comments" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.Comment(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>", components.ReviewCommentRequest{
        ExpectedVersion: 181980,
        Body: "Reviewed the retained account evidence and customer report.",
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ReviewCommentRequest](../../models/components/reviewcommentrequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.CommentReviewResponse](../../models/operations/commentreviewresponse.md), error**

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

## ListResolutions

Page through immutable review resolutions.

### Example Usage

<!-- UsageSnippet language="go" operationID="listReviewResolutions" method="get" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/resolutions" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.ListResolutions(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>", noresponse.Pointer[int64](50), nil)
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `limit`                                                                                                                                                                | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `beforeVersion`                                                                                                                                                        | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListReviewResolutionsResponse](../../models/operations/listreviewresolutionsresponse.md), error**

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

## ListViews

List private and shared saved review filters visible to you.

### Example Usage

<!-- UsageSnippet language="go" operationID="listReviewViews" method="get" path="/v1/workspaces/{workspace_id}/review_views/{environment}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.ListViews(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive)
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

**[*operations.ListReviewViewsResponse](../../models/operations/listreviewviewsresponse.md), error**

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

## CreateView

Save a private or shared review filter.

### Example Usage

<!-- UsageSnippet language="go" operationID="createReviewView" method="post" path="/v1/workspaces/{workspace_id}/review_views/{environment}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.CreateView(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, components.ReviewViewRequest{
        Name: "High priority reviews",
        Shared: true,
        Filter: components.Filter{
            Assignee: noresponse.Pointer("usr_customer_admin"),
            AssigneePrincipalID: noresponse.Pointer("mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM"),
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
| `body`                                                                                                                                                                 | [components.ReviewViewRequest](../../models/components/reviewviewrequest.md)                                                                                           | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.CreateReviewViewResponse](../../models/operations/createreviewviewresponse.md), error**

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

## UpdateView

Update your saved review filter after checking its expected version.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateReviewView" method="patch" path="/v1/workspaces/{workspace_id}/review_views/{environment}/{view_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.UpdateView(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>", components.ReviewViewRequest{
        Name: "High priority reviews",
        Shared: false,
        Filter: components.Filter{
            Assignee: noresponse.Pointer("usr_customer_admin"),
            AssigneePrincipalID: noresponse.Pointer("mem_01JA2Z8K6Q4V7W9XBCDEFGHJKM"),
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
| `viewID`                                                                                                                                                               | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ReviewViewRequest](../../models/components/reviewviewrequest.md)                                                                                           | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.UpdateReviewViewResponse](../../models/operations/updatereviewviewresponse.md), error**

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

## DeleteView

Delete your saved review filter after checking its expected version.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteReviewView" method="delete" path="/v1/workspaces/{workspace_id}/review_views/{environment}/{view_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.DeleteView(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>", components.ReviewViewDeleteRequest{
        ExpectedVersion: 165419,
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
| `viewID`                                                                                                                                                               | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ReviewViewDeleteRequest](../../models/components/reviewviewdeleterequest.md)                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.DeleteReviewViewResponse](../../models/operations/deletereviewviewresponse.md), error**

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

## Export

Export retained review details, comments, resolutions and linked subjects. Requires current reviews.read and exports.read; API keys need management together with exports, history or audit_manage. The audit event commits before any bytes are returned. Oversized exports are refused without truncation; use the bounded collection APIs.

### Example Usage

<!-- UsageSnippet language="go" operationID="exportReview" method="get" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/export" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.Export(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "<id>")
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ExportReviewResponse](../../models/operations/exportreviewresponse.md), error**

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

## ListSubjects

List retained subjects linked to this review. A link supplies investigation context and does not prove a relationship or change account risk.

### Example Usage

<!-- UsageSnippet language="go" operationID="listReviewSubjects" method="get" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/subjects" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.ListSubjects(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "<id>")
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListReviewSubjectsResponse](../../models/operations/listreviewsubjectsresponse.md), error**

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

## AddSubject

Link an account, exact assessment revision or issued-key association within this environment. Supply the expected review version. A review permits 100 links. Erasure removes links to erased subjects. This action records context without creating feedback or enforcement.

### Example Usage

<!-- UsageSnippet language="go" operationID="addReviewSubject" method="post" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/subjects" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.AddSubject(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "<id>", components.ReviewSubjectRequest{
        Kind: components.ReviewSubjectKindAccount,
        AccountID: "acct_7f3a9c",
        AssessmentID: noresponse.Pointer("asm_01JA2Z8K6Q4V7W9XBCDEFGHJKM"),
        CredentialID: noresponse.Pointer("cred_01JA2Z8K6Q4V7W9XBCDEFGHJKM"),
        Reason: "Linked during review of the account evidence.",
        ExpectedVersion: 691733,
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `body`                                                                                                                                                                 | [components.ReviewSubjectRequest](../../models/components/reviewsubjectrequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.AddReviewSubjectResponse](../../models/operations/addreviewsubjectresponse.md), error**

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

## DeleteSubject

Remove a subject link with the expected review version. The linked account, assessment or key association remains unchanged.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteReviewSubject" method="delete" path="/v1/workspaces/{workspace_id}/reviews/{environment}/{review_id}/subjects/{subject_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v9"
	"github.com/noresponsehq/go/v9/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Reviews.DeleteSubject(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "<id>", "rvs_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.ReviewViewDeleteRequest{
        ExpectedVersion: 299624,
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
| `reviewID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `subjectID`                                                                                                                                                            | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | rvs_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `body`                                                                                                                                                                 | [components.ReviewViewDeleteRequest](../../models/components/reviewviewdeleterequest.md)                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.DeleteReviewSubjectResponse](../../models/operations/deletereviewsubjectresponse.md), error**

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