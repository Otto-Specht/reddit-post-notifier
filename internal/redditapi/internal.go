package redditapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
)

func generateNewAccessToken() {
	data := []byte(`grant_type=client_credentials`)
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", bytes.NewBuffer(data))
	if err != nil {
		logger.FatalAndExit(err.Error())
	}

	// Loading from env to allow secret rotation and to make life easier
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		logger.FatalAndExit("Missing CLIENT_ID and/or CLIENT_SECRET")
	}

	basicAuth := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+basicAuth)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		logger.FatalAndExit(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.FatalAndExit(fmt.Sprintf("Failed to generate access token. Status code: %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.FatalAndExit(err.Error())
	}

	var tokenResponse AccessTokenResponse
	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		logger.FatalAndExit(err.Error())
	}

	logger.Debug("Successfully generated new API access token")

	api = RedditApi{
		AccessToken:       tokenResponse.AccessToken,
		AccessTokenExpire: time.Now().Unix() + tokenResponse.ExpiresIn,
	}
}

func refreshTokenIfNeeded() {
	if api.AccessToken == "" || time.Now().Unix()+3600 > api.AccessTokenExpire {
		logger.Debug("Generating new access token...")

		generateNewAccessToken()
	}
}

func buildRequest(method string, url string, body io.Reader) *http.Request {
	refreshTokenIfNeeded()

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logger.FatalAndExit(fmt.Sprintf("Failed to create http request. Error: %s", err))
	}

	req.Header.Set("User-Agent", "PostNotifier/1.0")
	req.Header.Set("Authorization", "Bearer "+api.AccessToken)

	return req
}
