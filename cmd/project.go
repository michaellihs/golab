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

	. "github.com/michaellihs/golab/cmd/helpers"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// see https://docs.gitlab.com/ce/api/projects.html
var projectsCmd = &golabCommand{
	Parent: RootCmd,
	Cmd: &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
		Long:  `List, create, edit and delete projects`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot run this command without further sub-commands")
	},
}

// see https://docs.gitlab.com/ce/api/projects.html#list-all-projects
type projectsListFlags struct {
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

var projectsListCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &projectsListFlags{},
	Opts:   &gitlab.ListProjectsOptions{},
	Paged:  true,
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List all projects",
		Long:  `Get a list of all visible projects across GitLab for the authenticated user.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.ListProjectsOptions)
		projects, _, err := gitlabClient.Projects.ListProjects(opts)
		if err != nil {
			return err
		}
		return OutputJson(projects)
	},
}

type getFlags struct {
	Id *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"either the project ID (numeric) or 'namespace/project-name'"`
	// TODO currently not supported by go-gitlab
	Statistics *bool `flag_name:"statistics" short:"s" required:"no" description:"include project statistics"`
}

var projectGetCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &getFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get detailed information for a project",
		Long:  `Get detailed information for a project identified by either project ID or 'namespace/project-name'`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*getFlags)
		project, _, err := gitlabClient.Projects.GetProject(parsePid(*flags.Id)) // make sure, parsedPid is of type int if numeric
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type createFlags struct {
	Name                                      *string   `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the new project"`
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

var projectCreateCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &createFlags{},
	Opts:   &gitlab.CreateProjectOptions{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Long:  `Create a new project for the given parameters`,
	},
	Run: func(cmd golabCommand) error {
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
		opts := cmd.Opts.(*gitlab.CreateProjectOptions)
		project, _, err := gitlabClient.Projects.CreateProject(opts)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type editFlags struct {
	Id                                        *string   `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Name                                      *string   `flag_name:"name" short:"n" type:"string" required:"yes" description:"The name of the project"`
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

var projectEditCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &editFlags{},
	Opts:   &gitlab.EditProjectOptions{},
	Cmd: &cobra.Command{
		Use:   "edit",
		Short: "Edit project",
		Long:  `Updates an existing project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*editFlags)
		opts := cmd.Opts.(*gitlab.EditProjectOptions)
		project, _, err := gitlabClient.Projects.EditProject(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type forkFlags struct {
	Id        *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	Namespace *string `flag_name:"namespace" short:"n" type:"integer/string" required:"yes" description:"The ID or path of the namespace that the project will be forked to"`
}

var projectForkCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &forkFlags{},
	Cmd: &cobra.Command{
		Use:   "fork",
		Short: "Fork project",
		Long:  `Forks a project into the user namespace of the authenticated user or the one provided.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*forkFlags)
		// TODO target namespace is currently not supported by go-gitlab
		project, _, err := gitlabClient.Projects.ForkProject(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type listForksFlags struct {
	Id                       *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
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

var projectListForksCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &listForksFlags{},
	Cmd: &cobra.Command{
		Use:   "list-forks",
		Short: "List Forks of a project",
		Long:  `List the projects accessible to the calling user that have an established, forked relationship with the specified project (available since Gitlab 10.1).`,
	},
	Run: func(cmd golabCommand) error {
		// TODO currently not available in go-gitlab
		return errors.New("list forks of a project is currently not implemented")
	},
}

type starFlags struct {
	Id *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var projectStarCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &starFlags{},
	Cmd: &cobra.Command{
		Use:   "star",
		Short: "Star a project ",
		Long:  `Stars a given project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*starFlags)
		project, _, err := gitlabClient.Projects.StarProject(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type unstarFlags struct {
	Id *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var projectUnstarCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &unstarFlags{},
	Cmd: &cobra.Command{
		Use:   "unstar",
		Short: "Unstar a project",
		Long:  `Unstars a given project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*unstarFlags)
		project, _, err := gitlabClient.Projects.UnstarProject(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type archiveFlags struct {
	Id *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var projectArchiveCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &archiveFlags{},
	Cmd: &cobra.Command{
		Use:   "archive",
		Short: "Archive a project",
		Long:  `Archives the project if the user is either admin or the project owner of this project. This action is idempotent, thus archiving an already archived project will not change the project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*archiveFlags)
		project, _, err := gitlabClient.Projects.ArchiveProject(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type unarchiveFlags struct {
	Id *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var projectUnarchiveCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &unarchiveFlags{},
	Cmd: &cobra.Command{
		Use:   "unarchive",
		Short: "Unarchive a project",
		Long:  `Unarchives the project if the user is either admin or the project owner of this project. This action is idempotent, thus unarchiving an non-archived project will not change the project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*unarchiveFlags)
		project, _, err := gitlabClient.Projects.UnarchiveProject(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

type deleteFlags struct {
	Id *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var projectDeleteCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &deleteFlags{},
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Remove project",
		Long:  `Removes a project including all associated resources (issues, merge requests etc.)`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*deleteFlags)
		_, err := gitlabClient.Projects.DeleteProject(*flags.Id)
		return err
	},
}

type uploadFlags struct {
	Id   *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	File *string `flag_name:"file" short:"f" type:"string" required:"yes" description:"Path to the file to be uploaded"`
}

var projectUploadFileCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &uploadFlags{},
	Cmd: &cobra.Command{
		Use:   "upload-file",
		Short: "Upload a file",
		Long:  `Uploads a file to the specified project to be used in an issue or merge request description, or a comment.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*uploadFlags)
		projectFile, _, err := gitlabClient.Projects.UploadFile(*flags.Id, *flags.File)
		if err != nil {
			return err
		}
		return OutputJson(projectFile)
	},
}

type shareFlags struct {
	Id          *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	GroupID     *int    `flag_name:"group_id" short:"g" type:"integer" required:"yes" description:"The ID of the group to share with"`
	GroupAccess *string `flag_name:"group_access" short:"a" type:"integer" transform:"str2AccessLevel" required:"yes" description:"The permissions level to grant the group"`
	ExpiresAt   *string `flag_name:"expires_at" short:"e" type:"string" required:"no" description:"Share expiration date in ISO 8601 format: 2016-09-26"`
	// gitlab opts should use ISOTime instead of string, then this line is valid:
	//ExpiresAt   *string  `flag_name:"expires_at" short:"e" type:"string" transform:"string2IsoTime" required:"no" description:"Share expiration date in ISO 8601 format: 2016-09-26"`
}

var projectShareWithGroupCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &shareFlags{},
	Opts:   &gitlab.ShareWithGroupOptions{},
	Cmd: &cobra.Command{
		Use:   "share",
		Short: "Share project with group",
		Long:  `Allow to share project with group.`,
	},
	Run: func(cmd golabCommand) error {
		opts := cmd.Opts.(*gitlab.ShareWithGroupOptions)
		flags := cmd.Flags.(*shareFlags)
		_, err := gitlabClient.Projects.ShareProjectWithGroup(*flags.Id, opts)
		return err
	},
}

type unshareFlags struct {
	Id      *string `flag_name:"id" short:"i" type:"string" required:"yes" description:"The ID or URL-encoded path of the project"`
	GroupID *int    `flag_name:"group_id" short:"g" type:"integer" required:"yes" description:"The ID of the group to share with"`
}

var projectUnshareWithGroupCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &unshareFlags{},
	Cmd: &cobra.Command{
		Use:   "unshare",
		Short: "Delete a shared project link within a group",
		Long:  `Unshare the project from the group.`,
	},
	Run: func(cmd golabCommand) error {
		// TODO delete a share is currently missing in go-gitlab
		return errors.New("not implemented...")
	},
}

var projectHooksCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Cmd: &cobra.Command{
		Use:   "hooks",
		Short: "Manage project hooks.",
		Long:  `Also called Project Hooks and Webhooks. These are different for System Hooks that are system wide.`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot run this command without further sub-commands")
	},
}

type hooksListFlags struct {
	Id *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
}

var projectHooksListCmd = &golabCommand{
	Parent: projectHooksCmd.Cmd,
	Flags:  &hooksListFlags{},
	Opts:   &gitlab.ListProjectHooksOptions{},
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List project hooks",
		Long:  `Get a list of project hooks.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*hooksListFlags)
		opts := cmd.Opts.(*gitlab.ListProjectHooksOptions)
		hooks, _, err := gitlabClient.Projects.ListProjectHooks(*flags.Id, opts)
		if err != nil {
			return err
		}
		return OutputJson(hooks)
	},
}

type hooksGetFlags struct {
	Id     *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	HookId *int    `flag_name:"hook_id" short:"h" type:"integer" required:"yes" description:"The ID of the hook"`
}

var projectHooksGetCmd = &golabCommand{
	Parent: projectHooksCmd.Cmd,
	Flags:  &hooksGetFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Get project hook",
		Long:  `Get a specific hook for a project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*hooksGetFlags)
		hook, _, err := gitlabClient.Projects.GetProjectHook(*flags.Id, *flags.HookId)
		if err != nil {
			return err
		}
		return OutputJson(hook)
	},
}

