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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"errors"
	"github.com/xanzy/go-gitlab"
	"fmt"
)

var name string
var id string
var group string

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long: `List, create, edit and delete projects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO do something useful with the resonse
		projects, _, err := gitlabClient.Projects.ListAllProjects(&gitlab.ListProjectsOptions{})
		if err != nil {
			fmt.Println("kaputt")
			return err
		}
		err = OutputJson(projects)
		return err
	},
}

var projectGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get detailed information for a project",
	Long: `Get detailed information for a project identified by either project ID or 'namespace/project-name'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == "" {
			return errors.New("You have to provide a project ID or 'namespace/project-name' with the -i --id flag")
		}
		// TODO do something useful with the response
		project, _, err := gitlabClient.Projects.GetProject(id)
		if err != nil {
			return err
		}
		err = OutputJson(project)
		return err
	},
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

func init() {
	initProjectGetCommand()
	initProjectCreateCommand()
	initProjectDeleteCommand()
	RootCmd.AddCommand(projectCmd)
}

func initProjectGetCommand() {
	projectGetCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "(required) Either ID of project or 'namespace/project-name'")
	viper.BindPFlag("id", projectGetCmd.PersistentFlags().Lookup("id"))
	projectCmd.AddCommand(projectGetCmd)
}

func initProjectCreateCommand() {
	projectCreateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "(required) Name of the project")
	projectCreateCmd.PersistentFlags().StringVarP(&group, "group", "g", "", "Group to add project to (either ID or namespace)")
	viper.BindPFlag("name", projectCreateCmd.PersistentFlags().Lookup("name"))
	projectCmd.AddCommand(projectCreateCmd)
}

func initProjectDeleteCommand() {
	projectDeleteCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "(required) Either ID of project or 'namespace/project-name'")
	viper.BindPFlag("id", projectDeleteCmd.PersistentFlags().Lookup("id"))
	projectCmd.AddCommand(projectDeleteCmd)
}
