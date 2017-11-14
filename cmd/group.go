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
	"fmt"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/michaellihs/go-gitlab"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage Gitlab Groups",
	Long: `Show, create, update and delete Gitlab groups.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO how can we make Cobra do this
		fmt.Println("use one of the subcommands: get")
	},
}

var groupGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get detailed information for a group",
	Long: `Get detailed information for a group identified by either ID or the namespace / path of the group`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("Required parameter `-i` or `--id` not given. Exiting.")
		}
		// TODO do something useful with the response
		group, _, err := gitlabClient.Groups.GetGroup(id)
		if err != nil {
			return err
		}
		err = OutputJson(group)
		return err
	},
}

var groupCreateCommand = &cobra.Command{
	Use: "create",
	Short: "Create a new group",
	Long: `Create a new group for the given parameters`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if name == "" {
			return errors.New("Required parameter `-n` or `--name` not given. Exiting.")
		}
		group, _, err := gitlabClient.Groups.CreateGroup(&gitlab.CreateGroupOptions{Name: &name, Path: &name})
		if err != nil {
			return err
		}
		err = OutputJson(group)
		return err
	},
}

func init() {
	initGroupGetCommand()
	initGroupCreateCommand()
	RootCmd.AddCommand(groupCmd)
}

func initGroupGetCommand() {
	groupGetCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) Either ID or namespace of group")
	viper.BindPFlag("id", groupGetCmd.PersistentFlags().Lookup("id"))
	groupCmd.AddCommand(groupGetCmd)
}

func initGroupCreateCommand() {
	groupCreateCommand.PersistentFlags().StringVarP(&name, "name", "n", "", "(required) name of the new group")
	viper.BindPFlag("name", groupCreateCommand.PersistentFlags().Lookup("name"))
	groupCmd.AddCommand(groupCreateCommand)
}