type addHookFlags struct {
	Id                    *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	URL                   *string `flag_name:"url" short:"u" type:"string" required:"yes" description:"The hook URL"`
	PushEvents            *bool   `flag_name:"push_events" type:"bool" required:"no" description:"Trigger hook on push events"`
	IssuesEvents          *bool   `flag_name:"issues_events" type:"bool" required:"no" description:"Trigger hook on issues events"`
	MergeRequestsEvents   *bool   `flag_name:"merge_requests_events" type:"bool" required:"no" description:"Trigger hook on merge requests events"`
	TagPushEvents         *bool   `flag_name:"tag_push_events" type:"bool" required:"no" description:"Trigger hook on tag push events"`
	NoteEvents            *bool   `flag_name:"note_events" type:"bool" required:"no" description:"Trigger hook on note events"`
	JobEvents             *bool   `flag_name:"job_events" type:"bool" required:"no" description:"Trigger hook on job events"`
	PipelineEvents        *bool   `flag_name:"pipeline_events" type:"bool" required:"no" description:"Trigger hook on pipeline events"`
	WikiEvents            *bool   `flag_name:"wiki_events" type:"bool" required:"no" description:"Trigger hook on wiki events"`
	EnableSslVerification *bool   `flag_name:"enable_ssl_verification" type:"bool" required:"no" description:"Do SSL verification when triggering the hook"`
	Token                 *string `flag_name:"token" type:"string" required:"no" description:"Secret token to validate received payloads; this will not be returned in the response"`
}

