package bnet

import (
	"github.com/nmccrory/go-bnet/sc2"
	"github.com/nmccrory/go-bnet/wow"
)
// ProfileService has OAuth Profile APIs. See Client.
type ProfileService struct {
	client *Client
}

// SC2Profile represents the profile information for a user's Starcraft 2 profile.
type SC2Profile struct {
	Characters []sc2.SC2Character `json:"characters"`
}

// WoWProfile is a collection of a user's World of Warcraft characters.
type WoWProfile struct {
	Characters []wow.WoWCharacter `json:"characters"`
}

// SC2() calls the /sc2/profile/user endpoint. This endpoint uses OAuth2
// to retrieve a user's Starcraft 2 profile. See Battle.net docs.
func (s *ProfileService) SC2() (*SC2Profile, *Response, error) {
	req, err := s.client.NewRequest("GET", "sc2/profile/user", nil)
	if err != nil {
		return nil, nil, err
	}

	var profile SC2Profile
	resp, err := s.client.Do(req, &profile)
	if err != nil {
		return nil, resp, err
	}

	return &profile, resp, nil
}

// WoW() calls the /wow/user/characters endpoint. This endpoint uses OAuth2
// to retrieve a user's World of Warcraft character list. See Battle.net docs.
func (s *ProfileService) WoW() (*WoWProfile, *Response, error) {
	req, err := s.client.NewRequest("GET", "wow/user/characters", nil)
	if err != nil {
		return nil, nil, err
	}

	var profile WoWProfile
	resp, err := s.client.Do(req, &profile)
	if err != nil {
		return nil, resp, err
	}

	return &profile, resp, nil
}
