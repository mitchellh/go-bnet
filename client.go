package bnet

import (
	"bytes"
	"encoding/json"
	"io"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
)


const (
	libraryVersion = "0.1"
	userAgent      = "go-bnet/" + libraryVersion
)

// Client is the API client for Battle.net.
// Create this using one of the Battle.net libary's NewClient functions.
// This can also be constructed manually but it isn't recommended.
type Client struct {
	// Client is the HTTP client to use for communication.
	Client *http.Client

	// BaseURL is the base URL for API requests. This should match
	// the region with the auth region used for Client.
	BaseURL *url.URL

	// UserAgent is the user agent to set on API requests.
	UserAgent string
}

// NewClient creates a new Battle.net client.
//
// region must be a valid Battle.net region. This will not validate it
// is valid.
//
// The http.Client argument should usually be retrieved via the
// oauth2 Go library NewClient function. It must be a client that
// automatically injects authentication details into requests.
func NewClient(region string, c *http.Client) *Client {
	region = strings.ToLower(region)

	if c == nil {
		c = http.DefaultClient
	}

	// Determine the API base URL based on the region
	baseURLStr := fmt.Sprintf("https://%s.api.battle.net/", region)
	if region == "cn" {
		baseURLStr = "https://api.battlenet.com.cn/"
	}

	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		// We panic because we manually construct it above so it should
		// never really fail unless the user gives us a REALLY bad region.
		panic(err)
	}
	return &Client{
		Client:    c,

		BaseURL:   baseURL,

		UserAgent: userAgent,
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	if err := CheckError(resp); err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}
