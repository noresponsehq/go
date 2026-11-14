package noresponse

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/noresponsehq/go/v12/models/components"
)

// A person's report returns the contract's worked records as typed models:
// a platform account and a leaked row, with kinds, roles, ratings and
// likelihood.
func TestGetPersonReturnsTypedRecords(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "person-report-records.json"))
	if err != nil {
		t.Fatal(err)
	}
	var ids struct {
		Person struct {
			ID      string `json:"id"`
			GroupID string `json:"group_id"`
		} `json:"person"`
	}
	if err := json.Unmarshal(body, &ids); err != nil {
		t.Fatal(err)
	}
	const workspace = "tnt_01JC8Y0A2B4C6D8E0F2G4H6J8K"
	var path string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	client := NewClient("nrk_test_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	res, err := client.Groups.GetPerson(context.Background(), workspace, components.KeyEnvironmentTest, ids.Person.GroupID, ids.Person.ID)
	if err != nil {
		t.Fatal(err)
	}
	if want := "/v1/workspaces/" + workspace + "/groups/test/" + ids.Person.GroupID + "/people/" + ids.Person.ID; path != want {
		t.Fatalf("path %s, want %s", path, want)
	}
	lookup := res.Result.GetLookup()
	if lookup == nil || len(lookup.Records) != 2 || len(lookup.Remaining) != 1 {
		t.Fatalf("lookup: %+v", lookup)
	}
	account, row := lookup.Records[0], lookup.Records[1]
	field := func(record components.Record, name string) components.RecordField {
		for _, each := range record.Fields {
			if each.Field == name {
				return each
			}
		}
		t.Fatalf("no field %s in %s", name, record.ID)
		return components.RecordField{}
	}

	if account.Source.Kind != "platform" || account.Source.GetSite() == nil || *account.Source.GetSite() != "github.com" {
		t.Fatalf("account source: %+v", account.Source)
	}
	login := field(account, "login")
	if *login.Kind != components.RecordKindUsername || *login.Role != components.RecordRoleHandle || *login.Value != "jordan-ellis" ||
		login.Distinctive.Rating != components.RatingDistinctive {
		t.Fatalf("login: %+v", login)
	}
	if name := field(account, "name"); name.Original != "Jordan  Ellis" || *name.Value != "Jordan Ellis" {
		t.Fatalf("name: %+v", name)
	}
	if hireable := field(account, "hireable"); hireable.Role != nil {
		t.Fatalf("a field only its source returns has a role: %+v", hireable)
	}
	if *account.Likelihood != 0.74 || !*account.Joined || *account.Verdict != components.VerdictJoined || len(account.Agreed) != 2 {
		t.Fatalf("account linkage: %v %v %v %v", *account.Likelihood, *account.Joined, *account.Verdict, account.Agreed)
	}

	if row.Source.Kind != "site_breach" || *row.Source.Table != "FORUM_EXAMPLE_NET_2019" || *row.Source.Date != "2019-03" {
		t.Fatalf("row source: %+v", row.Source)
	}
	if password := field(row, "password"); *password.Kind != components.RecordKindPassword || *password.Value != "Leeds1990!" {
		t.Fatalf("password: %+v", password)
	}
	if hash := field(row, "password_hash"); *hash.Detail != "bcrypt" {
		t.Fatalf("hash: %+v", hash)
	}
	if ip := field(row, "lastip"); *ip.Role != components.RecordRoleLastIP || *ip.Value != "203.0.113.5" {
		t.Fatalf("lastip: %+v", ip)
	}
	if birthday := field(row, "birthday"); birthday.Value != nil || birthday.Original != "07/03/1990" {
		t.Fatalf("an ambiguous date: %+v", birthday)
	}
	if *row.Likelihood != 0.99 || row.Verdict != nil {
		t.Fatalf("row linkage: %v %v", *row.Likelihood, row.Verdict)
	}
	if remaining := lookup.Remaining[0]; remaining.Source.Kind != "credential_list" || remaining.By != "username" || remaining.Rows != 37 {
		t.Fatalf("remaining: %+v", remaining)
	}
}
