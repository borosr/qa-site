package oauth

import (
	"fmt"
	"net/http"

	"github.com/borosr/qa-site/pkg/settings"
)

const (
	GithubProvider providerType = "github"
)

var (
	config = settings.Get()

	providers = map[providerType]Provider{
		GithubProvider: &GithubOAuth{},
	}

	configs = map[providerType]Config{
		GithubProvider: {
			ClientID:     config.GithubClientID,
			ClientSecret: config.GithubClientSecret,
		},
	}
)

type providerType string

type Config struct {
	ClientID     string
	ClientSecret string
}

type Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type ResponseWrapper struct {
	Response    Response
	UserDetails UserDetails
}

type Provider interface {
	Redirect(http.ResponseWriter, *http.Request)
	Callback(http.ResponseWriter, *http.Request) (ResponseWrapper, error)
	Available() bool
}

type UserDetails interface {
	Username() string
	FullName() string
	// TODO add avatarURL later and more needed getters
}

func Redirect(w http.ResponseWriter, r *http.Request, provider string) error {
	if p, ok := providers[providerType(provider)]; ok {
		p.Redirect(w, r)
		return nil
	}

	return fmt.Errorf("invalid provider type [%s]", provider)
}

func Callback(w http.ResponseWriter, r *http.Request, provider string) (ResponseWrapper, error) {
	if p, ok := providers[providerType(provider)]; ok {
		return p.Callback(w, r)
	}
	return ResponseWrapper{}, fmt.Errorf("invalid provider type [%s]", provider)
}

func Availability() map[providerType]bool {
	returnMap := make(map[providerType]bool)
	for key, element := range providers {
		returnMap[key] = element.Available()
	}
	return returnMap
}
