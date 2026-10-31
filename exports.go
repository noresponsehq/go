package noresponse

import (
	"bytes"
	"context"
	"fmt"
	"github.com/noresponsehq/go/v10/internal/config"
	"github.com/noresponsehq/go/v10/internal/hooks"
	"github.com/noresponsehq/go/v10/internal/utils"
	"github.com/noresponsehq/go/v10/models/apierrors"
	"github.com/noresponsehq/go/v10/models/components"
	"github.com/noresponsehq/go/v10/models/operations"
	"github.com/noresponsehq/go/v10/retry"
	"github.com/spyzhov/ajson"
	"net/http"
	"strconv"
	"strings"
)

// Exports - Bounded customer export jobs and encrypted artifacts.
type Exports struct {
	rootSDK          *NoResponse
	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

func newExports(rootSDK *NoResponse, sdkConfig config.SDKConfiguration, hooks *hooks.Hooks) *Exports {
	return &Exports{
		rootSDK:          rootSDK,
		sdkConfiguration: sdkConfig,
		hooks:            hooks,
	}
}

// Create - Request an asynchronous customer export.
// Request an encrypted JSON Lines snapshot of retained account or audit records. Requires exports plus history or audit scope. An account export without legacy history needs investigate with evidence or detailed, because it holds submitted identifiers; detailed account data always needs evidence or detailed. Criteria are frozen and hashed. Each environment permits four active jobs, 64 MiB per artifact, and 256 KiB per chunk. Artifacts expire after 24 hours. Send an Idempotency-Key to safely repeat a request. Every operation checks current permissions and source retention.
func (s *Exports) Create(ctx context.Context, workspaceID string, body components.CustomerExportRequest, idempotencyKey *string, opts ...operations.Option) (*operations.CreateCustomerExportResponse, error) {
	request := operations.CreateCustomerExportRequest{
		WorkspaceID:    workspaceID,
		Body:           body,
		IdempotencyKey: idempotencyKey,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionTimeout,
		operations.SupportedOptionSkipDeserialization,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v1/workspaces/{workspace_id}/exports", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "createCustomerExport",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Body", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	utils.PopulateHeaders(ctx, req, request, nil)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.CreateCustomerExportResponse{}

	switch {
	case httpRes.StatusCode == 202:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CustomerExport
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Result = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.AuthenticationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.PermissionDeniedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.RateLimitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// List exports for one collection.
// List authorised jobs for one kind and environment. Detailed jobs require current detailed-data permission and are filtered before pagination. Keep the kind and environment unchanged when following next_cursor.
func (s *Exports) List(ctx context.Context, workspaceID string, kind components.CustomerExportRequestKind, environment components.Environment, limit *int64, cursor *string, opts ...operations.Option) (*operations.ListCustomerExportsResponse, error) {
	request := operations.ListCustomerExportsRequest{
		WorkspaceID: workspaceID,
		Kind:        kind,
		Environment: environment,
		Limit:       limit,
		Cursor:      cursor,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
		operations.SupportedOptionSkipDeserialization,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v1/workspaces/{workspace_id}/exports", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "listCustomerExports",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}
	paginationCtx := ctx

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     8000,
					Exponent:        2,
					MaxElapsedTime:  30000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"5XX",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ListCustomerExportsResponse{}
	res.Next = func() (*operations.ListCustomerExportsResponse, error) {
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}

		b, err := ajson.Unmarshal(rawBody)
		if err != nil {
			return nil, err
		}
		nC, err := ajson.Eval(b, "$.next_cursor")
		if err != nil {
			return nil, err
		}
		var nCVal string

		if nC.IsNumeric() {
			numVal, err := nC.GetNumeric()
			if err != nil {
				return nil, err
			}
			// GetNumeric returns as float64 so convert to the appropriate type.
			nCVal = strconv.FormatFloat(numVal, 'f', 0, 64)
		} else {
			val, err := nC.Value()
			if err != nil {
				return nil, err
			}
			if val == nil {
				return nil, nil
			}
			nCVal = val.(string)
			if strings.TrimSpace(nCVal) == "" {
				return nil, nil
			}
		}

		return s.List(
			paginationCtx,
			workspaceID,
			kind,
			environment,
			limit,
			&nCVal,
			opts...,
		)
	}

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CustomerExportList
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Result = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.AuthenticationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.PermissionDeniedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.RateLimitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// Get - Read export state and progress.
// Read the status of an authorised export. Records and bytes describe committed progress in the current generation. Only completed, unexpired artifacts can be downloaded. Source erasure, retention changes and workspace removal can invalidate a job.
func (s *Exports) Get(ctx context.Context, workspaceID string, exportID string, opts ...operations.Option) (*operations.GetCustomerExportResponse, error) {
	request := operations.GetCustomerExportRequest{
		WorkspaceID: workspaceID,
		ExportID:    exportID,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
		operations.SupportedOptionSkipDeserialization,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v1/workspaces/{workspace_id}/exports/{export_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "getCustomerExport",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     8000,
					Exponent:        2,
					MaxElapsedTime:  30000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"5XX",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.GetCustomerExportResponse{}

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CustomerExport
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Result = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.AuthenticationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.PermissionDeniedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.RateLimitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// Cancel and remove an export artifact.
// Cancel an export and remove its stored chunks. Repeating cancellation is safe. An already admitted download may finish; bytes already returned cannot be recalled. Send an empty JSON object.
func (s *Exports) Cancel(ctx context.Context, workspaceID string, exportID string, body components.CustomerExportActionRequest, opts ...operations.Option) (*operations.CancelCustomerExportResponse, error) {
	request := operations.CancelCustomerExportRequest{
		WorkspaceID: workspaceID,
		ExportID:    exportID,
		Body:        body,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
		operations.SupportedOptionSkipDeserialization,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v1/workspaces/{workspace_id}/exports/{export_id}/cancel", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "cancelCustomerExport",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Body", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     8000,
					Exponent:        2,
					MaxElapsedTime:  30000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"5XX",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.CancelCustomerExportResponse{}

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CustomerExport
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Result = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.AuthenticationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.PermissionDeniedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.RateLimitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// Retry a failed export.
// Start a new generation of a failed export before its original expiry, up to ten generations. The request criteria stay fixed and the new generation takes a fresh source snapshot. Cancelled, erased and expired jobs cannot be retried. Send an empty JSON object.
func (s *Exports) Retry(ctx context.Context, workspaceID string, exportID string, body components.CustomerExportActionRequest, opts ...operations.Option) (*operations.RetryCustomerExportResponse, error) {
	request := operations.RetryCustomerExportRequest{
		WorkspaceID: workspaceID,
		ExportID:    exportID,
		Body:        body,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionTimeout,
		operations.SupportedOptionSkipDeserialization,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v1/workspaces/{workspace_id}/exports/{export_id}/retry", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "retryCustomerExport",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Body", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.RetryCustomerExportResponse{}

	switch {
	case httpRes.StatusCode == 202:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CustomerExport
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Result = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.AuthenticationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.PermissionDeniedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.RateLimitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// DownloadChunk - Download one authorized export chunk.
// Download one chunk of a completed, unexpired export after a fresh permission and privacy check. The response contains base64-encoded JSON Lines, a SHA-256, generation and next index. Decode and concatenate chunks in order from one generation. Audit storage failure withholds the bytes.
func (s *Exports) DownloadChunk(ctx context.Context, workspaceID string, exportID string, chunkIndex int64, opts ...operations.Option) (*operations.DownloadCustomerExportChunkResponse, error) {
	request := operations.DownloadCustomerExportChunkRequest{
		WorkspaceID: workspaceID,
		ExportID:    exportID,
		ChunkIndex:  chunkIndex,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
		operations.SupportedOptionSkipDeserialization,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v1/workspaces/{workspace_id}/exports/{export_id}/chunks/{chunk_index}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "downloadCustomerExportChunk",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 500,
					MaxInterval:     8000,
					Exponent:        2,
					MaxElapsedTime:  30000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"5XX",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.DownloadCustomerExportChunkResponse{}

	switch {
	case httpRes.StatusCode == 200:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CustomerExportChunk
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Result = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.AuthenticationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.PermissionDeniedError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.NotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 409:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ConflictError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 429:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.RateLimitError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, apierrors.NewResponseValidationError("response did not match the declared error schema", httpRes.StatusCode, string(rawBody), httpRes, err)
			}

			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIStatusError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIStatusError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
