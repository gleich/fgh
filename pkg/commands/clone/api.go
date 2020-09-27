package clone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"golang.org/x/oauth2"
)

type Repository struct {
	Owner        string
	Name         string
	MainLanguage string
	Private      bool
	Archived     bool
	Template     bool
	Disabled     bool
	Mirror       bool
	Fork         bool
}

func GetRepository(secrets configure.SecretsOutline, args []string) Repository {
	query, owner, name := fillQuery(secrets, args)
	spin := spinner.New(spinner.CharSets[13], 40*time.Millisecond)
	spin.Suffix = fmt.Sprintf(" ℹ️  Getting metadata for %v/%v", owner, name)
	spin.Start()
	repo := getData(secrets, query, owner, name)
	spin.Stop()
	statuser.Success(fmt.Sprintf("Got metadata for %v/%v", owner, name))
	return repo
}

// Fill in the variables to the query
func fillQuery(secrets configure.SecretsOutline, args []string) (string, string, string) {
	query := `
	query Repository {
		repository(name: "$name", owner: "$owner") {
			isPrivate
			isTemplate
			isMirror
			isFork
			isArchived
			isDisabled
			languages(first: 1, orderBy: {field: SIZE, direction: DESC}) {
				nodes {
					name
				}
			}
		}
	}`

	// Getting repo name and owner
	var (
		owner string
		name  string
	)
	if strings.Contains(args[0], "/") {
		parts := strings.Split(args[0], "/")
		owner = parts[0]
		name = parts[1]
	} else {
		owner = secrets.Username
		name = args[0]
	}

	// Filling in variables
	query = strings.ReplaceAll(query, "$name", name)
	query = strings.ReplaceAll(query, "$owner", owner)

	return query, owner, name
}

// Make the actual request
func getData(secrets configure.SecretsOutline, query string, owner string, name string) Repository {
	// Creating a http client with the PAT (Personal Access Token)
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: secrets.PAT},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Creating the request
	jsonValue, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		statuser.Error("Failed to create JSON for GitHub GraphQL API request", err, 1)
	}
	request, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		statuser.Error("Failed to formulate request to GitHub GraphQL API", err, 1)
	}

	// Making the request
	response, err := httpClient.Do(request)
	if err != nil || response.StatusCode != 200 {
		statuser.Error("Failed to get data from GitHub GraphQL API", err, 1)
	}
	defer response.Body.Close()

	// Getting the actual data from the response
	binaryData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		statuser.Error("Failed to binary data from response", err, 1)
	}
	var data struct {
		Data struct {
			Repository struct {
				IsPrivate  bool `json:"isPrivate"`
				IsTemplate bool `json:"isTemplate"`
				IsMirror   bool `json:"isMirror"`
				IsFork     bool `json:"isFork"`
				IsArchived bool `json:"isArchived"`
				IsDisabled bool `json:"isDisabled"`
				Languages  struct {
					Nodes []struct {
						Name string `json:"name"`
					} `json:"nodes"`
				} `json:"languages"`
			} `json:"repository"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	err = json.Unmarshal(binaryData, &data)
	if err != nil {
		statuser.Error("Failed to fill binary into initial data struct", err, 1)
	}
	if len(data.Errors) != 0 {
		statuser.ErrorMsg("Failed to get repo data. "+data.Errors[0].Message, 1)
	}

	// Mapping data to Repository struct
	return Repository{
		Owner:        owner,
		Name:         name,
		MainLanguage: data.Data.Repository.Languages.Nodes[0].Name,
		Private:      data.Data.Repository.IsPrivate,
		Archived:     data.Data.Repository.IsArchived,
		Template:     data.Data.Repository.IsTemplate,
		Disabled:     data.Data.Repository.IsDisabled,
		Mirror:       data.Data.Repository.IsMirror,
		Fork:         data.Data.Repository.IsFork,
	}
}
