package client

import (
	"github.com/dghubble/sling"
	"net/http"
	"github.com/michaellihs/golab/model"
	"net/url"
	"github.com/michaellihs/golab/client/services"
)

type GitlabClient struct {
	sling   *sling.Sling
	token   string
	client  *http.Client
	baseUrl *url.URL

	ProjectService services.ProjectsService
}

func NewClient(gitlabUrl string, token string, httpClient *http.Client) *GitlabClient {
	base := sling.New().Client(httpClient).Base(gitlabUrl)
	return &GitlabClient{
		sling: base,
		token: token,
		client: httpClient}
}

func (client *GitlabClient) NewGetRequest(url string) (*http.Request, error) {
	req, err := client.sling.New().Get(url).Request()
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", client.token)
	return req, nil
}

func (client GitlabClient) ListProjects() *[]model.Project {
	projects := new([]model.Project)
	error := new(string)
	req, _ := client.NewGetRequest("/api/v3/projects")
	client.sling.Do(req, projects, error)
	return projects
}

