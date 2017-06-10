package bnet

// ProfileService has OAuth Profile APIs. See Client.
type ProfileService struct {
	client *Client
}

// SC2Profile represents the profile information for a user's Starcraft 2 profile.
type SC2Profile struct {
	Characters []SC2Character `json:"characters"`
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
