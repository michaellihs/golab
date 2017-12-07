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

package cmd

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"github.com/michaellihs/golab/cmd/mapper"
)

var createOptsMapper, listOptsMapper, getOptsMapper, editOptsMapper, forkOptsMapper, listForksOptsMapper mapper.FlagMapper

var id int

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  `List, create, edit and delete projects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, _, err := gitlabClient.Projects.ListProjects(&gitlab.ListProjectsOptions{})
		if err != nil {
			return err
		}
		err = OutputJson(projects)
		return err
	},
}

type listFlags struct {
	Archived                 *bool   `flag_name:"archived" type:"bool" required:"no" description:"Limit by archived status"`
	Visibility               *string `flag_name:"visibility" type:"string" required:"no" description:"Limit by visibility public, internal, or private"`
	OrderBy                  *string `flag_name:"order_by" type:"string" required:"no" description:"Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at"`
	Sort                     *string `flag_name:"sort" type:"string" required:"no" description:"Return projects sorted in asc or desc order. Default is desc"`
	Search                   *string `flag_name:"search" type:"string" required:"no" description:"Return list of projects matching the search criteria"`
	Simple                   *bool   `flag_name:"simple" type:"bool" required:"no" description:"Return only the ID, URL, name, and path of each project"`
	Owned                    *bool   `flag_name:"owned" type:"bool" required:"no" description:"Limit by projects owned by the current user"`
	Membership               *bool   `flag_name:"membership" type:"bool" required:"no" description:"Limit by projects that the current user is a member of"`
	Starred                  *bool   `flag_name:"starred" type:"bool" required:"no" description:"Limit by projects starred by the current user"`
	Statistics               *bool   `flag_name:"statistics" type:"bool" required:"no" description:"Include project statistics"`
	WithIssuesEnabled        *bool   `flag_name:"with_issues_enabled" type:"bool" required:"no" description:"Limit by enabled issues feature"`
	WithMergeRequestsEnabled *bool   `flag_name:"with_merge_requests_enabled" type:"bool" required:"no" description:"Limit by enabled merge requests feature"`
	// TODO custom attributes are currently not supported
}

var projectListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all projects",
	Long:  `Get a list of all visible projects across GitLab for the authenticated user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := createListOpts()
		projects, _, err := gitlabClient.Projects.ListProjects(opts)
		if err != nil {
			return err
		}
		return OutputJson(projects)
	},
}

func createListOpts() (*gitlab.ListProjectsOptions, error) {
	opts := &gitlab.ListProjectsOptions{}
	flags := &listFlags{}
	listOptsMapper.Map(flags, opts)
	return opts, nil
}

type getFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"either the project ID (numeric) or 'namespace/project-name'"`
	// TODO currently not supported by go-gitlab
	Statistics *bool `flag_name:"statistics" short:"s" required:"no" description:"include project statistics"`
}

var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get detailed information for a project",
	Long:  `Get detailed information for a project identified by either project ID or 'namespace/project-name'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := initProjectGetOpts()
		if *opts.Id == "" {
			return errors.New("you have to provide a project ID or 'namespace/project-name' with the -i --id flag")
		}
		project, _, err := gitlabClient.Projects.GetProject(parsePid(*opts.Id)) // make sure, parsedPid is of type int if numeric
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

func initProjectGetOpts() (*getFlags, error) {
	flags := &getFlags{}
	opts := &getFlags{}
	getOptsMapper.Map(flags, opts)
	return opts, nil
}

