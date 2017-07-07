package sc2

import (
	"fmt"
	"net/http"
	"testing"
	"reflect"
)

const matchesResp = `{
			"matches": [{
				"map": "foobar",
				"type": "SOLO",
				"decision": "WIN",
				"speed": "NORMAL",
				"date": 123456789
			}]
		}`

func TestProfileService_Matches(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/profile/1234567/1/foobar/matches", func (w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, matchesResp)
	})
	actual, _, err := client.Profile().Matches(1234567, 1, "foobar")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	want := Match{  Map: "foobar",
			Type: "SOLO",
			Decision: "WIN",
			Speed: "NORMAL",
			Date: 123456789,}

	if !reflect.DeepEqual(actual.Matches[0], want) {
		t.Fatalf("returned %+v, want %+v", actual, want)
	}
}
