package main

import "github.com/nicklaw5/helix"

func login() (*helix.Client, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return nil, err
	}
	resp, err := client.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		return nil, err
	}
	client.SetAppAccessToken(resp.Data.AccessToken)
	return client, nil
}