var projectAddHookCmd = &golabCommand{
	Parent: projectHooksCmd.Cmd,
	Flags:  &addHookFlags{},
	Opts:   &gitlab.AddProjectHookOptions{},
	Cmd: &cobra.Command{
		Use:   "add",
		Short: "Add project hook",
		Long:  `Adds a hook to a specified project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*addHookFlags)
		opts := cmd.Opts.(*gitlab.AddProjectHookOptions)
		hook, _, err := gitlabClient.Projects.AddProjectHook(parsePid(*flags.Id), opts)
		if err != nil {
			return err
		}
		return OutputJson(hook)
	},
}

type editHookFlags struct {
	Id                    *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	HookId                *int    `flag_name:"hook_id" type:"integer" required:"yes" description:"The ID of the project hook"`
	URL                   *string `flag_name:"url" short:"u" type:"string" required:"yes" description:"The hook URL"`
	PushEvents            *bool   `flag_name:"push_events" type:"bool" required:"no" description:"Trigger hook on push events"`
	IssuesEvents          *bool   `flag_name:"issues_events" type:"bool" required:"no" description:"Trigger hook on issues events"`
	MergeRequestsEvents   *bool   `flag_name:"merge_requests_events" type:"bool" required:"no" description:"Trigger hook on merge requests events"`
	TagPushEvents         *bool   `flag_name:"tag_push_events" type:"bool" required:"no" description:"Trigger hook on tag push events"`
	NoteEvents            *bool   `flag_name:"note_events" type:"bool" required:"no" description:"Trigger hook on note events"`
	JobEvents             *bool   `flag_name:"job_events" type:"bool" required:"no" description:"Trigger hook on job events"`
	PipelineEvents        *bool   `flag_name:"pipeline_events" type:"bool" required:"no" description:"Trigger hook on pipeline events"`
	WikiEvents            *bool   `flag_name:"wiki_events" type:"bool" required:"no" description:"Trigger hook on wiki events"`
	EnableSslVerification *bool   `flag_name:"enable_ssl_verification" type:"bool" required:"no" description:"Do SSL verification when triggering the hook"`
	Token                 *string `flag_name:"token" type:"string" required:"no" description:"Secret token to validate received payloads; this will not be returned in the response"`
}

var projectEditHookCmd = &golabCommand{
	Parent: projectHooksCmd.Cmd,
	Flags:  &editHookFlags{},
	Opts:   &gitlab.EditProjectHookOptions{},
	Cmd: &cobra.Command{
		Use:   "edit",
		Short: "Edit project hook",
		Long:  `Edits a hook for a specified project.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*editHookFlags)
		opts := cmd.Opts.(*gitlab.EditProjectHookOptions)
		hook, _, err := gitlabClient.Projects.EditProjectHook(parsePid(*flags.Id), *flags.HookId, opts)
		if err != nil {
			return err
		}
		return OutputJson(hook)
	},
}

type hooksDeleteFlags struct {
	Id     *string `flag_name:"id" short:"i" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the project"`
	HookId *int    `flag_name:"hook_id" short:"h" type:"integer" required:"yes" description:"The ID of the hook"`
}

