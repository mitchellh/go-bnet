package bnet

import (
	//"fmt"
	//"net/http"
	//"net/url"
)

type Profile struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func (s *AccountService) Profile(id int) (*Profile, *Response, error) {
	req, err := s.client.NewRequest("GET", "sc2/profile/user", nil)
	if err != nil {
		return nil, nil, err
	}

	var profile Profile
	resp, err := s.client.Do(req, &profile)
	if err != nil {
		return nil, resp, err
	}

	return &profile, resp, nil
}