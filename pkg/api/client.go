package api

import (
	"context"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Generate a githubv4 client
func GenerateClient() *githubv4.Client {
	token := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: configuration.GetSecrets().PAT},
	)
	httpClient := oauth2.NewClient(context.Background(), token)
	return githubv4.NewClient(httpClient)
}
