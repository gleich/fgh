package api

import (
	"context"

	"github.com/shurcooL/githubv4"
)

// Get data about a repo
func RepoData(client *githubv4.Client, owner string, name string) (Repo, error) {
	query := struct {
		Repository struct {
			IsPrivate  bool
			IsTemplate bool
			IsMirror   bool
			IsFork     bool
			IsArchived bool
			IsDisabled bool
			Languages  struct {
				Nodes []struct {
					Name string
				}
			} `graphql:"languages(first: 1, orderBy: {field: SIZE, direction: DESC})"`
			Name  string
			Owner struct {
				User struct {
					Login string
				} `graphql:"... on User"`
			}
		} `graphql:"repository(owner: $owner, name: $name)"`
	}{}

	vars := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}
	err := client.Query(context.Background(), &query, vars)
	if err != nil {
		return Repo{}, err
	}

	language := "Other"
	if len(query.Repository.Languages.Nodes) != 0 {
		language = query.Repository.Languages.Nodes[0].Name
	}
	return Repo{
		Owner:        query.Repository.Owner.User.Login,
		Name:         query.Repository.Name,
		MainLanguage: language,
		Private:      query.Repository.IsPrivate,
		Archived:     query.Repository.IsArchived,
		Template:     query.Repository.IsTemplate,
		Disabled:     query.Repository.IsDisabled,
		Mirror:       query.Repository.IsMirror,
		Fork:         query.Repository.IsFork,
	}, nil
}
