package client

import (
	"github.com/dghubble/sling"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/michaellihs/golab/model"
)

type GitlabClient struct {
	sling  *sling.Sling
	token  string
	client *http.Client
}

func NewClient(gitlabUrl string, token string, httpClient *http.Client) *GitlabClient {
	base := sling.New().Client(httpClient).Base(gitlabUrl)
	return &GitlabClient{
		sling: base,
		token: token,
		client: httpClient}
}

func (client GitlabClient) ListProjects() *[]model.Project {
	req, _ := client.sling.New().Get("/api/v3/projects?private_token=" + client.token).Request()
	fmt.Println(req.URL)
	resp, _ := client.client.Do(req)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)

	projects := new([]model.Project)
	client.sling.New().Get("/api/v3/projects?private_token=" + client.token).ReceiveSuccess(projects)
	return projects
}

