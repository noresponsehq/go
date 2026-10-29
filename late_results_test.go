package noresponse

// The shared late enforcement vectors through HandleRiskChange, plus what the
// vectors cannot say in JSON: verification, the custom provider and missing
// credentials.

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/noresponsehq/go/v9/models/components"
)

type lateExchange struct {
	Request struct {
		Method string          `json:"method"`
		URL    string          `json:"url"`
		Body   json.RawMessage `json:"body"`
	} `json:"request"`
	Response struct {
		Status  int               `json:"status"`
		Headers map[string]string `json:"headers"`
		Body    json.RawMessage   `json:"body"`
	} `json:"response"`
}

type lateVector struct {
	Name     string                              `json:"name"`
	Provider string                              `json:"provider"`
	Domain   string                              `json:"domain"`
	Flows    map[string]struct{ Reinstate bool } `json:"flows"`
	Event    json.RawMessage                     `json:"event"`
	State    *struct {
		State               string `json:"state"`
		UserID              string `json:"user_id"`
		EnforcedRiskVersion *int64 `json:"enforced_risk_version"`
	} `json:"state"`
	Exchanges []lateExchange `json:"exchanges"`
	Result    struct {
		Skipped      *string          `json:"skipped"`
		Enforcements []map[string]any `json:"enforcements"`
	} `json:"result"`
	Error *struct {
		Provider string `json:"provider"`
		Status   int    `json:"status"`
		Code     string `json:"code"`
	} `json:"error"`
}

type lateVectors struct {
	APIBaseURL    string       `json:"api_base_url"`
	APIKey        string       `json:"api_key"`
	Secret        string       `json:"secret_utf8"`
	KeyID         string       `json:"key_id"`
	Now           int64        `json:"now"`
	ProviderToken string       `json:"provider_token"`
	Cases         []lateVector `json:"cases"`
}

func loadLateVectors(t *testing.T) lateVectors {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", "late-results-vectors.json"))
	if err != nil {
		t.Fatal(err)
	}
	var vectors lateVectors
	if err := json.Unmarshal(raw, &vectors); err != nil {
		t.Fatal(err)
	}
	return vectors
}

func (v lateVectors) sign(event []byte) http.Header {
	stamp := strconv.FormatInt(v.Now, 10)
	mac := hmac.New(sha256.New, []byte(v.Secret))
	mac.Write([]byte(stamp + "."))
	mac.Write(event)
	var id struct {
		ID string `json:"id"`
	}
	_ = json.Unmarshal(event, &id)
	header := http.Header{}
	header.Set("NR-Webhook-Id", id.ID)
	header.Set("NR-Webhook-Timestamp", stamp)
	header.Set("NR-Webhook-Key-Id", v.KeyID)
	header.Set("NR-Webhook-Signature", "v1="+hex.EncodeToString(mac.Sum(nil)))
	return header
}

func (v lateVectors) secrets() map[string]string { return map[string]string{v.KeyID: v.Secret} }

func (v lateVectors) clock() WebhookOption {
	return WithWebhookClock(func() time.Time { return time.Unix(v.Now, 0) })
}

// normalizeRequest is the method, path and query as an order-free set of decoded pairs.
func normalizeRequest(method string, u *url.URL) string {
	var pairs []string
	for key, values := range u.Query() {
		for _, value := range values {
			pairs = append(pairs, key+"="+value)
		}
	}
	sort.Strings(pairs)
	return method + " " + u.Scheme + "://" + u.Host + u.EscapedPath() + "?" + strings.Join(pairs, "&")
}

// transcript serves the exchanges in order and fails on anything else.
type transcript struct {
	t         *testing.T
	exchanges []lateExchange
	next      int
}

