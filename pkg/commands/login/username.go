package login

import (
	"context"

	"github.com/gleich/fgh/pkg/api"
	"github.com/gleich/fgh/pkg/utils"
)

// Get the username of the authed user
func Username(token string) (string, utils.CtxErr) {
	query := struct {
		Viewer struct {
			Login string
		}
	}{}

	client := api.GenerateClient(token)
	err := client.Query(context.Background(), &query, nil)
	if err != nil || query.Viewer.Login == "" {
		return "", utils.CtxErr{
			Context: "Failed to yet your username",
			Error:   err,
		}
	}
	return query.Viewer.Login, utils.CtxErr{}
}
