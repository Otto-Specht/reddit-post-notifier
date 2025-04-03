package discordapi

import "net/http"

var api DiscordApi = DiscordApi{
	httpClient:        http.Client{},
	AccessToken:       "",
	AccessTokenExpire: 0,
}
