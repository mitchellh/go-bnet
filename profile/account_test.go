package profile

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_User(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/account/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"id": 12345,
			"battletag": "foobar#1234"
		}`)
	})

	actual, _, err := client.Account().User()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	want := &User{ID: 12345, BattleTag: "foobar#1234"}
	if !reflect.DeepEqual(actual, want) {
		t.Fatalf("returned %+v, want %+v", actual, want)
	}
}
