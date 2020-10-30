package login

import (
	"context"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/statuser/v2"
)

// Get the username of the authed user
func Username(token string) string {
	query := struct {
		Viewer struct {
			Login string
		}
	}{}

	client := api.GenerateClient(token)
	err := client.Query(context.Background(), &query, nil)
	if err != nil || query.Viewer.Login == "" {
		statuser.Error("Failed to yet your username", err, 1)
	}
	return query.Viewer.Login
}
