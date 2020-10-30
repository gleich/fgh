package api

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Generate a githubv4 client
func GenerateClient(rawToken string) *githubv4.Client {
	token := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: rawToken},
	)
	httpClient := oauth2.NewClient(context.Background(), token)
	return githubv4.NewClient(httpClient)
}
