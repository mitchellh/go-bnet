# go-bnet

go-bnet is a Go client library for accessing the
[Battle.net API](https://dev.battle.net). In addition to providing an API
client, this package provides OAuth endpoints.

**Documentation:** [![GoDoc](https://godoc.org/github.com/mitchellh/go-bnet?status.svg)](https://godoc.org/github.com/mitchellh/go-bnet)
**Build Status:** [![Build Status](https://travis-ci.org/mitchellh/go-bnet.svg?branch=master)](https://travis-ci.org/mitchellh/go-bnet)

## Usage

```go
import "github.com/mitchellh/go-bnet"
```

### Authentication

Authenticate using the [Go OAuth2](https://golang.org/x/oauth2) library.
Endpoints are provided via the `Endpoint` function. A guide to using OAuth2
to authenticate [is available in this blog post](https://blog.kowalczyk.info/article/f/Accessing-GitHub-API-from-Go.html).
The blog post uses GitHub as an example but it is almost identical for
Battle.net and this library.

Battle.net endpoints are region-specific, so specify the region to the
`Endpoint` function and use the resulting value. Example:

```go
oauthCfg := &oauth2.Config{
    // Get from dev.battle.net
    ClientID:     "",
    ClientSecret: "",

    // Endpoint from this library
    Endpoint: bnet.Endpoint("us"),
}
```

Once you have access to the OAuth client, you can initilize the Battle.net
API client:

```go
// Token from prior auth
authClient := oauthCfg.Client(oauth2.NoContext, token)

// Initialize the client
client := bnet.NewClient(oauthClient)

// ... API calls
```

### API Calls

Once a client is made, basic API calls can easliy be made:

```
user, resp, err := client.Account().User()
fmt.Printf("User: %#v", user)
```

All API calls return a `*Response` value in addition to a richer type
and error. The response contains the http.Response as well as metadata
such as quotas, QPS, etc. from Battle.net
