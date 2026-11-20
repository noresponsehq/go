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
* [RemovePerson](#removeperson) - Remove a person from a group
* [ReviewPerson](#reviewperson) - Record a decision on a person.
* [JudgeRecord](#judgerecord) - Join one of a person's records to them, or reject it.
* [StartSearch](#startsearch) - Search again on a value of one of a person's records.
* [RunSearch](#runsearch) - Run a counted password or hash search.
* [ListScreenings](#listscreenings) - List a group's screenings.
* [StartScreening](#startscreening) - Start a screening.
* [GetScreening](#getscreening) - Read a screening.
* [CancelScreening](#cancelscreening) - Cancel a screening.
* [ExportPeople](#exportpeople) - Export the people in a group.

## List

Groups, newest first. The people scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="listGroups" method="get" path="/v1/groups" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.List(ctx, noresponse.Pointer[int64](20), nil)
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
| `limit`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | At most this many items.                                 |
| `before`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Read groups older than this one.                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

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

Creates a group with its name, start, venue, host, the principals it protects and the sensitive topics its screening reads posts against. An environment holds at most 1000 groups. Send an Idempotency-Key to retry safely: the same key and request return the same resource in its current state while it is retained, even after work completes. Reusing the key with different input answers 409 idempotency_key_reused. Use a new key for a deliberate new operation. The people_write scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createGroup" method="post" path="/v1/groups" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"github.com/noresponsehq/go/v20/optionalnullable"
	"github.com/noresponsehq/go/v20/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Create(ctx, components.GroupRequest{
        Name: "Annual shareholder meeting",
        Type: components.GroupTypeOrganisation.ToPointer(),
        Attribute: noresponse.Pointer("Example Fund"),
        StartsAt: optionalnullable.From(noresponse.Pointer("2026-01-01T00:00:00Z")),
        EndsAt: optionalnullable.From(noresponse.Pointer("2026-01-02T00:00:00Z")),
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

| Parameter                                                                                                                                                                                             | Type                                                                                                                                                                                                  | Required                                                                                                                                                                                              | Description                                                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                   |
| `body`                                                                                                                                                                                                | [components.GroupRequest](../../models/components/grouprequest.md)                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                    | N/A                                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                    | The options for this request. Available options are `WithGroupsCreateIdempotencyKey`, `WithServerURL`, `WithTemplatedServerURL`, `WithOperationTimeout`, `WithSetHeaders`, `WithSkipDeserialization`. |

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

The group and how many people it holds. The people scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroup" method="get" path="/v1/groups/{group_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Get(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

Changes the fields the body names. With expected_version, the change applies only to that version (409 group_conflict otherwise). The people_write scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGroup" method="patch" path="/v1/groups/{group_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/optionalnullable"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Update(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.GroupUpdate{
        Name: noresponse.Pointer("Annual shareholder meeting"),
        Attribute: noresponse.Pointer("Example Fund"),
        StartsAt: optionalnullable.From(noresponse.Pointer("2026-01-01T00:00:00Z")),
        EndsAt: optionalnullable.From(noresponse.Pointer("2026-01-02T00:00:00Z")),
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `groupID`                                                        | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                   |
| `body`                                                           | [components.GroupUpdate](../../models/components/groupupdate.md) | :heavy_check_mark:                                               | N/A                                                              |                                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

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

Deletes the group, its memberships and its review context, and cancels unfinished group work. Workspace people, source evidence and completed investigations remain available. Retained investigations no longer name the deleted group. The delete scope is required; group read, write, investigation and review access do not grant deletion. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGroup" method="delete" path="/v1/groups/{group_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.Delete(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

The roster in upload order, with each person's role, match state and decision. Filter by decision (or undecided) and role (or unstated). The people scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPeople" method="get" path="/v1/groups/{group_id}/people" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ListPeople(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", noresponse.Pointer[int64](100), nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |                                                                                 |
| `groupID`                                                                       | `string`                                                                        | :heavy_check_mark:                                                              | N/A                                                                             | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                  |
| `limit`                                                                         | `*int64`                                                                        | :heavy_minus_sign:                                                              | At most this many people.                                                       |                                                                                 |
| `after`                                                                         | `*string`                                                                       | :heavy_minus_sign:                                                              | Read people after this one.                                                     |                                                                                 |
| `decision`                                                                      | [*components.ListPeopleDecision](../../models/components/listpeopledecision.md) | :heavy_minus_sign:                                                              | Only people with this decision; undecided for none yet.                         |                                                                                 |
| `role`                                                                          | [*components.ListPeopleRole](../../models/components/listpeoplerole.md)         | :heavy_minus_sign:                                                              | Only people with this role; unstated for people whose role was not given.       |                                                                                 |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |                                                                                 |

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

Adds 1 to 500 workspace people to this group, at most 512 KiB. An explicit person_id or external_id reuses the named workspace person. A reference-only row keeps submitted details; an identity supplied with that reference replaces those details. A shared email or phone only suggests matches and does not merge people. Each person's match and coverage in the group come from their latest investigation. The add queues one standard screening of the group's people with no investigation yet, unless another screening of the group is running, no platform is offered or it would pass the daily screening cap; a screening started on the group replaces that queued one; a changed identity sets aside investigations started before the change, and prior investigation history remains available. Identities are encrypted under the workspace key. Idempotency-Key is required: the same key and body return the original import result with Idempotent-Replayed: true while its people remain retained. An expired or deleted imported person answers 404; expired or deleted match suggestions are omitted. A different body answers 409 idempotency_key_reused. A group holds at most 20000 people. The people_write scope grants this operation; groups is also accepted. Suggested matches require people-read access. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="addPeople" method="post" path="/v1/groups/{group_id}/people" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.AddPeople(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.PersonBatchRequest{
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

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       | Example                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |                                                                                                                                                                   |
| `groupID`                                                                                                                                                         | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                    |
| `idempotencyKey`                                                                                                                                                  | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours. | account-acct_7f3a9c                                                                                                                                               |
| `body`                                                                                                                                                            | [components.PersonBatchRequest](../../models/components/personbatchrequest.md)                                                                                    | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |                                                                                                                                                                   |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |                                                                                                                                                                   |

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

Returns submitted details, this group's latest screening evidence and contextual decisions, and the other groups that explicitly include the same workspace person. Membership requires an explicit person or customer reference. The people scope grants this operation; groups is also accepted. Each successful read is durably audited as person.viewed. Full source records, posts and media are returned by default. Explicit evidence.read denial keeps submitted details, coverage, job status and decision summaries, hides collected evidence and free-text notes, and sets evidence_restricted to true. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerson" method="get" path="/v1/groups/{group_id}/people/{person_id}" example="records" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.GetPerson(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `personID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

## RemovePerson

Removes this membership and its group review history. Keeps the workspace person, source evidence and completed investigations. Needs delete access. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeGroupPerson" method="delete" path="/v1/groups/{group_id}/people/{person_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.RemovePerson(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `personID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RemoveGroupPersonResponse](../../models/operations/removegrouppersonresponse.md), error**

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

Marks the person's latest finished investigation cleared, watch or escalated with a note; a person with no finished investigation answers 409 group_conflict. Decisions are kept in order; the latest is the person's until a newer investigation finishes. A decision changes nothing else: no screening reads it. Deprecated: decisions are leaving people and groups. The console no longer records them and no screening reads them. This route is removed in the next contract release; stop calling it. The decisions scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="reviewPerson" method="post" path="/v1/groups/{group_id}/people/{person_id}/review" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ReviewPerson(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.PersonReviewRequest{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `groupID`                                                                        | `string`                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                   |
| `personID`                                                                       | `string`                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                   |
| `body`                                                                           | [components.PersonReviewRequest](../../models/components/personreviewrequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |

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

Records a verdict on one of the records the person's latest lookups found: joined makes the record the person's whatever its likelihood, and rejected keeps it a lead. The verdict sticks for as long as a lookup finds the same record, and a later verdict on the same record replaces it. A record the person's latest lookup does not hold answers 404. Each verdict is audited as person.record_judged. The decisions scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="judgeRecord" method="post" path="/v1/groups/{group_id}/people/{person_id}/records" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.JudgeRecord(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.RecordVerdictRequest{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `groupID`                                                                          | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                     |
| `personID`                                                                         | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                     |
| `body`                                                                             | [components.RecordVerdictRequest](../../models/components/recordverdictrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

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

Starts a new search on the value of one field of one of the person's records: a lookup's, a profile's or one an earlier search found. Name the record and the field; the value itself is not sent. The field's kind must be email, phone, username, person_name, password, hash, ip or domain. A domain search finds the breach rows of every email address at the domain, which must be a registrable domain or a host under one; any other domain answers 400. After a workspace person's own investigation or a group's screening, while the environment's enrichment mode, or an investigation's override, is automatic, enrichment also starts searches by itself, with requested_by enrichment, on the distinctive values of the enrichment kinds (usernames by default; email addresses, phone numbers and person names when named) of records more likely than min_likelihood (0.5 by default) to be the person's: each value at most once per person (once more if that search failed), at most 5 searches from one search's records and 20 per person, at most 6 searches deep, within max_spend and the deadline when set, and within the screening cap; at the screening cap it waits and tries again later. A record a search finds starts from the likelihood of the record it came from, divided by one plus the results the search found on that site, or anywhere for a name search, so a chain grows less likely with each search and stops. A group takes mode suggest while the workspace has not set a mode. In suggest mode enrichment starts none by itself, and the person report lists what it would start in suggested_searches. The search runs the breach search and the account and phone lookups that take the value's kind, under the environment's source controls, and finishes seconds later. A password or hash search counts first: it stops at counted with rows, how many breach rows hold the value, and runs only when a person runs it. Name a table, as a lookup remaining entry names it, to fetch the rows an earlier search left out, up to 100. What the search finds joins the person's lookup records as leads, each match naming the search; the person report lists the searches, each with the value and the record and field it came from, and its records are scored like the rest. Without an Idempotency-Key, starting the same open search returns it. Its estimate counts against the environment's screening_daily_cap, and one the cap cannot hold answers 409 screening_cap_exceeded. A record or field the person does not hold answers 404. Each search is audited as person.search_started. Send an Idempotency-Key to retry safely: the same key and request return the same resource in its current state while it is retained, even after work completes. Reusing the key with different input answers 409 idempotency_key_reused. Use a new key for a deliberate new operation. The investigations scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="startSearch" method="post" path="/v1/groups/{group_id}/people/{person_id}/searches" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.StartSearch(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.SearchRequest{
        RecordID: "rec_3f9a1c22d4e5b6a73f9a1c22d4e5b6a7",
        Field: "username",
        Table: noresponse.Pointer("FORUM_EXAMPLE_COM_2019"),
    }, noresponse.Pointer("request-01JA2Z8K6Q4V7W9XBCDEFGHJKM"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                          |
| `groupID`                                                                                                                                                                                                                                                                                                                                                | `string`                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                                                                                                                                                                           |
| `personID`                                                                                                                                                                                                                                                                                                                                               | `string`                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                      | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                                                                                                                                                                           |
| `body`                                                                                                                                                                                                                                                                                                                                                   | [components.SearchRequest](../../models/components/searchrequest.md)                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                          |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                         | `*string`                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                       | Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person. | request-01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                                                                                                                                                                       |

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

Runs a password or hash search once a person has seen its count. A search still counting answers 409 group_conflict; one already queued, running or finished answers as it is. The body is an empty object. Its estimate counts against the environment's screening_daily_cap. Each run is audited as person.search_run. The investigations scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="runSearch" method="post" path="/v1/groups/{group_id}/people/{person_id}/searches/{search_id}/run" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.RunSearch(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "srh_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.ScreeningCancelRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `groupID`                                                                              | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `personID`                                                                             | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `searchID`                                                                             | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | srh_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `body`                                                                                 | [components.ScreeningCancelRequest](../../models/components/screeningcancelrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

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

Screenings, newest first. The people scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="listScreenings" method="get" path="/v1/groups/{group_id}/screenings" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ListScreenings(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", noresponse.Pointer[int64](20), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `limit`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | At most this many items.                                 |                                                          |
| `before`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Read screenings older than this one.                     |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

Queues a screening that reruns the full search for every person in the group. One screening of a group runs at a time; it finishes completed, partial (some people not covered, counted in people_failed) or failed, and fails with error_code expired after 24 hours. Idempotency-Key is required and binds the body for the screening's life. The 202 body carries estimated_spend, what the selected people would cost at the depth and platforms. A screening whose estimate exceeds the environment's screening_daily_cap, less the day's recorded spend and what open screenings may still spend, answers 409 screening_cap_exceeded and queues nothing. The test environment screens with a deterministic simulator that contacts no source. A platform that no connector serves in the environment answers 409 screening_unavailable. The investigations scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="startScreening" method="post" path="/v1/groups/{group_id}/screenings" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.StartScreening(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "account-acct_7f3a9c", components.ScreeningRequest{
        Depth: components.ScreeningDepthStandard,
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

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       | Example                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |                                                                                                                                                                   |
| `groupID`                                                                                                                                                         | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                    |
| `idempotencyKey`                                                                                                                                                  | `string`                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours. | account-acct_7f3a9c                                                                                                                                               |
| `body`                                                                                                                                                            | [components.ScreeningRequest](../../models/components/screeningrequest.md)                                                                                        | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |                                                                                                                                                                   |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |                                                                                                                                                                   |

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

The screening's status and progress. The people scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="getScreening" method="get" path="/v1/groups/{group_id}/screenings/{screening_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.GetScreening(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `groupID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `screeningID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

Stops a queued or running screening. People it already screened keep their results. Cancelling a finished screening answers it unchanged. The body is an empty object. The investigations scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelScreening" method="post" path="/v1/groups/{group_id}/screenings/{screening_id}/cancel" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"github.com/noresponsehq/go/v20/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.CancelScreening(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.ScreeningCancelRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `groupID`                                                                              | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `screeningID`                                                                          | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | scn_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `body`                                                                                 | [components.ScreeningCancelRequest](../../models/components/screeningcancelrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

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

## ExportPeople

Returns one page for export. Each returned person is durably audited as person.exported with the actor, group, person ID and request ID. Requires both people read access and exports. Explicit evidence restrictions also hide free-text decision notes. No raw evidence enters audit logs. Follow next_after until absent for a complete export. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="exportPeople" method="get" path="/v1/groups/{group_id}/export" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v20"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.Groups.ExportPeople(ctx, "grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM", noresponse.Pointer[int64](100), nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |                                                                                 |
| `groupID`                                                                       | `string`                                                                        | :heavy_check_mark:                                                              | N/A                                                                             | grp_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                  |
| `limit`                                                                         | `*int64`                                                                        | :heavy_minus_sign:                                                              | At most this many people.                                                       |                                                                                 |
| `after`                                                                         | `*string`                                                                       | :heavy_minus_sign:                                                              | Read people after this one.                                                     |                                                                                 |
| `decision`                                                                      | [*components.ListPeopleDecision](../../models/components/listpeopledecision.md) | :heavy_minus_sign:                                                              | Only people with this decision; undecided for none yet.                         |                                                                                 |
| `role`                                                                          | [*components.ListPeopleRole](../../models/components/listpeoplerole.md)         | :heavy_minus_sign:                                                              | Only people with this role; unstated for people whose role was not given.       |                                                                                 |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |                                                                                 |

### Response

**[*operations.ExportPeopleResponse](../../models/operations/exportpeopleresponse.md), error**

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