type createOpts struct {
	Name                                      *string   `flag_name:"name" type:"string" required:"yes" description:"The name of the new project"`
	Path                                      *string   `flag_name:"path" type:"string" required:"no" description:"Custom repository name for new project.By default generated based on name"`
	DefaultBranch                             *string   `flag_name:"default_branch" type:"string" required:"no" description:"master by default"`
	NamespaceID                               *int      `flag_name:"namespace_id" type:"integer" required:"no" description:"Namespace ID (Group ID) for the new project (defaults to the current user's namespace)"`
	Description                               *string   `flag_name:"description" type:"string" required:"no" description:"Short project description"`
	IssuesEnabled                             *bool     `flag_name:"issues_enabled" type:"bool" required:"no" description:"Enable issues for this project"`
	MergeRequestsEnabled                      *bool     `flag_name:"merge_requests_enabled" type:"bool" required:"no" description:"Enable merge requests for this project"`
	JobsEnabled                               *bool     `flag_name:"jobs_enabled" type:"bool" required:"no" description:"Enable jobs for this project"`
	WikiEnabled                               *bool     `flag_name:"wiki_enabled" type:"bool" required:"no" description:"Enable wiki for this project"`
	SnippetsEnabled                           *bool     `flag_name:"snippets_enabled" type:"bool" required:"no" description:"Enable snippets for this project"`
	ResolveOutdatedDiffDiscussions            *bool     `flag_name:"resolve_outdated_diff_discussions" type:"bool" required:"no" description:"Automatically resolve merge request diffs discussions on lines changed with a push"`
	ContainerRegistryEnabled                  *bool     `flag_name:"container_registry_enabled" type:"bool" required:"no" description:"Enable container registry for this project"`
	SharedRunnersEnabled                      *bool     `flag_name:"shared_runners_enabled" type:"bool" required:"no" description:"Enable shared runners for this project"`
	Visibility                                *string   `flag_name:"visibility" type:"string" required:"no" description:"See project visibility level"`
	ImportUrl                                 *string   `flag_name:"import_url" type:"string" required:"no" description:"URL to import repository from"`
	PublicJobs                                *bool     `flag_name:"public_jobs" type:"bool" required:"no" description:"If true, jobs can be viewed by non-project-members"`
	OnlyAllowMergeIfPipelineSucceeds          *bool     `flag_name:"only_allow_merge_if_pipeline_succeeds" type:"bool" required:"no" description:"Set whether merge requests can only be merged with successful jobs"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool     `flag_name:"only_allow_merge_if_all_discussions_are_resolved" type:"bool" required:"no" description:"Set whether merge requests can only be merged when all the discussions are resolved"`
	LfsEnabled                                *bool     `flag_name:"lfs_enabled" type:"bool" required:"no" description:"Enable LFS"`
	RequestAccessEnabled                      *bool     `flag_name:"request_access_enabled" type:"bool" required:"no" description:"Allow users to request member access"`
	TagList                                   *[]string `flag_name:"tag_list" type:"array" required:"no" description:"The list of tags for a project; put array of tags, that should be finally assigned to a project"`
	Avatar                                    *string   `flag_name:"avatar" type:"mixed" required:"no" description:"Image file for avatar of the project"`
	PrintingMergeRequestLinkEnabled           *bool     `flag_name:"printing_merge_request_link_enabled" type:"bool" required:"no" description:"Show link to create/view merge request when pushing from the command line"`
	CiConfigPath                              *string   `flag_name:"ci_config_path" type:"string" required:"no" description:"The path to CI config file"`
}

var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  `Create a new project for the given parameters`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO add this to use name of group instead of namespace_id
		//groups, _, err := gitlabClient.Groups.SearchGroup(group)
		//if err != nil {
		//	// TODO make sure we stop here when namespace_id cannot be properly resolved
		//	return errors.New("An error occurred while detecting namespace ID for " + group + ":" + err.Error())
		//}
		//if len(groups) > 1 {
		//	return errors.New("More than one group was found for given group" + group)
		//}
		//
		//p := &gitlab.CreateProjectOptions{
		//	Name:        &name,
		//	NamespaceID: &groups[0].ID,
		//}

		opts, err := createProjectOpts()
		if err != nil {
			return err
		}
		project, _, err := gitlabClient.Projects.CreateProject(opts)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

func createProjectOpts() (*gitlab.CreateProjectOptions, error) {
	flags := &createOpts{}
	opts := &gitlab.CreateProjectOptions{}
	createOptsMapper.Map(flags, opts)
	if flags.Visibility != nil {
		opts.Visibility = str2Visibility(*flags.Visibility)
	}
	return opts, nil
}

type editOpts struct {
	Id                                        *string   `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Name                                      *string   `flag_name:"name" type:"string" required:"yes" description:"The name of the project"`
	Path                                      *string   `flag_name:"path" type:"string" required:"no" description:"Custom repository name for the project. By default generated based on name"`
	DefaultBranch                             *string   `flag_name:"default_branch" type:"string" required:"no" description:"master by default"`
	Description                               *string   `flag_name:"description" type:"string" required:"no" description:"Short project description"`
	IssuesEnabled                             *bool     `flag_name:"issues_enabled" type:"bool" required:"no" description:"Enable issues for this project"`
	MergeRequestsEnabled                      *bool     `flag_name:"merge_requests_enabled" type:"bool" required:"no" description:"Enable merge requests for this project"`
	JobsEnabled                               *bool     `flag_name:"jobs_enabled" type:"bool" required:"no" description:"Enable jobs for this project"`
	WikiEnabled                               *bool     `flag_name:"wiki_enabled" type:"bool" required:"no" description:"Enable wiki for this project"`
	SnippetsEnabled                           *bool     `flag_name:"snippets_enabled" type:"bool" required:"no" description:"Enable snippets for this project"`
	ResolveOutdatedDiffDiscussions            *bool     `flag_name:"resolve_outdated_diff_discussions" type:"bool" required:"no" description:"Automatically resolve merge request diffs discussions on lines changed with a push"`
	ContainerRegistryEnabled                  *bool     `flag_name:"container_registry_enabled" type:"bool" required:"no" description:"Enable container registry for this project"`
	SharedRunnersEnabled                      *bool     `flag_name:"shared_runners_enabled" type:"bool" required:"no" description:"Enable shared runners for this project"`
	Visibility                                *string   `flag_name:"visibility" type:"string" required:"no" description:"See project visibility level"`
	ImportUrl                                 *string   `flag_name:"import_url" type:"string" required:"no" description:"URL to import repository from"`
	PublicJobs                                *bool     `flag_name:"public_jobs" type:"bool" required:"no" description:"If true, jobs can be viewed by non-project-members"`
	OnlyAllowMergeIfPipelineSucceeds          *bool     `flag_name:"only_allow_merge_if_pipeline_succeeds" type:"bool" required:"no" description:"Set whether merge requests can only be merged with successful jobs"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool     `flag_name:"only_allow_merge_if_all_discussions_are_resolved" type:"bool" required:"no" description:"Set whether merge requests can only be merged when all the discussions are resolved"`
	LfsEnabled                                *bool     `flag_name:"lfs_enabled" type:"bool" required:"no" description:"Enable LFS"`
	RequestAccessEnabled                      *bool     `flag_name:"request_access_enabled" type:"bool" required:"no" description:"Allow users to request member access"`
	TagList                                   *[]string `flag_name:"tag_list" type:"array" required:"no" description:"The list of tags for a project; put array of tags, that should be finally assigned to a project"`
	CiConfigPath                              *string   `flag_name:"ci_config_path" type:"string" required:"no" description:"The path to CI config file"`
}

var projectEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit project",
	Long:  `Updates an existing project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := editProjectOpts()
		if err != nil {
			return err
		}
		pid, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		project, _, err := gitlabClient.Projects.EditProject(pid, opts)
		return OutputJson(project)
	},
}

