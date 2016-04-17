package bnet

import (
	"fmt"
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

// The variables below synchronize access to register an endpoint as
// broken with the OAuth2 library. Battle.net uses a "broken" implementation
// of OAuth2.
var brokenLock sync.Mutex
var brokenMap = map[string]struct{}{}

// Endpoint returns the endpoint for the given region. This doesn't
// validate the region name so you must use one that is valid from Battle.net.
func Endpoint(region string) oauth2.Endpoint {
	region = strings.ToLower(region)
	domain := fmt.Sprintf("https://%s.battle.net/", region)
	if region == "cn" {
		domain = "https://www.battlenet.com.cn/"
	}

	// Register the broken provider
	brokenLock.Lock()
	defer brokenLock.Unlock()
	if _, ok := brokenMap[domain]; !ok {
		brokenMap[domain] = struct{}{}
		oauth2.RegisterBrokenAuthHeaderProvider(domain)
	}

	// Build the endpoint
	return oauth2.Endpoint{
		AuthURL:  domain + "oauth/authorize",
		TokenURL: domain + "oauth/token",
	}
}
