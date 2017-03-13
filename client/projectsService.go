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
	"github.com/michaellihs/golab/model"
	"fmt"
	"strings"
	"net/http"
	"errors"
)

type ProjectsService struct {
	Client *GitlabClient
}

type ProjectParams struct {
	Name                                      string `json:"name,omitempty"`
	Path                                      string `json:"path,omitempty"`
	NamespaceId                               int    `json:"namespace_id,omitempty"`
	DefaultBranch                             string `json:"default_branch,omitempty"`
	Description                               string `json:"description,omitempty"`
	IssuesEnabled                             bool   `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled                      bool   `json:"merge_requests_enabled,omitempty"`
	BuildsEnabled                             bool   `json:"builds_enabled,omitempty"`
	WikiEnabled                               bool   `json:"wiki_enabled,omitempty"`
	SnippetsEnabled                           bool   `json:"snippets_enabled,omitempty"`
	ContainerRegistryEnabled                  bool   `json:"container_registry_enabled,omitempty"`
	SharedRunnersEnabled                      bool   `json:"shared_runners_enabled,omitempty"`
	Visibility                                string `json:"visibility,omitempty"`
	ImportUrl                                 string `json:"import_url,omitempty"`
	PublicBuilds                              bool   `json:"public_builds,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          bool   `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool   `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	LfsEnabled                                bool   `json:"lfs_enabled,omitempty"`
	RequestAccessEnabled                      bool   `json:"request_access_enabled,omitempty"`
	RepositoryStorage                         string `json:"repository_storage,omitempty"`
	ApprovalsBeforeMerge                      int    `json:"approvals_before_merge,omitempty"`
}

func (service *ProjectsService) Get(projectId string) (*model.Project, error) {
	encodedProjectId := strings.Replace(projectId, "/", "%2F", -1)
	project := new(model.Project)
	req, err1 := service.Client.NewGetRequest("/projects/" + encodedProjectId)
	if err1 != nil {
		return nil, err1
	}
	_, err2 := service.Client.Do(req, project)
	if err2 != nil {
		return nil, err2
	}
	return project, nil
}

func (service *ProjectsService) List() (*[]model.Project, error) {
	projects := new([]model.Project)
	req, err := service.Client.NewGetRequest("/projects")
	if err != nil {
		return nil, err
	}
	_, err = service.Client.Do(req, projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (service *ProjectsService) Create(projectParams *ProjectParams) (*model.Project, error) {
	req, _ := service.Client.NewPostRequest("/projects", projectParams)
	project := new(model.Project)
	_, err := service.Client.Do(req, project)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
		return nil, err
	}
	return project, nil
}

func (service *ProjectsService) Delete(projectId string) (bool, error) {
	encodedProjectId := strings.Replace(projectId, "/", "%2F", -1)
	req, _ := service.Client.NewDeleteRequest("/projects/" + encodedProjectId)
	resp, err := service.Client.Do(req, nil)
	if err != nil {
		fmt.Println("An error occured: " + err.Error())
		return false, err
	}
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, errors.New(resp.Status)
	}
}
