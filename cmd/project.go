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
	"reflect"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"github.com/spf13/viper"
	"github.com/jinzhu/copier"
)

var name string
var id int
var group string
var pid string

var createProjectParams, listProjectParams *interface{}

var flagMap = make(map[string]interface{})

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

type listOpts struct {
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
		var err error
		listProjectOpts := &listOpts{}
		listProjectParams, err = flagMapToOpts(listProjectOpts, cmd)
		if err != nil { return err }

		opts, err := currParams2listProjectOpts()

		projects, _, err := gitlabClient.Projects.ListProjects(opts)
		if err != nil {
			return err
		}
		return OutputJson(projects)
	},
}

var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get detailed information for a project",
	Long:  `Get detailed information for a project identified by either project ID or 'namespace/project-name'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if pid == "" {
			return errors.New("you have to provide a project ID or 'namespace/project-name' with the -i --id flag")
		}
		project, _, err := gitlabClient.Projects.GetProject(parsePid(pid)) // make sure, parsedPid is of type int if numeric
		if err != nil {
			return err
		}
		return OutputJson(project)
	},
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

		opts, err := createProjectOpts(cmd)
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

func createProjectOpts(cmd *cobra.Command) (*gitlab.CreateProjectOptions, error) {
	var err error
	createProjectOpts := &createOpts{}
	createProjectParams, err = flagMapToOpts(createProjectOpts, cmd)
	if err != nil {
		return nil, err
	}

	return currParams2createProjectOpts()
}

func currParams2createProjectOpts() (*gitlab.CreateProjectOptions, error) {
	//// TODO how can we de-pointer this?
	createOpts, ok := (*createProjectParams).(*createOpts)
	if !ok {
		return &gitlab.CreateProjectOptions{}, errors.New("casting of createOpts went wrong")
	}
	opts := &gitlab.CreateProjectOptions{}
	copier.Copy(opts, *createProjectParams)
	opts.Visibility = str2Visibility(*createOpts.Visibility)
	return opts, nil
}

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

func currParams2listProjectOpts() (*gitlab.ListProjectsOptions, error) {
	//listOpts, ok := (*listProjectParams).(*listOpts)
	//if !ok {
	//	return &gitlab.ListProjectsOptions{}, errors.New("casting of listOpts went wrong")
	//}
	opts := &gitlab.ListProjectsOptions{}
	copier.Copy(opts, *listProjectParams)
	return opts,nil
}



func parsePid(value string) interface{} {
	if pid, err := strconv.Atoi(value); err == nil {
		return pid
	} else {
		return value
	}
}

func init() {
	initProjectGetCommand()

	createOpts := &createOpts{}
	createProjectParams, _ = paramsToMap(createOpts, projectCreateCmd, projectCmd)

	listOpts := &listOpts{}
	listProjectParams, _ = paramsToMap(listOpts, projectListCmd, projectCmd)

	initProjectDeleteCommand()

	RootCmd.AddCommand(projectCmd)
}

func initProjectGetCommand() {
	projectGetCmd.PersistentFlags().StringVarP(&pid, "id", "i", "", "(required) Either the project ID (numeric) or 'namespace/project-name'")
	// TODO currently not supported by go-gitlab
	projectGetCmd.PersistentFlags().BoolVarP(&statistics, "statistics", "s", false, "(optional) Include project statistics")
	projectCmd.AddCommand(projectGetCmd)
}

func initProjectDeleteCommand() {
	projectDeleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) Either ID of project or 'namespace/project-name'")
	viper.BindPFlag("id", projectDeleteCmd.PersistentFlags().Lookup("id"))
	projectCmd.AddCommand(projectDeleteCmd)
}

func paramsToMap(opts interface{}, cmd *cobra.Command, baseCmd *cobra.Command) (*interface{}, error) {
	v := reflect.ValueOf(opts).Elem()
	for i := 0; i < v.NumField(); i++ {
		// this gives us the type of a struct field
		//fieldType := v.Field(i).Type().String()
		//fmt.Println(fieldType)

		// this gives us the name of a struct field
		//fieldName := v.Type().Field(i).Name
		//fmt.Println(fieldName)

		// this gives us the tag of a struct field
		tag := v.Type().Field(i).Tag
		f := v.Field(i)

		// avoid usage of flag map and directly set values to opts struct
		flagName := tag.Get("flag_name")
		switch f.Type().String() {
		case "*int":
			flagMap[flagName] = cmd.PersistentFlags().Int(flagName, 0, tag.Get("description"))
		case "*string":
			flagMap[flagName] = cmd.PersistentFlags().String(flagName, "", tag.Get("description"))
		case "*bool":
			flagMap[flagName] = cmd.PersistentFlags().Bool(flagName, false, tag.Get("description"))
		case "*[]string":
			flagMap[flagName] = cmd.PersistentFlags().StringArray(flagName, nil, tag.Get("description"))
		default:
			panic("Unknown type " + f.Type().String())
		}

	}
	baseCmd.AddCommand(cmd)
	return &opts, nil
}

func flagMapToOpts(opts interface{}, cmd *cobra.Command) (*interface{}, error) {
	v := reflect.ValueOf(opts).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		tag := v.Type().Field(i).Tag

		flagName := tag.Get("flag_name")
		flagChanged := cmd.PersistentFlags().Changed(flagName)

		// see https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
		// see https://stackoverflow.com/questions/40060131/reflect-assign-a-pointer-struct-value
		if f.IsValid() {
			// A Value can be changed only if it is addressable and was not obtained by the use of unexported struct fields.
			if f.CanSet() {
				if flagChanged {
					f.Set(reflect.ValueOf(flagMap[flagName]))
				} else {
					// TODO implement an additional tag that allows setting of "default" values
				}
			} else {
				return nil, errors.New("can not set " + flagName)
			}
		} else {
			return nil, errors.New(flagName + "is not valid")
		}
	}
	return &opts, nil
}
