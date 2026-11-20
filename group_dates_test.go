package noresponse

import (
	"encoding/json"
	"github.com/noresponsehq/go/v19/models/components"
	"github.com/noresponsehq/go/v19/optionalnullable"
	"reflect"
	"testing"
)

func TestGroupDatesPreserveNullAndOmission(t *testing.T) {
	start, end := "2026-12-03T09:00:00Z", "2026-12-03T17:00:00Z"
	for _, tc := range []struct {
		body components.GroupUpdate
		want string
	}{
		{components.GroupUpdate{}, `{}`},
		{components.GroupUpdate{StartsAt: optionalnullable.From[string](nil), EndsAt: optionalnullable.From[string](nil)}, `{"starts_at":null,"ends_at":null}`},
		{components.GroupUpdate{StartsAt: optionalnullable.From(&start), EndsAt: optionalnullable.From(&end)}, `{"starts_at":"2026-12-03T09:00:00Z","ends_at":"2026-12-03T17:00:00Z"}`},
	} {
		got, err := json.Marshal(tc.body)
		var actual, expected any
		json.Unmarshal(got, &actual)
		json.Unmarshal([]byte(tc.want), &expected)
		if err != nil || !reflect.DeepEqual(actual, expected) {
			t.Fatalf("got %s (%v), want %s", got, err, tc.want)
		}
	}
	var group components.Group
	if err := json.Unmarshal([]byte(`{"id":"grp_example","environment":"test","name":"Harbour building","type":"event","attribute":"","starts_at":null,"ends_at":null,"principals":[],"topics":[],"people":0,"version":1,"created_by":"test","created_at":"2026-11-15T10:00:00Z","updated_at":"2026-11-15T10:00:00Z"}`), &group); err != nil {
		t.Fatal(err)
	}
	if group.StartsAt != nil || group.EndsAt != nil {
		t.Fatal("ongoing group acquired dates")
	}
}

func TestRelatedOngoingGroupDecodesNullStart(t *testing.T) {
	var group components.OtherGroup
	if err := json.Unmarshal([]byte(`{"group_id":"grp_building","group_name":"Harbour building","group_starts_at":null,"person_id":"per_example","matched_on":["person_id"],"match_state":"unscreened","decision":null}`), &group); err != nil {
		t.Fatal(err)
	}
	if group.GroupStartsAt != nil {
		t.Fatal("ongoing group acquired a date")
	}
}
