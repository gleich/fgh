package api

import (
	"context"
	"fmt"

	"github.com/gleich/fgh/pkg/utils"
	"github.com/shurcooL/githubv4"
)

// Get data about a repo
func RepoData(client *githubv4.Client, owner string, name string) (Repo, utils.CtxErr) {
	query := struct {
		Repository struct {
			IsPrivate       bool
			IsTemplate      bool
			IsMirror        bool
			IsFork          bool
			IsArchived      bool
			IsDisabled      bool
			PrimaryLanguage struct {
				Name string
			}
			Name  string
			Owner struct {
				Login string
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}{}

	vars := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}
	err := client.Query(context.Background(), &query, vars)
	if err != nil {
		return Repo{}, utils.CtxErr{
			Context: fmt.Sprintf("Failed to get repo data for %v/%v", owner, name),
			Error:   err,
		}
	}

	if query.Repository.PrimaryLanguage.Name == "" {
		query.Repository.PrimaryLanguage.Name = "Other"
	}
	return Repo{
		Owner:        query.Repository.Owner.Login,
		Name:         query.Repository.Name,
		MainLanguage: query.Repository.PrimaryLanguage.Name,
		Private:      query.Repository.IsPrivate,
		Archived:     query.Repository.IsArchived,
		Template:     query.Repository.IsTemplate,
		Disabled:     query.Repository.IsDisabled,
		Mirror:       query.Repository.IsMirror,
		Fork:         query.Repository.IsFork,
	}, utils.CtxErr{}
}
