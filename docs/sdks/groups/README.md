# Groups

## Overview

Groups, lists of people, screenings and reviewer decisions.

### Available Operations

* [List](#list) - List an environment's groups.
* [Create](#create) - Create a group to screen.
* [Get](#get) - Read a group.
* [Update](#update) - Change a group.
* [Delete](#delete) - Erase a group.
* [ListPeople](#listpeople) - List a group's people.
* [AddPeople](#addpeople) - Add or update a batch of people.
* [GetPerson](#getperson) - Read a person's report.
* [ReviewPerson](#reviewperson) - Record a decision on a person.
* [JudgeRecord](#judgerecord) - Join one of a person's records to them, or reject it.
* [StartSearch](#startsearch) - Search again on a value of one of a person's records.
* [RunSearch](#runsearch) - Run a counted password or hash search.
* [ListScreenings](#listscreenings) - List a group's screenings.
* [StartScreening](#startscreening) - Start a screening.
* [GetScreening](#getscreening) - Read a screening.
* [CancelScreening](#cancelscreening) - Cancel a screening.

## List

List this environment's groups, newest first. Pass next_before back as before for the next page. Requires the groups scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGroups" method="get" path="/v1/workspaces/{workspace_id}/groups/{environment}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.List(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, noresponse.Pointer[int64](20), nil)
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
| `limit`                                                                                                                                                                | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | At most this many items.                                                                                                                                               |                                                                                                                                                                        |
| `before`                                                                                                                                                               | `*string`                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                     | Read groups older than this one.                                                                                                                                       |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListGroupsResponse](../../models/operations/listgroupsresponse.md), error**

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

Create a group to screen: its name, start, venue, host, the principals it protects and the sensitive topics posts are read against. An environment holds at most 1000 groups. Send an Idempotency-Key to repeat a request safely while the group is retained. A replay returns the same group in its current state; changed input returns 409.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGroup" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"github.com/noresponsehq/go/v13/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Create(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, components.GroupRequest{
        Name: "Annual shareholder meeting",
        StartsAt: "2026-01-01T00:00:00Z",
        EndsAt: noresponse.Pointer("2026-01-02T00:00:00Z"),
        Venue: noresponse.Pointer("Example Convention Centre"),
        Host: noresponse.Pointer("Example Holdings"),
        Principals: []components.Principal{
            components.Principal{
                Name: "Alex Principal",
                Aliases: []string{
                    "Alex P",
                },
            },
        },
        Topics: []string{
            "executive pay",
        },
    }, operations.WithGroupsCreateIdempotencyKey("request-01JA2Z8K6Q4V7W9XBCDEFGHJKM"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                             | Type                                                                                                                                                                                                  | Required                                                                                                                                                                                              | Description                                                                                                                                                                                           | Example                                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                   |                                                                                                                                                                                                       |
| `workspaceID`                                                                                                                                                                                         | `string`                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                    | Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.                                                                    | tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                        |
| `environment`                                                                                                                                                                                         | [components.KeyEnvironment](../../models/components/keyenvironment.md)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                    | The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.                                |                                                                                                                                                                                                       |
| `body`                                                                                                                                                                                                | [components.GroupRequest](../../models/components/grouprequest.md)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                    | N/A                                                                                                                                                                                                   |                                                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                    | The options for this request. Available options are `WithGroupsCreateIdempotencyKey`, `WithServerURL`, `WithTemplatedServerURL`, `WithOperationTimeout`, `WithSetHeaders`, `WithSkipDeserialization`. |                                                                                                                                                                                                       |

### Response

**[*operations.CreateGroupResponse](../../models/operations/creategroupresponse.md), error**

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

Read a group and how many people it holds.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroup" method="get" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Get(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.GetGroupResponse](../../models/operations/getgroupresponse.md), error**

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

Change the fields you name. Send expected_version to apply the change only to that version.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGroup" method="patch" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Update(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.GroupUpdate{
        Name: noresponse.Pointer("Annual shareholder meeting"),
        StartsAt: noresponse.Pointer("2026-01-01T00:00:00Z"),
        EndsAt: noresponse.Pointer("2026-01-02T00:00:00Z"),
        Venue: noresponse.Pointer("Example Convention Centre"),
        Host: noresponse.Pointer("Example Holdings"),
        Principals: []components.Principal{
            components.Principal{
                Name: "Alex Principal",
                Aliases: []string{
                    "Alex P",
                },
            },
        },
        Topics: []string{
            "executive pay",
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `body`                                                                                                                                                                 | [components.GroupUpdate](../../models/components/groupupdate.md)                                                                                                       | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.UpdateGroupResponse](../../models/operations/updategroupresponse.md), error**

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

Erase a group with its people, their decisions and its screenings, and stop any screening still running.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGroup" method="delete" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Delete(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.DeleteGroupResponse](../../models/operations/deletegroupresponse.md), error**

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

## ListPeople

Read a group's roster in upload order, with each person's match state, highest finding band and decision. Filter by band or decision; pass next_after back as after.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPeople" method="get" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"github.com/noresponsehq/go/v13/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ListPeople(ctx, operations.ListPeopleRequest{
        WorkspaceID: "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        Environment: components.KeyEnvironmentTest,
        GroupID: "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListPeopleRequest](../../models/operations/listpeoplerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListPeopleResponse](../../models/operations/listpeopleresponse.md), error**

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

## AddPeople

Add or update 1 to 500 people with an Idempotency-Key. A row with the same external_id, else email, else phone as an existing person updates it. Identities are stored encrypted.

### Example Usage

<!-- UsageSnippet language="go" operationID="addPeople" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.AddPeople(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.PersonBatchRequest{
        People: []components.PersonIdentity{},
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `idempotencyKey`                                                                                                                                                       | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.      | account-acct_7f3a9c                                                                                                                                                    |
| `body`                                                                                                                                                                 | [components.PersonBatchRequest](../../models/components/personbatchrequest.md)                                                                                         | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.AddPeopleResponse](../../models/operations/addpeopleresponse.md), error**

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

## GetPerson

Read a person's report: matched profiles, findings with their evidence, every recorded decision, and the other groups where the same email or phone appears.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerson" method="get" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people/{person_id}" example="records" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.GetPerson(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `personID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.GetPersonResponse](../../models/operations/getpersonresponse.md), error**

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

## ReviewPerson

Record a cleared, watch or escalated decision with a note. The latest decision is the person's; the next screening refreshes everyone on watch or escalated.

### Example Usage

<!-- UsageSnippet language="go" operationID="reviewPerson" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people/{person_id}/review" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ReviewPerson(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.PersonReviewRequest{
        Decision: components.PersonDecisionEscalated,
        Note: noresponse.Pointer("Known to the venue team."),
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `personID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `body`                                                                                                                                                                 | [components.PersonReviewRequest](../../models/components/personreviewrequest.md)                                                                                       | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |

### Response

**[*operations.ReviewPersonResponse](../../models/operations/reviewpersonresponse.md), error**

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

## JudgeRecord

Join one of a person's lookup records to them whatever its likelihood, or reject it so it stays a lead. The verdict sticks while lookups find the same record.

### Example Usage

<!-- UsageSnippet language="go" operationID="judgeRecord" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people/{person_id}/records" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.JudgeRecord(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentTest, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.RecordVerdictRequest{
        RecordID: "rec_3f9a1c22d4e5b6a73f9a1c22d4e5b6a7",
        Verdict: components.VerdictRejected,
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `personID`                                                                                                                                                             | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `body`                                                                                                                                                                 | [components.RecordVerdictRequest](../../models/components/recordverdictrequest.md)                                                                                     | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.JudgeRecordResponse](../../models/operations/judgerecordresponse.md), error**

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

## StartSearch

Search again on the value of one field of one of a person's records, named by the record and field. Nothing searches again by itself. A password or hash search counts its breach rows first and runs only when you run it; read the person to see each search and what it found. Send an Idempotency-Key to return the same search after a lost answer, even after it completes. A different request with the same key returns 409. Use a new key for a deliberate new search.

### Example Usage

<!-- UsageSnippet language="go" operationID="startSearch" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people/{person_id}/searches" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"github.com/noresponsehq/go/v13/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.StartSearch(ctx, operations.StartSearchRequest{
        WorkspaceID: "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        Environment: components.KeyEnvironmentLive,
        GroupID: "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        PersonID: "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        IdempotencyKey: noresponse.Pointer("request-01JA2Z8K6Q4V7W9XBCDEFGHJKM"),
        Body: components.SearchRequest{
            RecordID: "rec_3f9a1c22d4e5b6a73f9a1c22d4e5b6a7",
            Field: "username",
            Table: noresponse.Pointer("FORUM_EXAMPLE_COM_2019"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.StartSearchRequest](../../models/operations/startsearchrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |

### Response

**[*operations.StartSearchResponse](../../models/operations/startsearchresponse.md), error**

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

## RunSearch

Run a counted password or hash search once you have seen how many breach rows hold the value.

### Example Usage

<!-- UsageSnippet language="go" operationID="runSearch" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/people/{person_id}/searches/{search_id}/run" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"github.com/noresponsehq/go/v13/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.RunSearch(ctx, operations.RunSearchRequest{
        WorkspaceID: "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        Environment: components.KeyEnvironmentLive,
        GroupID: "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        PersonID: "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        SearchID: "srh_01JA2Z8K6Q4V7W9XBCDEFGHJKM",
        Body: components.ScreeningCancelRequest{},
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.RunSearchRequest](../../models/operations/runsearchrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.RunSearchResponse](../../models/operations/runsearchresponse.md), error**

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

## ListScreenings

List a group's screenings, newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="listScreenings" method="get" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/screenings" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ListScreenings(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", noresponse.Pointer[int64](20), nil)
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `limit`                                                                                                                                                                | `*int64`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | At most this many items.                                                                                                                                               |                                                                                                                                                                        |
| `before`                                                                                                                                                               | `*string`                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                     | Read screenings older than this one.                                                                                                                                   |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.ListScreeningsResponse](../../models/operations/listscreeningsresponse.md), error**

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

## StartScreening

Start a screening of new people and everyone on watch or escalated, with an Idempotency-Key. It finishes completed, partial or failed minutes to hours later; poll it. The answer carries estimated_spend; a screening whose estimate exceeds what the daily screening cap has left answers 409 screening_cap_exceeded. The test environment screens with the synthetic connector, which makes up deterministic profiles, posts and findings.

### Example Usage

<!-- UsageSnippet language="go" operationID="startScreening" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/screenings" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.StartScreening(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.ScreeningRequest{
        Depth: components.ScreeningDepthStandard,
        Platforms: []string{
            "bluesky",
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `idempotencyKey`                                                                                                                                                       | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.      | account-acct_7f3a9c                                                                                                                                                    |
| `body`                                                                                                                                                                 | [components.ScreeningRequest](../../models/components/screeningrequest.md)                                                                                             | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.StartScreeningResponse](../../models/operations/startscreeningresponse.md), error**

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

## GetScreening

Read a screening's status and progress.

### Example Usage

<!-- UsageSnippet language="go" operationID="getScreening" method="get" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/screenings/{screening_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.GetScreening(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `screeningID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.GetScreeningResponse](../../models/operations/getscreeningresponse.md), error**

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

## CancelScreening

Stop a queued or running screening. People already screened keep their results.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelScreening" method="post" path="/v1/workspaces/{workspace_id}/groups/{environment}/{group_id}/screenings/{screening_id}/cancel" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v13"
	"github.com/noresponsehq/go/v13/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.CancelScreening(ctx, "tnt_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.KeyEnvironmentLive, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.ScreeningCancelRequest{})
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
| `groupID`                                                                                                                                                              | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `screeningID`                                                                                                                                                          | `string`                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    | scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                         |
| `body`                                                                                                                                                                 | [components.ScreeningCancelRequest](../../models/components/screeningcancelrequest.md)                                                                                 | :heavy_check_mark:                                                                                                                                                     | N/A                                                                                                                                                                    |                                                                                                                                                                        |
| `opts`                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                                     | The options for this request.                                                                                                                                          |                                                                                                                                                                        |

### Response

**[*operations.CancelScreeningResponse](../../models/operations/cancelscreeningresponse.md), error**

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