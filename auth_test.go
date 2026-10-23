package noresponse

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type authVector struct {
	Name      string         `json:"name"`
	Provider  string         `json:"provider"`
	Claims    map[string]any `json:"claims"`
	Overrides map[string]any `json:"overrides"`
	Expect    map[string]any `json:"expect"`
	Error     bool           `json:"error"`
}

func readAuthFixture(t *testing.T, name string, into any) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("tests", "fixtures", name))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, into); err != nil {
		t.Fatal(err)
	}
}

func TestAuthContextVectors(t *testing.T) {
	var file struct {
		Cases []authVector `json:"cases"`
	}
	readAuthFixture(t, "auth-context-vectors.json", &file)
	if len(file.Cases) == 0 {
		t.Fatal("no vectors")
	}
	for _, vector := range file.Cases {
		t.Run(vector.Name, func(t *testing.T) {
			var got any
			var err error
			switch vector.Provider {
			case "workos":
				got, err = FromWorkOSClaims(vector.Claims, vector.Overrides)
			case "clerk":
				got, err = FromClerkClaims(vector.Claims, vector.Overrides)
			case "auth0":
				got, err = FromAuth0Claims(vector.Claims, vector.Overrides)
			case "supabase":
				got, err = FromSupabaseClaims(vector.Claims, vector.Overrides)
			case "firebase":
				got, err = FromFirebaseClaims(vector.Claims, vector.Overrides)
			case "cognito":
				got, err = FromCognitoClaims(vector.Claims, vector.Overrides)
			case "entra":
				got, err = FromEntraClaims(vector.Claims, vector.Overrides)
			case "keycloak":
				got, err = FromKeycloakClaims(vector.Claims, vector.Overrides)
			case "ory":
				got, err = FromOrySession(vector.Claims, vector.Overrides)
			case "ory_oauth2":
				got, err = FromOryOAuth2Claims(vector.Claims, vector.Overrides)
			default:
				got, err = FromOIDCClaims(vector.Provider, vector.Claims, vector.Overrides)
			}
			if vector.Error {
				if err == nil {
					t.Fatalf("expected an error, got %v", got)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(got)
			if err != nil {
				t.Fatal(err)
			}
			var wire map[string]any
			if err := json.Unmarshal(encoded, &wire); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(wire, vector.Expect) {
				t.Fatalf("got %v\nwant %v", wire, vector.Expect)
			}
		})
	}
}

// A token can put an object or an array in any claim. Comparing one with a
// marker's value must not panic, and only an equal scalar matches.
func TestMarkerRuleHandlesEveryClaimType(t *testing.T) {
	var rule authRule
	ruleJSON := `{"op": "marker", "value": "low", "when": [
		{"claim": "kind", "equals": "anonymous"},
		{"claim": "count", "equals": 3},
		{"claim": "flag", "equals": true},
		{"claim": "shape", "equals": {"v": 1}}]}`
	if err := json.Unmarshal([]byte(ruleJSON), &rule); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		claims string
		match  bool
	}{
		{"object claims", `{"kind": {"v": "anonymous"}, "count": {"v": 3}, "flag": {"v": true}}`, false},
		{"array claims", `{"kind": ["anonymous"], "count": [3], "flag": [true]}`, false},
		{"object claim against an object marker", `{"shape": {"v": 1}}`, false},
		{"number claim equal", `{"count": 3}`, true},
		{"number claim different", `{"count": 4, "kind": 3}`, false},
		{"string claim equal", `{"kind": "anonymous"}`, true},
		{"string claim different", `{"kind": "password", "count": "3", "flag": "true"}`, false},
		{"bool claim equal", `{"flag": true}`, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var claims map[string]any
			if err := json.Unmarshal([]byte(tc.claims), &claims); err != nil {
				t.Fatal(err)
			}
			got := authMapRule(claims, rule)
			if tc.match && got != "low" {
				t.Fatalf("got %#v, want \"low\"", got)
			}
			if !tc.match && got != nil {
				t.Fatalf("got %#v, want no match", got)
			}
		})
	}
}

func TestEmbeddedClaimMapIsTheContractMap(t *testing.T) {
	var table map[string]any
	readAuthFixture(t, "auth-context-claim-map.json", &table)
	delete(table, "description")
	delete(table, "rules")
	for _, provider := range table["providers"].(map[string]any) {
		delete(provider.(map[string]any), "notes")
	}
	var embedded map[string]any
	if err := json.Unmarshal([]byte(authClaimMapJSON), &embedded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(embedded, table) {
		t.Fatal("the embedded claim map differs from tests/fixtures/auth-context-claim-map.json")
	}
}
