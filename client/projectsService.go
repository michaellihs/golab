package client

import (
	"github.com/michaellihs/golab/model"
)

type ProjectsService struct {
	Client *GitlabClient
}

func (service *ProjectsService) List() *[]model.Project {
	projects := new([]model.Project)
	req, _ := service.Client.NewGetRequest("/api/v3/projects")
	service.Client.Do(req, projects)
	return projects
}
