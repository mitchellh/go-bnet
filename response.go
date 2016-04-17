package bnet

import (
	"net/http"
	"strconv"
	"time"
)

// Reponse is a Battle.net API response. This wraps the standard http.Response
// and provides convenient access to some of the metadata returned.
type Response struct {
	*http.Response

	QPSCurrent    int
	QPSAllotted   int
	QuotaCurrent  int
	QuotaAllotted int
	QuotaReset    time.Time
}

func newResponse(r *http.Response) *Response {
	result := &Response{Response: r}
	result.parseMeta()
	return result
}

func (r *Response) parseMeta() {
	// Parse the basic ints
	intMaps := map[string]*int{
		"X-Plan-Qps-Allotted":   &r.QPSAllotted,
		"X-Plan-Qps-Current":    &r.QPSCurrent,
		"X-Plan-Quota-Allotted": &r.QuotaAllotted,
		"X-Plan-Quota-Current":  &r.QuotaCurrent,
	}

	for k, ptr := range intMaps {
		if v := r.Response.Header.Get(k); v != "" {
			*ptr, _ = strconv.Atoi(v)
		}
	}

	// Parse the reset time
	if v := r.Response.Header.Get("X-Plan-Quota-Reset"); v != "" {
		r.QuotaReset, _ = time.Parse("Monday, January 2, 2006 3:04:05 PM MST", v)
	}
}
