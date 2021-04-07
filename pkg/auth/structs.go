package auth

import "github.com/volatiletech/null/v8"

const (
	DefaultAuthKind Kind = "DefaultLogin"
	GithubAuthKind       = "Github"
)

type Kind string

type Request struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Token    null.String `json:"token,omitempty"`
}

type Response struct {
	Token       string `json:"token"`
	RevokeToken string `json:"revoke_token"`
	AuthKind    Kind   `json:"auth_kind"`
}