var projectDeleteHookCmd = &golabCommand{
	Parent: projectHooksCmd.Cmd,
	Flags:  &hooksDeleteFlags{},
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Delete project hook",
		Long:  `Removes a hook from a project. This is an idempotent method and can be called multiple times. Either the hook is available or not.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*hooksDeleteFlags)
		_, err := gitlabClient.Projects.DeleteProjectHook(*flags.Id, *flags.HookId)
		return err
	},
}

var projectForskCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Cmd: &cobra.Command{
		Use:   "forks",
		Short: "Admin fork relation",
		Long:  `Allows modification of the forked relationship between existing projects. Available only for admins.`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("cannot run this command without a sub-command")
	},
}

// see https://docs.gitlab.com/ce/api/projects.html#create-a-forked-from-to-relation-between-existing-projects
type forkCreateFlags struct {
	Id           *int `flag_name:"id" short:"i" type:"integer" required:"yes" description:"The ID or URL-encoded path of the project"`
	ForkedFromId *int `flag_name:"forked_from_id" short:"f" type:"integer" required:"yes" description:"The ID of the project that was forked from"`
}

var projectForksCreateCmd = &golabCommand{
	Parent: projectForskCmd.Cmd,
	Flags:  &forkCreateFlags{},
	Cmd: &cobra.Command{
		Use:   "create",
		Short: "Create a forked from/to relation between existing projects",
		Long:  `Create a forked from/to relation between existing projects (available only for admins)`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*forkCreateFlags)
		// go-gitlab assumes that a JSON with the fork info is returned, but it seems like only "OK" is returned upon success
		_, res, err := gitlabClient.Projects.CreateProjectForkRelation(*flags.Id, *flags.ForkedFromId)
		if res.Status != "201 Created" {
			return err
		}
		return nil
	},
}

// see https://docs.gitlab.com/ce/api/projects.html#delete-an-existing-forked-from-relationship
type forkDeleteFlags struct {
	Id *int `flag_name:"id" short:"i" type:"integer" required:"yes" description:"The ID of the project"`
}

var projectForksDeleteCmd = &golabCommand{
	Parent: projectForskCmd.Cmd,
	Flags:  &forkDeleteFlags{},
	Cmd: &cobra.Command{
		Use:   "delete",
		Short: "Delete an existing forked from relationship",
		Long:  `Delete an existing forked from relationship (available only for admins)`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*forkDeleteFlags)
		_, err := gitlabClient.Projects.DeleteProjectForkRelation(*flags.Id)

		return err
	},
}

// see https://docs.gitlab.com/ce/api/projects.html#search-for-projects-by-name
type projectSearchFlags struct {
	Search  *string `flag_name:"search" short:"s" type:"string" required:"yes" description:"A string contained in the project name"`
	OrderBy *string `flag_name:"order_by" type:"string" required:"no" description:"Return requests ordered by id, name, created_at or last_activity_at fields"`
	Sort    *string `flag_name:"sort" type:"string" required:"no" description:"Return requests sorted in asc or desc order"`
}

var projectSearchCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Flags:  &projectSearchFlags{},
	Cmd: &cobra.Command{
		Use:   "search",
		Short: "Search for projects by name",
		Long:  `Search for projects by name which are accessible to the authenticated user. This endpoint can be accessed without authentication if the project is publicly accessible.`,
	},
	Run: func(cmd golabCommand) error {
		return errors.New("currently not implemented")
		// TODO not implemented in go-gitlab so far
		//projectSearchOptsMapper.AutoMap()
		//opts := projectSearchOptsMapper.MappedOpts().(*gitlab...)
		//flags := projectSearchOptsMapper.MappedFlags().(*projectSearchFlags)
		//gitlabClient.Projects.Search()
	},
}

var projectHousekeepingCmd = &golabCommand{
	Parent: projectsCmd.Cmd,
	Cmd: &cobra.Command{
		Use:   "housekeeping",
		Short: "Start the Housekeeping task for a Project",
		Long:  `Start the Housekeeping task for a Project`,
	},
	Run: func(cmd golabCommand) error {
		// TODO this feature is missing in go-gitlab
		return errors.New("this feature is not yet implemented")
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
	projectsCmd.Init()
	projectsListCmd.Init()
	projectEditCmd.Init()
	projectGetCmd.Init()
	projectCreateCmd.Init()
	projectForkCmd.Init()
	projectListForksCmd.Init()
	projectStarCmd.Init()
	projectUnstarCmd.Init()
	projectArchiveCmd.Init()
	projectUnarchiveCmd.Init()
	projectDeleteCmd.Init()
	projectUploadFileCmd.Init()
	projectShareWithGroupCmd.Init()
	projectUnshareWithGroupCmd.Init()
	projectHooksCmd.Init()
	projectHooksListCmd.Init()
	projectHooksGetCmd.Init()
	projectAddHookCmd.Init()
	projectEditHookCmd.Init()
	projectDeleteHookCmd.Init()
	projectForskCmd.Init()
	projectForksCreateCmd.Init()
	projectForksDeleteCmd.Init()
	projectSearchCmd.Init()
	projectHousekeepingCmd.Init()
}
