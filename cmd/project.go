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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"strconv"
)

var name string
var id int
var group string
var pid string

var archived, simple, owned, membership, starred bool
var orderBy, sort string

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long: `List, create, edit and delete projects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, _, err := gitlabClient.Projects.ListProjects(&gitlab.ListProjectsOptions{})
		if err != nil {
			return err
		}
		err = OutputJson(projects)
		return err
	},
}

// TODO custom attributes are currently not supported
var projectListCmd = &cobra.Command{
	Use: "ls",
	Short: "List all projects",
	Long: `Get a list of all visible projects across GitLab for the authenticated user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, _, err := gitlabClient.Projects.ListProjects(flagsToListOptions())
		if err != nil { return err }
		return OutputJson(projects)
	},
}

var projectGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get detailed information for a project",
	Long: `Get detailed information for a project identified by either project ID or 'namespace/project-name'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if pid == "" {
			return errors.New("you have to provide a project ID or 'namespace/project-name' with the -i --id flag")
		}
		parsedPid := parsePid(pid) // make sure, parsedPid is of type int if numeric
		project, _, err := gitlabClient.Projects.GetProject(parsedPid)
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
}

func parsePid(value string) interface{} {
	if pid, err := strconv.Atoi(value); err == nil {
		return pid
	} else {
		return value
	}
}

var projectCreateCmd = &cobra.Command{
	Use: "create",
	Short: "Create a new project",
	Long: `Create a new project for the given parameters`,
	RunE: func(cmd *cobra.Command, args []string) error {
		groups, _, err := gitlabClient.Groups.SearchGroup(group)
		if err != nil {
			// TODO make sure we stop here when namespace_id cannot be properly resolved
			return errors.New("An error occurred while detecting namespace ID for " + group + ":" + err.Error())
		}
		if len(groups) > 1 {
			return errors.New("More than one group was found for given group" + group)
		}

		p := &gitlab.CreateProjectOptions{
			Name: &name,
			NamespaceID: &groups[0].ID,
		}

		// TODO do something useful with the response
		project, _, err := gitlabClient.Projects.CreateProject(p)
		if err != nil {
			return err
		}
		err = OutputJson(project)
		return err
	},
}

var projectDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete an existing project",
	Long: `Delete an existing project by either its project ID or namespace/project-name`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO maybe we want to return something upon success
		// TODO do something useful with the response
		_, err := gitlabClient.Projects.DeleteProject(id)
		return err
	},
}

func flagsToListOptions() *gitlab.ListProjectsOptions {
	listOptions := &gitlab.ListProjectsOptions{
		Archived: &archived,
		Membership: &membership,
		Owned: &owned,
		Search: &search,
		Simple: &simple,
		Starred: &starred,
		Statistics: &statistics,
	}
	if orderBy != "" {
		listOptions.OrderBy = &orderBy
	}
	if sort != "" {
		listOptions.Sort = &sort
	}
	if visibility != "" {
		listOptions.Visibility = str2Visibility(visibility)
	}
	return listOptions
}

func init() {
	initProjectGetCommand()
	initProjectCreateCommand()
	initProjectDeleteCommand()
	initProjectLsCommand()
	RootCmd.AddCommand(projectCmd)
}

func initProjectLsCommand() {
	projectListCmd.PersistentFlags().BoolVarP(&archived, "archived", "a", false, "(optional) Limit by archived status")
	projectListCmd.PersistentFlags().StringVarP(&visibility, "visibility", "v", "", "(optional) Limit by visibility public, internal, or private")
	projectListCmd.PersistentFlags().StringVarP(&orderBy, "order_by", "o", "", "(optional) Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at")
	projectListCmd.PersistentFlags().StringVar(&sort, "sort", "", "(optional) Return projects sorted in asc or desc order. Default is desc")
	projectListCmd.PersistentFlags().StringVar(&search, "search", "", "(optional) Return list of projects matching the search criteria")
	projectListCmd.PersistentFlags().BoolVarP(&simple, "simple", "s", false, "(optional) Return only the ID, URL, name, and path of each project")
	projectListCmd.PersistentFlags().BoolVarP(&owned, "owned", "", false, "(optional) Limit by projects owned by the current user")
	projectListCmd.PersistentFlags().BoolVarP(&membership, "membership", "m", false, "(optional) Limit by projects that the current user is a member of")
	projectListCmd.PersistentFlags().BoolVar(&starred, "starred", false, "(optional) Limit by projects starred by the current user")
	projectListCmd.PersistentFlags().BoolVar(&statistics, "statistics", false, "(optional) Include project statistics")
	// TODO not supported by go-gitlab
	//projectListCmd.PersistentFlags().BoolVarP(listOptions.with_issues_enabled, "	with_issues_enabled", "", false, "(optional) Limit by enabled issues feature")
	// TODO not supported by go-gitlab
	//projectListCmd.PersistentFlags().BoolVarP(listOptions.with_merge_requests_enabled, "	with_merge_requests_enabled", "", false, "(optional) Limit by enabled merge requests feature	")
	projectCmd.AddCommand(projectListCmd)
}

func initProjectGetCommand() {
	projectGetCmd.PersistentFlags().StringVarP(&pid, "id", "i", "", "(required) Either the project ID (numeric) or 'namespace/project-name'")
	projectCmd.AddCommand(projectGetCmd)
}

func initProjectCreateCommand() {
	projectCreateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "(required) Name of the project")
	projectCreateCmd.PersistentFlags().StringVarP(&group, "group", "g", "", "Group to add project to (either ID or namespace)")
	viper.BindPFlag("name", projectCreateCmd.PersistentFlags().Lookup("name"))
	projectCmd.AddCommand(projectCreateCmd)
}

func initProjectDeleteCommand() {
	projectDeleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) Either ID of project or 'namespace/project-name'")
	viper.BindPFlag("id", projectDeleteCmd.PersistentFlags().Lookup("id"))
	projectCmd.AddCommand(projectDeleteCmd)
}
