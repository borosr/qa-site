package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/borosr/qa-site/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type GithubOAuth struct{}

type GithubUserData struct {
	Login                   string    `json:"login"`
	Id                      int       `json:"id"`
	NodeId                  string    `json:"node_id"`
	AvatarUrl               string    `json:"avatar_url"`
	GravatarId              string    `json:"gravatar_id"`
	Url                     string    `json:"url"`
	HtmlUrl                 string    `json:"html_url"`
	FollowersUrl            string    `json:"followers_url"`
	FollowingUrl            string    `json:"following_url"`
	GistsUrl                string    `json:"gists_url"`
	StarredUrl              string    `json:"starred_url"`
	SubscriptionsUrl        string    `json:"subscriptions_url"`
	OrganizationsUrl        string    `json:"organizations_url"`
	ReposUrl                string    `json:"repos_url"`
	EventsUrl               string    `json:"events_url"`
	ReceivedEventsUrl       string    `json:"received_events_url"`
	Type                    string    `json:"type"`
	SiteAdmin               bool      `json:"site_admin"`
	Name                    string    `json:"name"`
	Company                 string    `json:"company"`
	Blog                    string    `json:"blog"`
	Location                string    `json:"location"`
	Email                   string    `json:"email"`
	Hireable                bool      `json:"hireable"`
	Bio                     string    `json:"bio"`
	TwitterUsername         string    `json:"twitter_username"`
	PublicRepos             int       `json:"public_repos"`
	PublicGists             int       `json:"public_gists"`
	Followers               int       `json:"followers"`
	Following               int       `json:"following"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	PrivateGists            int       `json:"private_gists"`
	TotalPrivateRepos       int       `json:"total_private_repos"`
	OwnedPrivateRepos       int       `json:"owned_private_repos"`
	Collaborators           int       `json:"collaborators"`
	TwoFactorAuthentication bool      `json:"two_factor_authentication"`
}

func (g GithubUserData) Username() string {
	return g.Login
}

func (g GithubUserData) FullName() string {
	return g.Name
}

func (g GithubOAuth) Redirect(w http.ResponseWriter, r *http.Request) {
	config := settings.Get()
	http.Redirect(w, r,
		fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s://%s:%s/login/github/callback",
			config.GithubClientID, "http", config.Hostname, "8081"),
		http.StatusMovedPermanently)
}

func (g GithubOAuth) Available() bool {
	config := settings.Get()
	return !(config.GithubClientID == "" || config.GithubClientSecret == "")
}

func (g GithubOAuth) Callback(w http.ResponseWriter, r *http.Request) (ResponseWrapper, error) {
	code := r.URL.Query().Get("code")

	token, err := g.getGithubAccessToken(code)
	if err != nil {
		return ResponseWrapper{}, err
	}
	var githubData UserDetails
	if githubData, err = g.getGithubData(token.AccessToken); err != nil {
		return ResponseWrapper{}, err
	}

	log.Infof("Github Data: %+v", githubData)

	return ResponseWrapper{Response: token, UserDetails: githubData}, nil
}

func (g GithubOAuth) getGithubAccessToken(code string) (Response, error) {
	config := settings.Get()
	req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token",
		bytes.NewBufferString(`{"client_id":"`+config.GithubClientID+`","client_secret":"`+config.GithubClientSecret+`","code":"`+code+`"}`))
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var oauthResponse Response
	if err := json.Unmarshal(respData, &oauthResponse); err != nil {
		return Response{}, err
	}

	return oauthResponse, nil
}

func (g GithubOAuth) getGithubData(accessToken string) (UserDetails, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userData GithubUserData
	if err := json.Unmarshal(rawResp, &userData); err != nil {
		return nil, err
	}

	return userData, nil
}
