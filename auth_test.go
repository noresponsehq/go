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
