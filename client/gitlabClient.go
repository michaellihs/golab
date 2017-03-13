// Copyright © 2017 Michael Lihs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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

