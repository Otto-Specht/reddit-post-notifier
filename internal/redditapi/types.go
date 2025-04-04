package redditapi

import "net/http"

type RedditApi struct {
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
	TotalKarma int64  `json:"total_karma"`
}

type UserSubmittedFeed struct {
	Entries []UserSubmittedEntry `xml:"entry"`
}

type UserSubmittedEntry struct {
	Id        string      `xml:"id"`
	Title     string      `xml:"title"`
	Published string      `xml:"published"`
	Link      Link        `xml:"link"`
	Author    EntryAuthor `xml:"author"`
	Category  Category    `xml:"category"`
}

type EntryAuthor struct {
	Name string `xml:"name"`
	Uri  string `xml:"uri"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Category struct {
	Term  string `xml:"term,attr"`
	Label string `xml:"label,attr"`
}
