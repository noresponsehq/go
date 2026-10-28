package noresponse

import (
	"context"
	"fmt"
	"github.com/noresponsehq/go/v8/internal/config"
	"github.com/noresponsehq/go/v8/internal/hooks"
	"github.com/noresponsehq/go/v8/internal/utils"
	"github.com/noresponsehq/go/v8/models/components"
	"github.com/noresponsehq/go/v8/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production
	"https://api.noresponse.org",
}

// HTTPClient provides an interface for supplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

// NoResponse API: Server-side risk assessment of account creation, login and purchase: a recommendation, polling and signed webhooks.
// Assess an account creation, login or purchase and get back the account's one current risk, poll a pending assessment, page through an account's history, record what you later learned, report usage of the API keys you issue to your accounts, rotate your API key, manage your webhook endpoints, and register hosted endpoints that answer your WorkOS, Auth0 or Supabase sign-up and sign-in hooks. Every response carries X-Request-ID; every error is an RFC 9457 problem document. Keep the API key on your server.
type NoResponse struct {
	SDKVersion string
	// Assessments of account creation, login and purchase.
	Assessments *Assessments
	// Your accounts, each with one current risk, and their assessment history.
	Accounts *Accounts
	KeyUsage *KeyUsage
	// Your own API keys.
	Keys *Keys
	// Hosted endpoints that answer your sign-up and sign-in hooks.
	Integrations *Integrations
	// Your webhook endpoints: register, list, rotate signing keys, send test events and delete.
	WebhookEndpoints *WebhookEndpoints
	Settings         *Settings
	Requests         *Requests
	Lists            *Lists

	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

type SDKOption func(*NoResponse)

// WithServerURL allows providing an alternative server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *NoResponse) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *NoResponse) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *NoResponse) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *NoResponse) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKey string) SDKOption {
	return func(sdk *NoResponse) {
		security := components.Security{APIKey: &apiKey}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *NoResponse) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *NoResponse) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *NoResponse) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *NoResponse {
	sdk := &NoResponse{
		SDKVersion: "8.4.1",
		sdkConfiguration: config.SDKConfiguration{
			UserAgent:         "noresponse-sdk/go 8.4.1",
			SDKVersion:        "8.4.1",
			GenVersion:        "internal",
			OpenAPIDocVersion: "1",
			ServerList:        ServerList,
		},
		hooks: hooks.New(),
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk.sdkConfiguration.Security == nil {
		var envVarSecurity components.Security
		if utils.PopulateSecurityFromEnv(&envVarSecurity) {
			sdk.sdkConfiguration.Security = utils.AsSecuritySource(envVarSecurity)
		}
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	sdk.sdkConfiguration = sdk.hooks.SDKInit(sdk.sdkConfiguration)

	sdk.Assessments = newAssessments(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Accounts = newAccounts(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.KeyUsage = newKeyUsage(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Keys = newKeys(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Integrations = newIntegrations(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.WebhookEndpoints = newWebhookEndpoints(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Settings = newSettings(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Requests = newRequests(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Lists = newLists(sdk, sdk.sdkConfiguration, sdk.hooks)

	return sdk
}
