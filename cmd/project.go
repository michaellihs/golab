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
	"fmt"
	"github.com/michaellihs/golab/client"
	"github.com/spf13/viper"
	"encoding/json"
	"os"
)

var name string
var id string

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long: `List, create, edit and delete projects`,
	Run: func(cmd *cobra.Command, args []string) {
		projects := gitlabClient.Projects.List()
		result, _ := json.MarshalIndent(projects, "", "  ")
		fmt.Println(string(result))
	},
}

var projectGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get detailed information for a project",
	Long: `Get detailed information for a project identified by either project ID or 'namespace/project-name'`,
	Run: func(cmd *cobra.Command, args []string) {
		project, err := gitlabClient.Projects.Get(id)
		// TODO introduce generic check method for required params
		if id == "" {
			fmt.Println("You have to provide the a project ID or 'namespace/project-name' with the -i --id flag")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("An error occurred: " + err.Error())
		}
		result, _ := json.MarshalIndent(project, "", "  ")
		fmt.Println(string(result))
	},
}

var projectCreateCmd = &cobra.Command{
	Use: "create",
	Short: "Create a new project",
	Long: `Create a new project for the given parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		params := &client.ProjectParams{
			Name: name}
		project, err := gitlabClient.Projects.Create(params)
		if err != nil {
			fmt.Println("An error occurred: " + err.Error())
		} else {
			result, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(result))
		}
	},
}

var projectDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete an existing project",
	Long: `Delete an existing project by either its project ID or namespace/project-name`,
	Run: func(cmd *cobra.Command, args []string) {
		success, err := gitlabClient.Projects.Delete(id)
		if !success {
			fmt.Println("Something went wrong: " + err.Error())
		}
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
	viper.BindPFlag("name", projectCreateCmd.PersistentFlags().Lookup("name"))
	projectCmd.AddCommand(projectCreateCmd)
}

func initProjectDeleteCommand() {
	projectDeleteCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "(required) Either ID of project or 'namespace/project-name'")
	viper.BindPFlag("id", projectDeleteCmd.PersistentFlags().Lookup("id"))
	projectCmd.AddCommand(projectDeleteCmd)
}
