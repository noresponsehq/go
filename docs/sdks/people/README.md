# People

## Overview

Workspace people, submitted identity, source evidence and expiry.

### Available Operations

* [List](#list) - List workspace people.
* [Create](#create) - Create or reuse a workspace person.
* [Get](#get) - Read a workspace person.
* [Delete](#delete) - Delete a workspace person.
* [JudgeRecord](#judgerecord) - Join one of a workspace person's records to them, or reject it.
* [StartSearch](#startsearch) - Search again on a value of one of a workspace person's records.
* [RunSearch](#runsearch) - Run a counted password or hash search on a workspace person.

## List

Lists retained people in the key workspace and environment, including people in no group. Each returned person records a person.viewed audit event. Expired people are excluded before background deletion. email and phone narrow the list to people holding that exact address or number, so a client can open the person it already has instead of adding a duplicate.

### Example Usage

<!-- UsageSnippet language="go" operationID="listWorkspacePeople" method="get" path="/v1/people" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.List(ctx, noresponse.Pointer[int64](100), nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `limit`                                                                                      | `*int64`                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `after`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Pass next_cursor from the prior page.                                                        |
| `email`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Return only people whose email is this address, compared without case or surrounding spaces. |
| `phone`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Return only people whose phone is this number, compared without surrounding spaces.          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListWorkspacePeopleResponse](../../models/operations/listworkspacepeopleresponse.md), error**

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

Creates a person without a group. An explicit person_id or external_id reuses the named workspace person. A reference-only request keeps submitted details; identity supplied with that reference replaces those details. A shared email or phone only suggests a match. Suggested matches are returned only with people-read access. The workspace retention setting determines expires_at. An existing person keeps their original expiry. An Idempotency-Key returns the same retained person on retries, without creating another. Changed input with the same key answers 409 idempotency_key_reused. A customer reference that names another person or several people answers 409 person_reference_conflict.

### Example Usage

<!-- UsageSnippet language="go" operationID="createWorkspacePerson" method="post" path="/v1/people" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"github.com/noresponsehq/go/v16/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.Create(ctx, components.WorkspacePersonRequest{
        ExternalID: noresponse.Pointer("case-person-1042"),
        Name: noresponse.Pointer("Jordan Example"),
        Email: noresponse.Pointer("jordan@example.com"),
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
| `body`                                                                                                                                                                                                                                                                                                                                                   | [components.WorkspacePersonRequest](../../models/components/workspacepersonrequest.md)                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                      | {<br/>"external_id": "case-person-1042",<br/>"name": "Jordan Example",<br/>"email": "jordan@example.com"<br/>}                                                                                                                                                                                                                                           |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                         | `*string`                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                       | Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person. | request-01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                                                                                                                                                                       |

### Response

**[*operations.CreateWorkspacePersonResponse](../../models/operations/createworkspacepersonresponse.md), error**

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

Returns submitted details, the current retention expiry and the most recent retained completed or partial investigation evidence, together with retained follow-up search evidence from all of this person's groups. Search evidence remains available after its group is deleted. Migrated results whose original per-person status is unknown may also supply evidence. Earlier results and pending jobs remain available through investigation history; each investigation retains its own collected result. Full source records, posts and media are returned by default; explicit evidence.read denial withholds collected evidence and free-text notes. Each read records person.viewed without raw evidence in the audit log. A missing or expired person answers 404.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspacePerson" method="get" path="/v1/people/{person_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.Get(ctx, "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `personID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetWorkspacePersonResponse](../../models/operations/getworkspacepersonresponse.md), error**

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

Deletes a person and their retained evidence, investigation history and group memberships. Requires the separate delete permission. The audit retains only the person id and deletion metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteWorkspacePerson" method="delete" path="/v1/people/{person_id}" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.Delete(ctx, "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM")
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
| `personID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteWorkspacePersonResponse](../../models/operations/deleteworkspacepersonresponse.md), error**

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

Records a verdict on one of the records the person report holds: joined makes the record the person's whatever its likelihood, and rejected keeps it a lead that leads nowhere. The verdict belongs to the person on their own: the person report and enrichment follow it, and each group keeps its own verdicts. It sticks for as long as the person's evidence holds the same record, and a later verdict on the same record replaces it. After each verdict, enrichment reads the person again under the environment's enrichment setting: it starts no search from a rejected record or from what a search of it found, and a joined record's values lead on. A record the person report does not hold answers 404. Each verdict is audited as person.record_judged. The decisions scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="judgeWorkspacePersonRecord" method="post" path="/v1/people/{person_id}/records" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"github.com/noresponsehq/go/v16/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.JudgeRecord(ctx, "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.RecordVerdictRequest{
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
| `personID`                                                                         | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                     |
| `body`                                                                             | [components.RecordVerdictRequest](../../models/components/recordverdictrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.JudgeWorkspacePersonRecordResponse](../../models/operations/judgeworkspacepersonrecordresponse.md), error**

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

Starts a search on the person alone, outside any group, as the group route does. The person's own enrichment then follows what it finds. Its records join the person report. Starts a new search on the value of one field of one of the person's records: a lookup's, a profile's or one an earlier search found. Name the record and the field; the value itself is not sent. The field's kind must be email, phone, username, person_name, password, hash, ip or domain. A domain search finds the breach rows of every email address at the domain, which must be a registrable domain or a host under one; any other domain answers 400. After a workspace person's own investigation or a group's screening, while the environment's enrichment mode, or an investigation's override, is automatic, enrichment also starts searches by itself, with requested_by enrichment, on the distinctive values of the enrichment kinds (usernames by default; email addresses and phone numbers when named) of records more likely than min_likelihood (0.5 by default) to be the person's: each value at most once per person (once more if that search failed), at most 5 searches from one search's records and 20 per person, at most 6 searches deep, within max_spend and the deadline when set, and within the screening cap; at the screening cap it waits and tries again later. A record a search finds starts from the likelihood of the record it came from, divided by one plus the results the search found on that site, or anywhere for a name search, so a chain grows less likely with each search and stops. A group takes mode suggest while the workspace has not set a mode. In suggest mode enrichment starts none by itself, and the person report lists what it would start in suggested_searches. The search runs the breach search and the account and phone lookups that take the value's kind, under the environment's source controls, and finishes seconds later. A password or hash search counts first: it stops at counted with rows, how many breach rows hold the value, and runs only when a person runs it. Name a table, as a lookup remaining entry names it, to fetch the rows an earlier search left out, up to 100. What the search finds joins the person's lookup records as leads, each match naming the search; the person report lists the searches, each with the value and the record and field it came from, and its records are scored like the rest. Without an Idempotency-Key, starting the same open search returns it. Its estimate counts against the environment's screening_daily_cap, and one the cap cannot hold answers 409 screening_cap_exceeded. A record or field the person does not hold answers 404. Each search is audited as person.search_started. Send an Idempotency-Key to retry safely: the same key and request return the same resource in its current state while it is retained, even after work completes. Reusing the key with different input answers 409 idempotency_key_reused. Use a new key for a deliberate new operation. The investigations scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="startWorkspacePersonSearch" method="post" path="/v1/people/{person_id}/searches" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"github.com/noresponsehq/go/v16/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.StartSearch(ctx, "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.SearchRequest{
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
| `personID`                                                                                                                                                                                                                                                                                                                                               | `string`                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                      | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                                                                                                                                                                           |
| `body`                                                                                                                                                                                                                                                                                                                                                   | [components.SearchRequest](../../models/components/searchrequest.md)                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                       | N/A                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                          |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                         | `*string`                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                       | Opaque request key, 16 to 128 characters. Repeating the key and request returns the same resource in its current state while it is retained. A different request with the same key answers 409. Use a new key for a deliberate new operation. Keys are scoped to the environment and operation; person searches also scope them to the group and person. | request-01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                                                                                                                                                                                                                                                                                       |

### Response

**[*operations.StartWorkspacePersonSearchResponse](../../models/operations/startworkspacepersonsearchresponse.md), error**

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

Runs a counted search that was started on the person alone (POST /v1/people/{person_id}/searches), as the group route does. Runs a password or hash search once a person has seen its count. A search still counting answers 409 group_conflict; one already queued, running or finished answers as it is. The body is an empty object. Its estimate counts against the environment's screening_daily_cap. Each run is audited as person.search_run. The investigations scope grants this operation. Explicit evidence.read restrictions override default evidence access. The groups scope is also accepted for existing integrations. The API key selects the workspace and environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="runWorkspacePersonSearch" method="post" path="/v1/people/{person_id}/searches/{search_id}/run" -->
```go
package main

import(
	"context"
	"os"
	noresponse "github.com/noresponsehq/go/v16"
	"github.com/noresponsehq/go/v16/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := noresponse.New(
        noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
    )

    res, err := s.People.RunSearch(ctx, "per_01JA2Z8K6Q4V7W9XBCDEFGHJKM", "srh_01JA2Z8K6Q4V7W9XBCDEFGHJKM", components.ScreeningCancelRequest{})
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
| `personID`                                                                             | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | per_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `searchID`                                                                             | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    | srh_01JA2Z8K6Q4V7W9XBCDEFGHJKM                                                         |
| `body`                                                                                 | [components.ScreeningCancelRequest](../../models/components/screeningcancelrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.RunWorkspacePersonSearchResponse](../../models/operations/runworkspacepersonsearchresponse.md), error**

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