func editProjectOpts() (*gitlab.EditProjectOptions, error) {
	flags := &editOpts{}
	opts := &gitlab.EditProjectOptions{}
	editOptsMapper.Map(flags, opts)
	return opts, nil
}

type forkOpts struct {
	Id        *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Namespace *string `flag_name:"namespace" type:"integer/string" required:"yes" description:"The ID or path of the namespace that the project will be forked to"`
}

var projectForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Fork project",
	Long: `Forks a project into the user namespace of the authenticated user or the one provided.

The forking operation for a project is asynchronous and is completed in a background job. The request will return immediately. To determine whether the fork of the project has completed, query the import_status for the new project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := forkProjectOpts()
		// TODO target namespace is currently not supported by go-gitlab
		project, _, err := gitlabClient.Projects.ForkProject(*opts.Id)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

func forkProjectOpts() (*forkOpts, error) {
	flags := &forkOpts{}
	opts := &forkOpts{}
	forkOptsMapper.Map(flags, opts)
	return opts, nil
}

type listForksOpts struct {
	Id                       *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Archived                 *bool   `flag_name:"archived" type:"bool" required:"no" description:"Limit by archived status"`
	Visibility               *string `flag_name:"visibility" type:"string" required:"no" description:"Limit by visibility public, internal, or private"`
	OrderBy                  *string `flag_name:"order_by" type:"string" required:"no" description:"Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at"`
	Sort                     *string `flag_name:"sort" type:"string" required:"no" description:"Return projects sorted in asc or desc order. Default is desc"`
	Search                   *string `flag_name:"search" type:"string" required:"no" description:"Return list of projects matching the search criteria"`
	Simple                   *bool   `flag_name:"simple" type:"bool" required:"no" description:"Return only the ID, URL, name, and path of each project"`
	Owned                    *bool   `flag_name:"owned" type:"bool" required:"no" description:"Limit by projects owned by the current user"`
	Membership               *bool   `flag_name:"membership" type:"bool" required:"no" description:"Limit by projects that the current user is a member of"`
	Starred                  *bool   `flag_name:"starred" type:"bool" required:"no" description:"Limit by projects starred by the current user"`
	Statistics               *bool   `flag_name:"statistics" type:"bool" required:"no" description:"Include project statistics"`
	WithIssuesEnabled        *bool   `flag_name:"with_issues_enabled" type:"bool" required:"no" description:"Limit by enabled issues feature"`
	WithMergeRequestsEnabled *bool   `flag_name:"with_merge_requests_enabled" type:"bool" required:"no" description:"Limit by enabled merge requests feature"`
}

var projectListForksCmd = &cobra.Command{
	Use:   "list-forks",
	Short: "List Forks of a project",
	Long:  `List the projects accessible to the calling user that have an established, forked relationship with the specified project (available since Gitlab 10.1).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO currently not available in go-gitlab
		//opts := listForkOpts()
		return errors.New("list forks of a project is currently not implemented")
	},
}

