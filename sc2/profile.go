package sc2

import (
	"fmt"
	"github.com/nmccrory/go-bnet"
	profile "github.com/nmccrory/go-bnet/profile"
)

type ProfileService struct {
	client *Client
}

func (s *ProfileService) Profile(id int, region int, name string) (*profile.SC2Profile, *bnet.Response, error) {
	url := fmt.Sprintf("profile/%d/%d/%s", id, region, name)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var profile profile.SC2Profile
	resp, err := s.client.Do(req, &profile)
	if err != nil {
		return nil, resp, err
	}

	return &profile, resp, nil
}