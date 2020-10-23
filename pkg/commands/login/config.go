package login

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var cfg = oauth2.Config{
	ClientID:     "96e23ef00d878a9a557c",
	ClientSecret: "6bd3ef059485fbabc8bc9d23668a7ac6263f025c",
	Endpoint:     github.Endpoint,
}
