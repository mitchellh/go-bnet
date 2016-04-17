package bnet

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Plan-Qps-Allotted", "100")
		w.Header().Set("X-Plan-Qps-Current", "5")
		w.Header().Set("X-Plan-Quota-Allotted", "36000")
		w.Header().Set("X-Plan-Quota-Current", "32")
		w.Header().Set("X-Plan-Quota-Reset", "Sunday, April 17, 2016 7:00:00 PM GMT")
	})

	req, err := client.NewRequest("GET", "test", nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	actual, err := client.Do(req, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	actual.Response = nil

	gmt, err := time.LoadLocation("GMT")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	want := &Response{
		QPSCurrent:    5,
		QPSAllotted:   100,
		QuotaCurrent:  32,
		QuotaAllotted: 36000,
		QuotaReset:    time.Date(2016, time.April, 17, 19, 0, 0, 0, gmt),
	}

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("returned %+v, want %+v", actual, want)
	}
}
