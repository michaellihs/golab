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

func NewClient(baseUrl *url.URL, token string, httpClient *http.Client) *GitlabClient {
	base := sling.New().Client(httpClient).Base(baseUrl.String())
	gitlabClient := &GitlabClient{sling: base, token: token, client: httpClient}

	gitlabClient.Projects = &ProjectsService{Client:gitlabClient}

	return gitlabClient
}

func (client *GitlabClient) NewGetRequest(url string) (*http.Request, error) {
	req, err := client.sling.New().Get(url).Request()
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", client.token)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (client *GitlabClient) NewPostRequest(url string, body interface{}) (*http.Request, error) {
	req, err := client.sling.New().Post(url).BodyJSON(body).Request()
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", client.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (client *GitlabClient) Do(req *http.Request, value interface{}) (*http.Response, error){
	resp, err := client.sling.Do(req, value, "")
	return resp, err
}