func (tr *transcript) Do(req *http.Request) (*http.Response, error) {
	tr.t.Helper()
	got := normalizeRequest(req.Method, req.URL)
	if tr.next >= len(tr.exchanges) {
		tr.t.Fatalf("unexpected request %s", got)
	}
	expected := tr.exchanges[tr.next]
	want, err := url.Parse(expected.Request.URL)
	if err != nil {
		tr.t.Fatal(err)
	}
	if wantText := normalizeRequest(expected.Request.Method, want); got != wantText {
		tr.t.Fatalf("request %d: got %s, want %s", tr.next, got, wantText)
	}
	if len(expected.Request.Body) > 0 {
		var body []byte
		if req.Body != nil {
			body, _ = io.ReadAll(req.Body)
		}
		var gotBody, wantBody any
		if err := json.Unmarshal(body, &gotBody); err != nil {
			tr.t.Fatalf("request %d body %q: %v", tr.next, body, err)
		}
		_ = json.Unmarshal(expected.Request.Body, &wantBody)
		if !reflect.DeepEqual(gotBody, wantBody) {
			tr.t.Fatalf("body of request %d: got %s, want %s", tr.next, body, expected.Request.Body)
		}
	}
	tr.next++
	header := http.Header{}
	for key, value := range expected.Response.Headers {
		header.Set(key, value)
	}
	content := []byte(nil)
	if len(expected.Response.Body) > 0 && expected.Response.Status != http.StatusNoContent {
		content = expected.Response.Body
	}
	return &http.Response{
		StatusCode: expected.Response.Status,
		Header:     header,
		Body:       io.NopCloser(bytes.NewReader(content)),
		Request:    req,
	}, nil
}

func (tr *transcript) done() {
	tr.t.Helper()
	if tr.next != len(tr.exchanges) {
		tr.t.Fatalf("%d of %d exchanges were requested", tr.next, len(tr.exchanges))
	}
}

