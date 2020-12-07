package tree

import "github.com/Matt-Gleich/fgh/pkg/repos"

// Get the repos for each user. User mapped to repo.
func GetRepos(clonedRepos []repos.LocalRepo) map[string][]repos.DetailedLocalRepo {
	mappedRepos := map[string][]repos.DetailedLocalRepo{}

	for _, repo := range clonedRepos {
		var (
			notCommitted, notPushed = repos.WorkingState(repo.Path)
			updatedTime             = repos.LastUpdated(repo.Path)
		)
		mappedRepos[repo.Owner] = append(mappedRepos[repo.Owner], repos.DetailedLocalRepo{
			Repo:        repo,
			ModTime:     updatedTime,
			Uncommitted: notCommitted,
			NotPushed:   notPushed,
		})
	}

	return mappedRepos
}
