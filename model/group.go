package model

type Group struct {
	AvatarURL            interface{} `json:"avatar_url"`
	Description          string      `json:"description"`
	FullName             string      `json:"full_name"`
	FullPath             string      `json:"full_path"`
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	ParentID             interface{} `json:"parent_id"`
	Path                 string      `json:"path"`
	Projects             []Project   `json:"projects"`
	RequestAccessEnabled bool        `json:"request_access_enabled"`
	SharedProjects       []Project   `json:"shared_projects"`
	Visibility           string      `json:"visibility"`
	WebURL               string      `json:"web_url"`
}
