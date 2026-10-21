# noresponse

Developer-friendly & type-safe Go SDK specifically catered to leverage *noresponse* API.


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

NoResponse API: Server-side signup and gateway risk assessment: a recommendation, polling and signed webhooks.

Assess a signup or gateway event, poll a pending assessment, page through an account's history, record what you later learned, and rotate your API key. Every response carries X-Request-ID; every error is an RFC 9457 problem document. Keep the API key on your server.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [noresponse](#noresponse)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/noresponsehq/go/v2
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type | Scheme      | Environment Variable |
| -------- | ---- | ----------- | -------------------- |
| `APIKey` | http | HTTP Bearer | `NORESPONSE_API_KEY` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Assessments](docs/sdks/assessments/README.md)

* [Create](docs/sdks/assessments/README.md#create) - Assess identifiers for a signup or gateway decision.
* [Get](docs/sdks/assessments/README.md#get) - Poll the authoritative assessment of one assessment.
* [RecordOutcome](docs/sdks/assessments/README.md#recordoutcome) - Report what was later learned about an assessed account.
* [Query](docs/sdks/assessments/README.md#query) - Page through one account's unexpired assessments.

### [Keys](docs/sdks/keys/README.md)

* [Rotate](docs/sdks/keys/README.md#rotate) - Rotate the API key that signs the request.

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Query(ctx, components.AssessmentQuery{
		AccountID: "acct_7f3a9c",
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"github.com/noresponsehq/go/v2/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"github.com/noresponsehq/go/v2/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIStatusError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type                      | Status Code | Content Type             |
| ------------------------------- | ----------- | ------------------------ |
| apierrors.BadRequestError       | 400         | application/problem+json |
| apierrors.AuthenticationError   | 401         | application/problem+json |
| apierrors.PermissionDeniedError | 403         | application/problem+json |
| apierrors.NotFoundError         | 404         | application/problem+json |
| apierrors.ConflictError         | 409         | application/problem+json |
| apierrors.RateLimitError        | 429         | application/problem+json |
| apierrors.APIStatusError        | 4XX, 5XX    | \*/\*                    |

### Example

```go
package main

import (
	"context"
	"errors"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/apierrors"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
		},
	})
	if err != nil {

		var e *apierrors.BadRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.AuthenticationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.PermissionDeniedError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.NotFoundError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ConflictError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.RateLimitError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIStatusError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := noresponse.New(
		noresponse.WithServerURL("https://api.noresponse.org"),
		noresponse.WithSecurity(os.Getenv("NORESPONSE_API_KEY")),
	)

	res, err := s.Assessments.Create(ctx, "<value>", components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/noresponsehq/go/v2"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = noresponse.New(noresponse.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future NoResponse SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 
