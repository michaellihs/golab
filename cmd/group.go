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
	"github.com/xanzy/go-gitlab"
)

var name, newName, path, visibility, description, lfsEnabledString, requestAccessEnabledString, search string

var lfsEnabled, requestAccessEnabled bool

var id, projectId int

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage Gitlab Groups",
	Long:  `Show, create, update and delete Gitlab groups.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("check usage of group with `golab group -h`")
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#list-groups
type groupLsFlags struct {
	SkipGroups   *[]string `flag_name:"skip_groups" type:"array" required:"no" description:"Skip the group IDs passed"`
	AllAvailable *bool     `flag_name:"all_available" type:"bool" required:"no" description:"Show all the groups you have access to (defaults to false for authenticated users)"`
	Search       *string   `flag_name:"search" type:"string" required:"no" description:"Return the list of authorized groups matching the search criteria"`
	OrderBy      *string   `flag_name:"order_by" type:"string" required:"no" description:"Order groups by name or path. Default is name"`
	Sort         *string   `flag_name:"sort" type:"string" required:"no" description:"Order groups in asc or desc order. Default is asc"`
	Statistics   *bool     `flag_name:"statistics" type:"bool" required:"no" description:"Include group statistics (admins only)"`
	Owned        *bool     `flag_name:"owned" type:"boolean" required:"no" description:"Limit to groups owned by the current user"`
}

var groupLsCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupLsFlags{},
	Opts:   &gitlab.ListGroupsOptions{},
	Cmd: &cobra.Command{
		Use:   "ls",
		Short: "List groups",
		Long:  `Get a list of visible groups for the authenticated user.`,
	},
	Run: func(cmd golabCommand) error {
		groups, _, err := gitlabClient.Groups.ListGroups(cmd.Opts.(*gitlab.ListGroupsOptions))
		if err != nil {
			return err
		}
		return OutputJson(groups)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#list-a-group-39-s-projects
type listGroupProjectsFlags struct {
	Id         *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the group owned by the authenticated user"`
	Archived   *bool   `flag_name:"archived" type:"bool" required:"no" description:"Limit by archived status"`
	Visibility *string `flag_name:"visibility" type:"string" required:"no" description:"Limit by visibility public, internal, or private"`
	OrderBy    *string `flag_name:"order_by" type:"string" required:"no" description:"Return projects ordered by id, name, path, created_at, updated_at, or last_activity_at fields. Default is created_at"`
	Sort       *string `flag_name:"sort" type:"string" required:"no" description:"Return projects sorted in asc or desc order. Default is desc"`
	Search     *string `flag_name:"search" type:"string" required:"no" description:"Return list of authorized projects matching the search criteria"`
	Simple     *bool   `flag_name:"simple" type:"bool" required:"no" description:"Return only the ID, URL, name, and path of each project"`
	Owned      *bool   `flag_name:"owned" type:"bool" required:"no" description:"Limit by projects owned by the current user"`
	Starred    *bool   `flag_name:"starred" type:"bool" required:"no" description:"Limit by projects starred by the current user"`
}

var groupProjectsCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &listGroupProjectsFlags{},
	Opts:   &gitlab.ListGroupProjectsOptions{},
	Cmd: &cobra.Command{
		Use:   "projects",
		Short: "List a group's projects",
		Long:  `Get a list of projects in this group. When accessed without authentication, only public projects are returned.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*listGroupProjectsFlags)
		projects, _, err := gitlabClient.Groups.ListGroupProjects(*flags.Id, cmd.Opts.(*gitlab.ListGroupProjectsOptions))
		if err != nil {
			return err
		}
		return OutputJson(projects)
	},
}

// see https://docs.gitlab.com/ce/api/groups.html#details-of-a-group
type groupGetFlags struct {
	Id *string `flag_name:"id" type:"integer/string" required:"yes" description:"The ID or URL-encoded path of the group owned by the authenticated user"`
}

var groupGetCmd = &golabCommand{
	Parent: groupCmd,
	Flags:  &groupGetFlags{},
	Cmd: &cobra.Command{
		Use:   "get",
		Short: "Details of a group",
		Long:  `Get all details of a group. This command can be accessed without authentication if the group is publicly accessible.`,
	},
	Run: func(cmd golabCommand) error {
		flags := cmd.Flags.(*groupGetFlags)
		group, _, err := gitlabClient.Groups.GetGroup(*flags.Id)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

var groupCreateCommand = &cobra.Command{
	Use:   "create",
	Short: "New group",
	Long:  `Creates a new project group. Available only for users who can create groups.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if name == "" {
			return errors.New("required parameter `-n` or `--name` not given - exiting")
		}
		if path == "" {
			return errors.New("required parameter `-p` or `--path` not given - exiting")
		}
		opts := &gitlab.CreateGroupOptions{
			Name:                 &name,
			Path:                 &path,
			Description:          &description,
			Visibility:           str2Visibility(visibility),
			LFSEnabled:           &lfsEnabled,
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

var transferProjectCmd = &cobra.Command{
	Use:   "transfer-project",
	Short: "Transfer project to group",
	Long:  `Transfer a project to the Group namespace. Available only for admin`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}
		if projectId == 0 {
			return errors.New("required parameter `-p` or `--project_id` not given - exiting")
		}
		group, _, err := gitlabClient.Groups.TransferGroup(id, projectId)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update group",
	Long:  `Updates the project group. Only available to group owners and administrators.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}

		trueVal := true;
		falseVal := false

		// we have to provide name and path, even if not set in command
		// therefore we read those values from current group
		currGroup, _, err := gitlabClient.Groups.GetGroup(id)
		if err != nil {
			return err
		}

		opts := &gitlab.UpdateGroupOptions{}
		if newName != "NIL" {
			opts.Name = &newName
		} else {
			opts.Name = &currGroup.Name
		}
		if path != "NIL" {
			opts.Path = &path
		} else {
			opts.Path = &currGroup.Path
		}
		if description != "NIL" {
			opts.Description = &description
		}
		if visibility != "NIL" {
			opts.Visibility = str2Visibility(visibility)
		}
		if lfsEnabledString != "NIL" {
			if lfsEnabledString == "true" || lfsEnabledString == "1" {
				opts.LFSEnabled = &trueVal
			}
			if lfsEnabledString == "false" || lfsEnabledString == "0" {
				opts.LFSEnabled = &falseVal
			}
		}
		if requestAccessEnabledString != "NIL" {
			if requestAccessEnabledString == "true" || requestAccessEnabledString == "1" {
				opts.RequestAccessEnabled = &trueVal
			}
			if requestAccessEnabledString == "false" || requestAccessEnabledString == "0" {
				opts.RequestAccessEnabled = &falseVal
			}
		}

		group, _, err := gitlabClient.Groups.UpdateGroup(id, opts)
		if err != nil {
			return err
		}
		return OutputJson(group)
	},
}

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove group",
	Long:  `Removes group with all projects inside.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if id == 0 {
			return errors.New("required parameter `-i` or `--id` not given - exiting")
		}
		_, err := gitlabClient.Groups.DeleteGroup(id)
		return err
	},
}

var groupSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for group",
	Long:  `Get all groups that match your string in their name or path.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if search == "" {
			return errors.New("required paramter `-s` or `--search` not given - exiting")
		}
		groups, _, err := gitlabClient.Groups.SearchGroup(search)
		if err != nil {
			return err
		}
		return OutputJson(groups)
	},
}

func str2Visibility(s string) *gitlab.VisibilityValue {
	if s == "private" {
		return gitlab.Visibility(gitlab.PrivateVisibility)
	}
	if s == "internal" {
		return gitlab.Visibility(gitlab.InternalVisibility)
	}
	if s == "public" {
		return gitlab.Visibility(gitlab.PublicVisibility)
	}
	return nil
}

func init() {
	groupLsCmd.Init()
	groupProjectsCmd.Init()
	groupGetCmd.Init()
	initGroupCreateCommand()
	initTransferProjectCmd()
	initGroupUpdateCommand()
	initGroupDeleteCommand()
	initGroupSearchCommand()
	RootCmd.AddCommand(groupCmd)
}

func initTransferProjectCmd() {
	transferProjectCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to transfer project to")
	transferProjectCmd.PersistentFlags().IntVarP(&projectId, "project_id", "p", 0, "(required) id of project to be transferred")
	groupCmd.AddCommand(transferProjectCmd)
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

func initGroupUpdateCommand() {
	groupUpdateCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) the id of the group to be updated")
	groupUpdateCmd.PersistentFlags().StringVarP(&newName, "name", "n", "NIL", "(optional) the name of the group")
	groupUpdateCmd.PersistentFlags().StringVarP(&path, "path", "p", "NIL", "(optional) the path of the group")
	groupUpdateCmd.PersistentFlags().StringVarP(&description, "description", "d", "NIL", "(optional) the description of the group")
	groupUpdateCmd.PersistentFlags().StringVarP(&visibility, "visibility", "v", "NIL", "(optional) The visibility level of the group. Can be 'private' (default), 'internal', or 'public'.")
	groupUpdateCmd.PersistentFlags().StringVarP(&lfsEnabledString, "lfs_enabled", "l", "NIL", "(optional) Enable/disable (default) Large File Storage (LFS) for the projects in this group")
	groupUpdateCmd.PersistentFlags().StringVarP(&requestAccessEnabledString, "request_access_enabled", "r", "NIL", "(optional) Allow users to request member access.")
	groupCmd.AddCommand(groupUpdateCmd)
}

func initGroupDeleteCommand() {
	groupDeleteCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "(required) id of group to be deleted")
	groupCmd.AddCommand(groupDeleteCmd)
}

func initGroupSearchCommand() {
	groupSearchCmd.PersistentFlags().StringVarP(&search, "search", "s", "", "(required) search phrase")
	groupCmd.AddCommand(groupSearchCmd)
}
