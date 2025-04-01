package api

import "net/http"

type API struct {
	httpClient        http.Client
	AccessToken       string
	AccessTokenExpire int64
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

type UserAbout struct {
	Data UserAboutData `json:"data"`
}

type UserAboutData struct {
	Name       string `json:"name"`
	TotalKarma string `json:"total_karma"`
}
