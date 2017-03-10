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
)

var name string

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long: `List, create, edit and delete projects`,
	Run: func(cmd *cobra.Command, args []string) {
		projects := gitlabClient.Projects.List()
		json, _ := json.MarshalIndent(projects, "", "  ")
		fmt.Println(string(json))
	},
}

var createProjectCmd = &cobra.Command{
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
			json, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(json))
		}
	},
}

func init() {
	createProjectCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the project")
	viper.BindPFlag("name", createProjectCmd.PersistentFlags().Lookup("name"))
	projectCmd.AddCommand(createProjectCmd)
	RootCmd.AddCommand(projectCmd)
}
