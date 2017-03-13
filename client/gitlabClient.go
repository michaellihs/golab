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
	Groups   *GroupsService
}

const (
	apiPath string = "/api/v3"
)

func NewClient(baseUrl *url.URL, token string, httpClient *http.Client) *GitlabClient {
	base := sling.New().Client(httpClient).Base(baseUrl.String())
	gitlabClient := &GitlabClient{sling: base, token: token, client: httpClient}

	gitlabClient.Projects = &ProjectsService{Client:gitlabClient}
	gitlabClient.Groups = &GroupsService{Client:gitlabClient}

	return gitlabClient
}

func (client *GitlabClient) NewGetRequest(url string) (*http.Request, error) {
	req, err := client.sling.New().Get(apiPath + url).Request()
	if err != nil {
		return nil, err
	}
	client.setHeaders(req)
	return req, nil
}

func (client *GitlabClient) NewPostRequest(url string, body interface{}) (*http.Request, error) {
	req, err := client.sling.New().Post(apiPath + url).BodyJSON(body).Request()
	if err != nil {
		return nil, err
	}
	client.setHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (client *GitlabClient) NewDeleteRequest(url string) (*http.Request, error) {
	req, err := client.sling.New().Delete(apiPath + url).Request()
	if err != nil {
		return nil, err
	}
	client.setHeaders(req)
	return req, nil
}

func (client *GitlabClient) Do(req *http.Request, value interface{}) (*http.Response, error) {
	resp, err := client.sling.Do(req, value, nil)
	return resp, err
}

func (client *GitlabClient) setHeaders(req *http.Request) (*http.Request) {
	req.Header.Set("PRIVATE-TOKEN", client.token)
	req.Header.Set("Accept", "application/json")
	return req
}