func TestLateResultVectors(t *testing.T) {
	vectors := loadLateVectors(t)
	if len(vectors.Cases) == 0 {
		t.Fatal("no vectors")
	}
	for _, vector := range vectors.Cases {
		t.Run(vector.Name, func(t *testing.T) {
			tr := &transcript{t: t, exchanges: vector.Exchanges}
			client := NewClient(vectors.APIKey, WithServerURL(vectors.APIBaseURL), WithClient(tr))
			provider, err := IdentityProvider(vector.Provider, IdentityProviderOptions{Token: vectors.ProviderToken, Domain: vector.Domain, HTTPClient: tr})
			if err != nil {
				t.Fatal(err)
			}
			var event struct {
				Data struct {
					Account struct {
						ID string `json:"id"`
					} `json:"account"`
				} `json:"data"`
			}
			_ = json.Unmarshal(vector.Event, &event)
			flows := map[components.Flow]FlowPolicy{}
			for flow, policy := range vector.Flows {
				flows[components.Flow(flow)] = FlowPolicy{Reinstate: policy.Reinstate}
			}
			opts := &EnforcementOptions{
				Provider: provider,
				Flows:    flows,
				Webhook:  []WebhookOption{vectors.clock()},
				GetAccountState: func(_ context.Context, account components.Account) (*AccountStateAnswer, error) {
					if account.ID != event.Data.Account.ID {
						t.Errorf("getAccountState got account %s", account.ID)
					}
					if vector.State == nil {
						return nil, nil
					}
					return &AccountStateAnswer{
						State:               components.EnforcementState(vector.State.State),
						UserID:              vector.State.UserID,
						EnforcedRiskVersion: vector.State.EnforcedRiskVersion,
					}, nil
				},
			}
			result, err := HandleRiskChange(context.Background(), client, vector.Event, vectors.sign(vector.Event), vectors.secrets(), opts)
			tr.done()
			if vector.Error != nil {
				var refused *ProviderActionError
				if !errors.As(err, &refused) {
					t.Fatalf("expected *ProviderActionError, got %v", err)
				}
				got := [3]any{refused.Provider, refused.Status, refused.Code}
				want := [3]any{vector.Error.Provider, vector.Error.Status, vector.Error.Code}
				if got != want {
					t.Fatalf("error %v, want %v", got, want)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			wantSkip := ""
			if vector.Result.Skipped != nil {
				wantSkip = *vector.Result.Skipped
			}
			if string(result.Skipped) != wantSkip {
				t.Fatalf("skipped %q, want %q", result.Skipped, wantSkip)
			}
			got := []map[string]any{}
			for _, e := range result.Enforcements {
				entry := map[string]any{"flow": string(e.Flow), "action": string(e.Action), "status": string(e.Status)}
				if e.ErrorCode != "" {
					entry["error_code"] = e.ErrorCode
				}
				got = append(got, entry)
			}
			want := vector.Result.Enforcements
			if want == nil {
				want = []map[string]any{}
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("enforcements %v, want %v", got, want)
			}
		})
	}
}

func TestLateResultVerificationFailureReadsNothing(t *testing.T) {
	vectors := loadLateVectors(t)
	tr := &transcript{t: t}
	client := NewClient(vectors.APIKey, WithServerURL(vectors.APIBaseURL), WithClient(tr))
	event := vectors.Cases[0].Event
	opts := &EnforcementOptions{
		Provider: CustomProvider(nil),
		Webhook:  []WebhookOption{vectors.clock()},
		GetAccountState: func(context.Context, components.Account) (*AccountStateAnswer, error) {
			t.Fatal("not called")
			return nil, nil
		},
	}
	tampered := append(append([]byte{}, event...), ' ')
	_, err := HandleRiskChange(context.Background(), client, tampered, vectors.sign(event), vectors.secrets(), opts)
	var refused *WebhookVerificationError
	if !errors.As(err, &refused) {
		t.Fatalf("expected *WebhookVerificationError, got %v", err)
	}
	tr.done()
}

func TestCustomProviderMissingHandlerFailsTheAction(t *testing.T) {
	vectors := loadLateVectors(t)
	const account = "acct_7f3a9c"
	record := func(action, status, code string) lateExchange {
		var e lateExchange
		e.Request.Method = "POST"
		e.Request.URL = vectors.APIBaseURL + "/v1/accounts/" + account + "/enforcements"
		body := map[string]any{"risk_version": 9, "flow": "login", "action": action, "state_before": "active", "status": status, "provider": "custom"}
		stored := map[string]any{"id": "enf_01JA2ZF5N6P7Q8R9S0T1V2W3X4", "account_id": account, "created_at": "2026-10-23T09:00:06.000Z"}
		if code != "" {
			body["error_code"] = code
		}
		for key, value := range body {
			stored[key] = value
		}
		e.Request.Body, _ = json.Marshal(body)
		e.Response.Status = 201
		e.Response.Headers = map[string]string{"content-type": "application/json"}
		e.Response.Body, _ = json.Marshal(stored)
		return e
	}
	tr := &transcript{t: t, exchanges: []lateExchange{record("restrict", "applied", ""), record("revoke_sessions", "failed", "custom.error")}}
	client := NewClient(vectors.APIKey, WithServerURL(vectors.APIBaseURL), WithClient(tr))
	var ran []string
	provider := CustomProvider(map[components.EnforcementAction]func(context.Context, string) ([]components.EnforcementAction, error){
		components.EnforcementActionRestrict: func(_ context.Context, userID string) ([]components.EnforcementAction, error) {
			ran = append(ran, "restrict "+userID)
			return nil, nil
		},
	})
	opts := &EnforcementOptions{
		Provider: provider,
		GetAccountState: func(context.Context, components.Account) (*AccountStateAnswer, error) {
			return &AccountStateAnswer{State: components.EnforcementStateActive, UserID: "u_1"}, nil
		},
	}
	decisions := components.AccountDecisionsByFlow{AccountCreation: "block", Login: "block", Purchase: "block"}
	_, err := EnforceAccount(context.Background(), client, components.Account{ID: account, RiskVersion: 9, DecisionsByFlow: decisions}, opts)
	if err == nil {
		t.Fatal("a missing handler must fail the action")
	}
	tr.done()
	if !reflect.DeepEqual(ran, []string{"restrict u_1"}) {
		t.Fatalf("ran %v", ran)
	}
}

func TestIdentityProviderNeedsItsCredentials(t *testing.T) {
	for _, c := range []struct {
		name    string
		options IdentityProviderOptions
	}{
		{"clerk", IdentityProviderOptions{}},
		{"auth0", IdentityProviderOptions{Token: "t"}},
		{"okta", IdentityProviderOptions{Token: "t"}},
	} {
		if _, err := IdentityProvider(c.name, c.options); err == nil {
			t.Errorf("%s %+v: expected an error", c.name, c.options)
		}
	}
}
