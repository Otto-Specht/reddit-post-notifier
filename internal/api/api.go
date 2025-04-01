package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
)

var api API = API{
	httpClient:        http.Client{},
	AccessToken:       "",
	AccessTokenExpire: 0,
}

func CheckIfUsersExistOrRemove(userList []string) []string {
	refreshTokenIfNeeded()

	existingUserList := []string{}

	for _, value := range userList {
		req := buildRequest("GET", "https://oauth.reddit.com/user/"+value+"/about.json", nil)
		resp, err := api.httpClient.Do(req)
		if err != nil {
			logger.FatalAndExit(err.Error())
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.FatalAndExit(err.Error())
			}

			var userAboutResponse UserAbout
			if err = json.Unmarshal(body, &userAboutResponse); err != nil {
				logger.FatalAndExit(err.Error())
			}

			logger.Info(fmt.Sprintf("Adding user u/%s (Karma: %v)", userAboutResponse.Data.Name, userAboutResponse.Data.TotalKarma))

			existingUserList = append(existingUserList, userAboutResponse.Data.Name)
		} else {
			logger.Warn(fmt.Sprintf("Cannot find user with name '%s', got status %s.", value, resp.Status))
		}
	}

	return existingUserList
}
