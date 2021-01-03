package auth

import "github.com/volatiletech/null/v8"

const (
	DefaultAuthKind AuthKind = "DefaultLogin"
	GithubAuthKind           = "Github"
)

type AuthKind string

type Request struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Token    null.String `json:"token,omitempty"`
}

type Response struct {
	Token       string   `json:"token"`
	RevokeToken string   `json:"revoke_token"`
	AuthKind    AuthKind `json:"auth_kind"`
}
