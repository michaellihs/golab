package model

// generated by http://json2struct.mervine.net/
type Project struct {
	ApprovalsBeforeMerge                      int         `json:"approvals_before_merge"`
	Archived                                  bool        `json:"archived"`
	AvatarURL                                 string      `json:"avatar_url"`
	BuildsEnabled                             bool        `json:"builds_enabled"`
	ContainerRegistryEnabled                  bool        `json:"container_registry_enabled"`
	CreatedAt                                 string      `json:"created_at"`
	CreatorID                                 int         `json:"creator_id"`
	DefaultBranch                             string      `json:"default_branch"`
	Description                               string      `json:"description"`
	ForksCount                                int         `json:"forks_count"`
	HTTPURLToRepo                             string      `json:"http_url_to_repo"`
	ID                                        int         `json:"id"`
	IssuesEnabled                             bool        `json:"issues_enabled"`
	LastActivityAt                            string      `json:"last_activity_at"`
	LfsEnabled                                bool        `json:"lfs_enabled"`
	MergeRequestsEnabled                      bool        `json:"merge_requests_enabled"`
	Name                                      string      `json:"name"`
	NameWithNamespace                         string      `json:"name_with_namespace"`
	Namespace                                 struct {
							  ID   int    `json:"id"`
							  Kind string `json:"kind"`
							  Name string `json:"name"`
							  Path string `json:"path"`
						  } `json:"namespace"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool   `json:"only_allow_merge_if_all_discussions_are_resolved"`
	OnlyAllowMergeIfBuildSucceeds             bool   `json:"only_allow_merge_if_build_succeeds"`
	OpenIssuesCount                           int    `json:"open_issues_count"`
	Path                                      string `json:"path"`
	PathWithNamespace                         string `json:"path_with_namespace"`
	Permissions                               struct {
							  GroupAccess   struct {
										AccessLevel       int `json:"access_level"`
										NotificationLevel int `json:"notification_level"`
									} `json:"group_access"`
							  ProjectAccess string `json:"project_access"`
						  } `json:"permissions"`
	Public                                    bool          `json:"public"`
	PublicBuilds                              bool          `json:"public_builds"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled"`
	SharedRunnersEnabled                      bool          `json:"shared_runners_enabled"`
	SharedWithGroups                          []string      `json:"shared_with_groups"`
	SnippetsEnabled                           bool          `json:"snippets_enabled"`
	SSHURLToRepo                              string        `json:"ssh_url_to_repo"`
	StarCount                                 int           `json:"star_count"`
	TagList                                   []string      `json:"tag_list"`
	VisibilityLevel                           int           `json:"visibility_level"`
	WebURL                                    string        `json:"web_url"`
	WikiEnabled                               bool          `json:"wiki_enabled"`
}