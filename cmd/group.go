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
)

var statistics bool

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage Gitlab Groups",
	Long: `Show, create, update and delete Gitlab groups.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("check usage of group with `golab group -h`")
	},
}

var groupLsCmd = &cobra.Command{
	Use: "ls",
	Short: "List groups",
	Long: `Get a list of visible groups for the authenticated user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := &gitlab.ListGroupsOptions{}
		if statistics == true {
			opts.Statistics = &statistics
		}
		groups, _, err := gitlabClient.Groups.ListGroups(opts)
		if err != nil { return err }
		return OutputJson(groups)
	},
}

var groupProjectsCmd = &cobra.Command{
	Use: "projects",
	Short: "List a group's projects",
	Long: `Get a list of projects in this group. When accessed without authentication, only public projects are returned.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("missing parameter `id`")
		}
		opts := &gitlab.ListGroupProjectsOptions{}
		projects, _, err := gitlabClient.Groups.ListGroupProjects(id, opts)
		if err != nil { return err }
		return OutputJson(projects)
	},
}

var groupGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get detailed information for a group",
	Long: `Get detailed information for a group identified by either ID or the namespace / path of the group`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}
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
			return errors.New("required parameter `-n` or `--name` not given - exiting")
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
	initGroupLsCommand()
	initGroupGetCommand()
	initGroupCreateCommand()
	initGroupProjectsCommand()
	RootCmd.AddCommand(groupCmd)
}

func initGroupLsCommand() {
	groupLsCmd.PersistentFlags().BoolVarP(&statistics, "statistics", "s", false, "(optional) if set to true, additional statistics are shown (admin only)")
	viper.BindPFlag("statistics", groupLsCmd.PersistentFlags().Lookup("statistics"))
	groupCmd.AddCommand(groupLsCmd)
}

func initGroupProjectsCommand() {
	groupProjectsCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to list projects for")
	viper.BindPFlag("id", groupProjectsCmd.PersistentFlags().Lookup("id"))
	groupCmd.AddCommand(groupProjectsCmd)
}

func initGroupGetCommand() {
	groupGetCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) either ID or namespace of group")
	viper.BindPFlag("id", groupGetCmd.PersistentFlags().Lookup("id"))
	groupCmd.AddCommand(groupGetCmd)
}

func initGroupCreateCommand() {
	groupCreateCommand.PersistentFlags().StringVarP(&name, "name", "n", "", "(required) name of the new group")
	viper.BindPFlag("name", groupCreateCommand.PersistentFlags().Lookup("name"))
	groupCmd.AddCommand(groupCreateCommand)
}
