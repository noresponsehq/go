package noresponse

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/noresponsehq/go/v19/models/components"
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
	var path string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
	defer server.Close()

	client := NewClient("nrk_test_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	res, err := client.Groups.GetPerson(context.Background(), ids.Person.GroupID, ids.Person.ID)
	if err != nil {
		t.Fatal(err)
	}
	if want := "/v1/groups/" + ids.Person.GroupID + "/people/" + ids.Person.ID; path != want {
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
	if *row.Likelihood != 0.8 || row.Verdict != nil {
		t.Fatalf("row linkage: %v %v", *row.Likelihood, row.Verdict)
	}
	if remaining := lookup.Remaining[0]; remaining.Source.Kind != "credential_list" || remaining.By != "username" || remaining.Rows != 37 {
		t.Fatalf("remaining: %+v", remaining)
	}
}

// The record the API serves for an account with every member set keeps its
// zeros, false values, nested values, timestamps and source through the SDK.
func TestGetPersonKeepsEveryMemberOfAnAccountsRecord(t *testing.T) {
	report, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "person-report-records.json"))
	if err != nil {
		t.Fatal(err)
	}
	served, err := os.ReadFile(filepath.Join("tests", "fixtures", "valid", "record-every-account-member.json"))
	if err != nil {
		t.Fatal(err)
	}
	var body map[string]any
	if err := json.Unmarshal(report, &body); err != nil {
		t.Fatal(err)
	}
	var record any
	if err := json.Unmarshal(served, &record); err != nil {
		t.Fatal(err)
	}
	body["lookup"].(map[string]any)["records"] = []any{record}
	person := body["person"].(map[string]any)
	answer, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(answer)
	}))
	defer server.Close()

	client := NewClient("nrk_test_kkkkkkkkkkkkkkkkkkkk", WithServerURL(server.URL))
	res, err := client.Groups.GetPerson(context.Background(), person["group_id"].(string), person["id"].(string))
	if err != nil {
		t.Fatal(err)
	}
	lookup := res.Result.GetLookup()
	if lookup == nil || len(lookup.Records) != 1 {
		t.Fatalf("lookup: %+v", lookup)
	}
	account := lookup.Records[0]
	if account.Source.Connector != "bluesky" || account.Source.Kind != "platform" || account.Source.GetSite() == nil || *account.Source.GetSite() != "bsky.app" {
		t.Fatalf("source: %+v", account.Source)
	}
	var want struct {
		Fields []struct {
			Field    string `json:"field"`
			Original string `json:"original"`
			Value    string `json:"value"`
		} `json:"fields"`
	}
	if err := json.Unmarshal(served, &want); err != nil {
		t.Fatal(err)
	}
	if len(account.Fields) != len(want.Fields) {
		t.Fatalf("%d fields, the API served %d", len(account.Fields), len(want.Fields))
	}
	for index, field := range account.Fields {
		expected := want.Fields[index]
		if field.Field != expected.Field || field.Original != expected.Original || field.Value == nil || *field.Value != expected.Value {
			t.Fatalf("field %d: %+v, the API served %+v", index, field, expected)
		}
	}
	for name, original := range map[string]string{"followers": "0", "verified": "false", "associated.lists.2": "true",
		"creation_date": "2023-04-30T10:15:00Z", "contacts.0.since": "2024-01-09T00:00:00Z"} {
		found := false
		for _, field := range account.Fields {
			found = found || field.Field == name && field.Original == original
		}
		if !found {
			t.Fatalf("no field %s holding %s", name, original)
		}
	}
}
