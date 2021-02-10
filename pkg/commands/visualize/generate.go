package visualize

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

// Generate the table to display.
func GenerateTable(clonedRepoMap map[string][]repos.DetailedLocalRepo, config configure.RegularOutline) table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"OWNER", "NAME", "LAST UPDATED", "CHANGES NOT COMMITTED", "CHANGES NOT PUSHED"})

	for _, clonedRepos := range clonedRepoMap {
		for _, repo := range clonedRepos {
			var (
				commitMsg = color.RedString("Yes")
				pushMsg   = color.RedString("Yes")
			)
			if repo.NotCommitted {
				commitMsg = color.GreenString("No")
			}
			if repo.NotPushed {
				pushMsg = color.GreenString("No")
			}
			t.AppendRow([]interface{}{
				repo.Repo.Owner,
				repo.Repo.Name,
				utils.FormatDate(repo.ModTime),
				commitMsg,
				pushMsg,
			})
		}
		t.AppendSeparator()
	}

	t.SetStyle(table.StyleLight)
	return t
}

// Output a list of all repos with the owner, name, and local path
func OutputOwnerNameList(clonedRepos []repos.LocalRepo) {
	for _, repo := range clonedRepos {
		fmt.Printf("%v/%v\n%v\n", repo.Owner, repo.Name, repo.Path)
	}
}
