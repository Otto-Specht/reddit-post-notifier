package api

import (
	"encoding/json"
	"encoding/xml"
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

func GetLatestPostId(user string) string {
	req := buildRequest("GET", "https://oauth.reddit.com/user/"+user+"/submitted.rss?limit=1", nil)
	resp, err := api.httpClient.Do(req)
	if err != nil {
		logger.FatalAndExit(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to get latest post id for user '%s', Error reading response body: %s.", user, err))
			return ""
		}

		var feed UserSubmittedFeed
		err = xml.Unmarshal(body, &feed)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to get latest post id for user '%s', Error parsing XML: %s.", user, err))
			return ""
		}

		if len(feed.Entries) == 0 {
			logger.Warn(fmt.Sprintf("User '%s' has no posts yet, the script will still work", user))
			return ""
		}

		lastPost := feed.Entries[0]

		logger.Debug(fmt.Sprintf("Lastest post from user u/%s: (%s) %s", user, lastPost.Published, lastPost.Title))

		return lastPost.Id
	} else {
		logger.Error(fmt.Sprintf("Failed to get latest post id for user '%s', got status %s.", user, resp.Status))
		return ""
	}
}