var projectStarCmd = &cobra.Command{
	Use: "star",
	Short: "Star a project ",
	Long: `Stars a given project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pid, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		project, _, err := gitlabClient.Projects.StarProject(pid)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

var projectUnstarCmd = &cobra.Command{
	Use: "unstar",
	Short: "Unstar a project",
	Long: `Unstars a given project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pid, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		project, _, err := gitlabClient.Projects.UnstarProject(pid)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

// TODO currently not available in go-gitlab
//func listForkOpts() gitlab.ListForkOptions, err {
//	flags := &listForksOpts{}
//	opts := &gitlab.ListForkOptions{}
//	listForksOptsMapper.Map(flags, opts)
//	return opts, nil
//}

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing project",
	Long:  `Delete an existing project by either its project ID or namespace/project-name`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO maybe we want to return something upon success
		// TODO do something useful with the response
		_, err := gitlabClient.Projects.DeleteProject(id)
		return err
	},
}

func parsePid(value string) interface{} {
	if pid, err := strconv.Atoi(value); err == nil {
		return pid
	} else {
		return value
	}
}

func init() {
	initProjectListCmd()
	initProjectGetCmd()
	initProjectCreateCmd()
	initProjectEditCmd()
	initProjectForkCmd()
	initProjectListForksCmd()
	initProjectStarCmd()
	initProjectUnstarCmd()
	initProjectDeleteCommand()
	RootCmd.AddCommand(projectCmd)
}

func initProjectListCmd() {
	flags := &listFlags{}
	listOptsMapper = mapper.New(projectListCmd)
	listOptsMapper.SetFlags(flags)
	projectCmd.AddCommand(projectListCmd)
}

func initProjectGetCmd() {
	flags := &getFlags{}
	getOptsMapper = mapper.New(projectGetCmd)
	getOptsMapper.SetFlags(flags)
	projectCmd.AddCommand(projectGetCmd)
}

func initProjectCreateCmd() {
	flags := &createOpts{}
	createOptsMapper = mapper.New(projectCreateCmd)
	createOptsMapper.SetFlags(flags)
	projectCmd.AddCommand(projectCreateCmd)
}

func initProjectEditCmd() {
	flags := &editOpts{}
	editOptsMapper = mapper.New(projectEditCmd)
	editOptsMapper.SetFlags(flags)
	projectCmd.AddCommand(projectEditCmd)
}

func initProjectForkCmd() {
	flags := &forkOpts{}
	forkOptsMapper = mapper.New(projectForkCmd)
	forkOptsMapper.SetFlags(flags)
	projectCmd.AddCommand(projectForkCmd)
}

func initProjectListForksCmd() {
	flags := &listForksOpts{}
	listForksOptsMapper = mapper.New(projectListForksCmd)
	listForksOptsMapper.SetFlags(flags)
	projectCmd.AddCommand(projectListForksCmd)
}

func initProjectStarCmd() {
	projectStarCmd.PersistentFlags().StringP("id", "i", "", "(required) The ID or URL-encoded path of the project")
	projectCmd.AddCommand(projectStarCmd)
}

func initProjectUnstarCmd() {
	projectUnstarCmd.PersistentFlags().StringP("id", "i", "", "(required) The ID or URL-encoded path of the project")
	projectCmd.AddCommand(projectUnstarCmd)
}

func initProjectDeleteCommand() {
	projectDeleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) Either ID of project or 'namespace/project-name'")
	projectCmd.AddCommand(projectDeleteCmd)
}
