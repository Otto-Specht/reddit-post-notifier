package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
)

type API struct {
	AccessToken       string `json:"access_token"`
	AccessTokenExpire int64  `json:"expires_in"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

func GenerateAccessToken(clientID, clientSecret string) API {
	data := []byte(`grant_type=client_credentials`)
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", bytes.NewBuffer(data))
	if err != nil {
		logger.FatalAndExit(err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+basicAuth(clientID, clientSecret))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.FatalAndExit(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.FatalAndExit(fmt.Sprintf("HTTP request failed with status code: %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.FatalAndExit(err.Error())
	}

	var tokenResponse AccessTokenResponse
	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		logger.FatalAndExit(err.Error())
	}

	logger.Debug("Successfully generated API access token")

	return API{
		AccessToken:       tokenResponse.AccessToken,
		AccessTokenExpire: time.Now().Unix() + tokenResponse.ExpiresIn,
	}
}

func basicAuth(clientID, clientSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
}
