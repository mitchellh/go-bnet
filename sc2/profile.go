package sc2

import (
	"fmt"
	"github.com/nmccrory/go-bnet"
)

// Profile() calls the /sc2/profile/:id/:region/:name endpoint to
// retrieve a user's Starcraft 2 profile. See Battle.net docs.
func (s *SC2Service) Profile(id int, region int, name string) (*bnet.SC2Profile, *bnet.Response, error) {
	endpoint := fmt.Sprintf("sc2/profile/%d/%d/%s", id, region, name)
	req, err := s.client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var profile bnet.SC2Profile
	resp, err := s.client.Do(req, &profile)
	if err != nil {
		return nil, resp, err
	}

	return &profile, resp, nil
}

