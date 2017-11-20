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

var path, visibility, description string

var statistics, lfsEnabled, requestAccessEnabled bool

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
			return errors.New("required parameter `-i` or `--id`not given - exiting")
		}
		opts := &gitlab.ListGroupProjectsOptions{}
		projects, _, err := gitlabClient.Groups.ListGroupProjects(id, opts)
		if err != nil { return err }
		return OutputJson(projects)
	},
}

var groupGetCmd = &cobra.Command{
	Use: "get",
	Short: "Details of a group",
	Long: `Get all details of a group. This command can be accessed without authentication if the group is publicly accessible.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}
		group, _, err := gitlabClient.Groups.GetGroup(id)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

var groupCreateCommand = &cobra.Command{
	Use: "create",
	Short: "New group",
	Long: `Creates a new project group. Available only for users who can create groups.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if name == "" {
			return errors.New("required parameter `-n` or `--name` not given - exiting")
		}
		if path == "" {
			return errors.New("required parameter `-p` or `--path` not given - exiting")
		}
		opts := &gitlab.CreateGroupOptions{
			Name: &name,
			Path: &path,
			Description: &description,
			Visibility: str2Visibility(visibility),
			LFSEnabled: &lfsEnabled,
			RequestAccessEnabled: &requestAccessEnabled,
		}
		group, _, err := gitlabClient.Groups.CreateGroup(opts)
		if err != nil {
			return err
		}
		err = OutputJson(group)
		return err
	},
}

func str2Visibility(s string) *gitlab.VisibilityValue {
	if s == "private" { return gitlab.Visibility(gitlab.PrivateVisibility) }
	if s == "internal" { return gitlab.Visibility(gitlab.InternalVisibility) }
	if s == "public" { return gitlab.Visibility(gitlab.PublicVisibility) }
	return nil
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
	groupCreateCommand.PersistentFlags().StringVarP(&name, "name", "n", "", "(required) the name of the group")
	groupCreateCommand.PersistentFlags().StringVarP(&path, "path", "p", "", "(required) the path of the group")
	groupCreateCommand.PersistentFlags().StringVarP(&description, "description", "d", "", "(optional) the description of the group")
	groupCreateCommand.PersistentFlags().StringVarP(&visibility, "visibility", "v", "private", "(optional) The visibility level of the group. Can be 'private' (default), 'internal', or 'public'.")
	groupCreateCommand.PersistentFlags().BoolVarP(&lfsEnabled, "lfs_enabled", "l", false, "(optional) Enable/disable (default) Large File Storage (LFS) for the projects in this group")
	groupCreateCommand.PersistentFlags().BoolVarP(&requestAccessEnabled, "request_access_enabled", "r", false, "(optional) Allow users to request member access.")
	groupCmd.AddCommand(groupCreateCommand)
}
