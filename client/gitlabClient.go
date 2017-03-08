package client

import (
	"github.com/dghubble/sling"
	"net/http"
	"net/url"
)

type GitlabClient struct {
	sling    *sling.Sling
	token    string
	client   *http.Client
	baseUrl  *url.URL

	Projects *ProjectsService
}

func NewClient(gitlabUrl string, token string, httpClient *http.Client) *GitlabClient {
	// TODO check gitlabUrl for being a proper URL
	// TODO use URL type for URL instead of string
	base := sling.New().Client(httpClient).Base(gitlabUrl)
	gitlabClient := &GitlabClient{
		sling: base,
		token: token,
		client: httpClient}

	gitlabClient.Projects = &ProjectsService{Client:gitlabClient}

	return gitlabClient
}

func (client *GitlabClient) NewGetRequest(url string) (*http.Request, error) {
	req, err := client.sling.New().Get(url).Request()
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", client.token)
	return req, nil
}

func (client *GitlabClient) Do(req *http.Request, value interface{}) {
	error := new(string)
	client.sling.Do(req, value, error)
}

