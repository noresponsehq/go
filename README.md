# NoResponse Go SDK

The Go client for the [NoResponse API](https://docs.noresponse.org): assess a signup or gateway event, wait for a pending result, page through an account's history, record what you later learned and verify webhooks.

Call it from your server only. The API key must never reach a browser or an app.

## Install

```sh
go get github.com/noresponsehq/go/v2
```

Go 1.24 or later.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	noresponse "github.com/noresponsehq/go/v2"
	"github.com/noresponsehq/go/v2/models/components"
)

func main() {
	client := noresponse.NewClient(os.Getenv("NORESPONSE_API_KEY"))

	res, err := client.Assessments.Create(context.Background(), noresponse.NewIdempotencyKey(), components.AssessmentRequest{
		AccountID: "acct_7f3a9c",
		Identifiers: components.AssessmentRequestIdentifiers{
			Email: noresponse.Pointer("new.user@example.com"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	assessment := res.Result
	if assessment.Status == components.StatusPending {
		// No recommendation yet. Apply your own fallback, then wait for the
		// webhook or call noresponse.WaitForAssessment.
		return
	}
	fmt.Println(*assessment.Recommendation, assessment.Risk.Score)
}
```

`NewClient` reads nothing from the environment itself; pass the key. `noresponse.New()` with no key option reads `NORESPONSE_API_KEY`.

The recommendation is `allow`, `step_up` or `review`. It is advice: your application decides what to do. A `pending` assessment has no recommendation, and a 429, a 503 or a timeout also needs your own fallback.

Every method takes a `context.Context` first. Cancelling it cancels the request and any retry wait.

## Methods

| Call | Does |
|---|---|
| `client.Assessments.Create(ctx, idempotencyKey, request)` | Assess identifiers. Returns 200 (complete or partial) or 202 (pending) as `res.Result`. |
| `client.Assessments.Get(ctx, assessmentID, nil)` | Fetch the latest version of one assessment. |
| `client.Assessments.RecordOutcome(ctx, assessmentID, idempotencyKey, request)` | Report what you later learned, for example `components.OutcomeConfirmedAbuse`. |
| `client.Assessments.Query(ctx, query)` | One page of an account's history. `res.Next()` fetches the next page. |
| `noresponse.AllAssessments(ctx, client, query)` | Range over every page of an account's history. |
| `client.Accounts.Get(ctx, accountID)` | The account's one current risk: score, a decision for each flow and the assessment it comes from (`history` scope). An account with no decided assessment is a 404. |
| `client.Accounts.ListAssessments(ctx, accountID, limit, cursor, responseMode)` | One page of the account's history with the account id in the URL; pass `nil` for the defaults. `res.Next()` fetches the next page. |
| `noresponse.WaitForAssessment(ctx, client, assessmentID, nil)` | Poll until the assessment is `complete`, honouring `Retry-After`, until `ctx`'s deadline (30 s when it has none). Then a `partial` answer is returned as it stands, and a `pending` one returns `*WaitTimeoutError`. |
| `noresponse.ConstructEvent(payload, headers, secrets)` | Verify a webhook delivery and return the typed event. |
| `client.Keys.Rotate(ctx, request)` | Rotate the API key that signs the request. |
| `client.Integrations.Create(ctx, request)` | Register a hosted WorkOS or Auth0 integration (`integrations` scope). For Auth0, `res.Result.Secret` is shown once. Never retried. |
| `client.Integrations.List(ctx)` | Every integration, oldest first, with no secrets. |
| `client.Integrations.Delete(ctx, integrationID)` | Delete an integration. |
| `client.Integrations.Rotate(ctx, integrationID, body)` | Replace an integration's secret. Never retried. |

Each response carries `Headers` (including `X-Request-ID`) and `Result`.

## Idempotency keys

`Create` and `RecordOutcome` require an idempotency key. The same key with the same body replays the stored answer and starts no new work. The same key with a different body is a 409. Use `noresponse.NewIdempotencyKey()` and store the key beside the signup so a retry after a crash sends the same one.

## Errors

Every failure is an `error` you can inspect with `errors.As`. Problem documents from the API are typed and carry the request id:

```go
_, err := client.Assessments.Get(ctx, id, nil)

var notFound *apierrors.NotFoundError
var limited *apierrors.RateLimitError
var other *apierrors.APIStatusError
switch {
case errors.As(err, &notFound):
	log.Println("no such assessment, request", notFound.RequestID)
case errors.As(err, &limited):
	log.Println("rate limited, request", limited.RequestID)
case errors.As(err, &other):
	log.Println("HTTP", other.StatusCode)
case err != nil:
	log.Fatal(err) // network failure or cancelled context
}
```

The typed errors are `BadRequestError` (400), `AuthenticationError` (401), `PermissionDeniedError` (403), `NotFoundError` (404), `ConflictError` (409) and `RateLimitError` (429), all in `github.com/noresponsehq/go/v2/models/apierrors`. Any other status is an `*apierrors.APIStatusError`.

## Retries and timeouts

The client retries 408, 429, 5xx and connection errors with exponential backoff and jitter: 500 ms first, 8 s at most, 30 s in total. The same idempotency key is sent each time, so a retry never repeats work. Change it for every call or for one:

```go
client := noresponse.NewClient(key,
	noresponse.WithTimeout(10*time.Second),
	noresponse.WithRetryConfig(retry.Config{
		Strategy:              "backoff",
		Backoff:               &retry.BackoffStrategy{InitialInterval: 200, MaxInterval: 2000, Exponent: 2, MaxElapsedTime: 5000}, // milliseconds
		RetryConnectionErrors: true,
	}),
)

res, err := client.Assessments.Get(ctx, id, nil, operations.WithRetries(retry.Config{Strategy: "none"}))
```

## Paging

```go
query := components.AssessmentQuery{AccountID: "acct_7f3a9c"}
for assessment, err := range noresponse.AllAssessments(ctx, client, query) {
	if err != nil {
		return err
	}
	fmt.Println(assessment.AssessmentID, assessment.Status)
}
```

## Issued keys

When your signup gate issues an API key to the new account, send only its hash, as `issued_credential.sha256`; the key itself never leaves your backend.

```go
sha256, err := noresponse.HashIssuedCredential(issuedKey) // 64 lower-case hex characters
```

## Webhooks

Register your endpoint with a key that has the `webhooks` scope. The signing secret is shown only in this answer, so store it, then send the verification challenge:

```go
created, err := client.WebhookEndpoints.Create(ctx, components.WebhookEndpointRequest{URL: "https://app.example.com/webhooks/noresponse"})
if err != nil {
	return err
}
secret := created.Result.SigningSecret // store secret.KeyID and secret.Secret
check, err := client.WebhookEndpoints.SendTestEvent(ctx, created.Result.Endpoint.EndpointID,
	components.WebhookTestRequest{Event: components.WebhookTestRequestEventEndpointVerification}) // check.Result.Verified
```

`List`, `RotateSigningKey` (the old key verifies for 7 more days) and `Delete` manage it afterwards. An `assessment.updated` test event arrives with `event.Test` true.

Pass the raw request body. Re-encoding parsed JSON breaks the signature.

```go
secrets := map[string]string{"whk_example0001": os.Getenv("NORESPONSE_WEBHOOK_SECRET")}

http.HandleFunc("/webhooks/noresponse", func(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 1<<20))
	if err != nil {
		http.Error(w, "unreadable body", http.StatusBadRequest)
		return
	}
	event, err := noresponse.ConstructEvent(payload, r.Header, secrets)
	if err != nil {
		http.Error(w, "invalid signature", http.StatusBadRequest)
		return
	}
	switch data := event.Data.(type) {
	case *components.Assessment:
		if event.Test {
			break // SendTestEvent's fixed example assessment: apply nothing
		}
		// The full assessment, as Get returns it.
	case *noresponse.AssessmentExpired:
		// Delete what you hold for data.AssessmentID.
	case *noresponse.EndpointVerification:
		body, err := noresponse.VerificationResponse(event)
		if err != nil {
			http.Error(w, "invalid verification", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(body)
		return
	}
	w.WriteHeader(http.StatusNoContent)
})
```

Deliveries arrive at least once and unordered. Deduplicate on `event.ID` and keep the highest `AssessmentVersion` per assessment.

### Late results

An assessment can rise after the signup was allowed. `HandleLateResult` verifies the `assessment.updated` delivery with `ConstructEvent`, applies a policy (default: revoke sessions on `review`) and acts through your auth provider:

```go
provider := &noresponse.ClerkActions{SecretKey: os.Getenv("CLERK_SECRET_KEY")} // or WorkOSActions, Auth0Actions
result, err := noresponse.HandleLateResult(r.Context(), payload, r.Header, secrets, provider, nil)
```

`LateResultOptions` sets `Policy` (recommendation to `LateResultNone`, `LateResultRevokeSessions` or `LateResultBan`), `ResolveUserID` (account id to provider user id, with the context; an error it returns fails the delivery so it is retried) and `Webhook` options. A test event from `SendTestEvent` acts on nothing. A refusal from the provider is a `*ProviderActionError`; answer 5xx so the delivery is retried. A provider 404 means the user or session is already gone and counts as done, so a repeated delivery is safe. WorkOS has no block, so its ban deletes the user permanently.

If your app has its own login, `&noresponse.LogOnlyActions{}` logs what it would revoke or ban (an Info record on `slog.Default()`, or set `Logger`) and calls nothing. Use it to try the handler, then write your own `LateResultProvider`: `Name()`, `RevokeSessions(ctx, userID)` and `Ban(ctx, userID)`.

## Other options

`noresponse.WithServerURL`, `noresponse.WithClient` (any `Do(*http.Request)` implementation) and `noresponse.WithSecuritySource` (a key fetched per call, for rotation) are available on both `New` and `NewClient`.


## Changes since 2.0.0

No name was renamed. New in the root package: `NewClient`, `WithAPIKey` (the same as `WithSecurity`), `NewIdempotencyKey` and `AllAssessments`. The module now needs Go 1.24, where 2.0.0 declared Go 1.26.8.
