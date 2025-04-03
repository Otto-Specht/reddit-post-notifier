package discordapi

import "net/http"

type DiscordApi struct {
	httpClient        http.Client
	AccessToken       string
	AccessTokenExpire int64
}
